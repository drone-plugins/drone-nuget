// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type (
	// Settings for the plugin.
	Settings struct {
		APIKey string
		Source string
		Name   string
		File   string

		nupkg string
	}

	nuspecMetadata struct {
		Name    string `xml:"id"`
		Version string `xml:"version"`
	}
)

const (
	nugetOrgName   = "nuget.org"
	nugetOrgSource = "https://api.nuget.org/v3/index.json"
)

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	if p.settings.APIKey == "" {
		return fmt.Errorf("no api key provided")
	}

	file := p.settings.File
	if file == "" {
		return fmt.Errorf("no package specified")
	}

	// Convert to / separators from os specific ones and then use path
	// Windows works fine with / but unix does not work with \
	file = filepath.ToSlash(file)

	// Clean the path
	file = path.Clean(file)

	// Determine file type
	if strings.HasSuffix(file, ".nuspec") {
		nuspec := file
		logrus.WithField("file", nuspec).Info("Loading .nuspec file")

		var err error
		file, err = nupkgFromNuspec(nuspec)
		if err != nil {
			return fmt.Errorf("could not determine nupkg file from %s: %w", nuspec, err)
		}
	} else if !strings.HasSuffix(file, ".nupkg") {
		return fmt.Errorf("file %s isn't a nuspec or a nupkg", file)
	}

	if !fileExists(file) {
		return fmt.Errorf(".nupkg file does not exist at %s", file)
	}

	logrus.WithField("file", file).Info("Publishing .nupkg file")

	// Store nupkg file and convert to os specific path
	p.settings.nupkg = filepath.FromSlash(file)

	// Set defaults for source and name
	if p.settings.Source == "" {
		p.settings.Source = nugetOrgSource
	}
	if p.settings.Name == "" {
		if p.settings.Source == nugetOrgSource {
			p.settings.Name = nugetOrgName
		} else {
			p.settings.Name = "drone-nuget"
		}
	}

	// Validate the source and name
	if p.settings.Name == nugetOrgName && p.settings.Source != nugetOrgSource {
		return fmt.Errorf("repository named %s must use %s as its source", nugetOrgName, nugetOrgSource)
	}
	if p.settings.Name != nugetOrgName && p.settings.Source == nugetOrgSource {
		return fmt.Errorf("source %s must use %s as its repository name", nugetOrgSource, nugetOrgName)
	}
	if _, err := url.Parse(p.settings.Source); err != nil {
		return fmt.Errorf("could not parse source url %s: %w", p.settings.Source, err)
	}

	logrus.WithFields(logrus.Fields{
		"name":   p.settings.Name,
		"source": p.settings.Source,
	}).Info("Using NuGet repository")

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	cmds := []*exec.Cmd{
		versionCmd(),
	}

	if p.settings.Name != nugetOrgName {
		cmds = append(cmds, addSourceCmd(p.settings.Source, p.settings.Name))
	}

	cmds = append(cmds, listSourcesCmd())
	cmds = append(cmds, pushPackageCmd(p.settings.nupkg, p.settings.Name, p.settings.APIKey))

	return runCommands(cmds, "")
}

// nupkgFromNuspec reads the nuspec file and determines the filename.
func nupkgFromNuspec(file string) (string, error) {
	if !fileExists(file) {
		return "", fmt.Errorf(".nuspec file not found at %s", file)
	}

	// Read the file
	nuspec, err := os.Open(file)
	if err != nil {
		return "", fmt.Errorf("unable to open .nuspec file %s: %w", file, err)
	}

	bytes, err := ioutil.ReadAll(nuspec)
	if err != nil {
		return "", fmt.Errorf("unable to read .nuspec file %s: %w", file, err)
	}

	// Unmarshal the file
	var doc struct {
		XMLName  xml.Name       `xml:"package"`
		Metadata nuspecMetadata `xml:"metadata"`
	}
	if err := xml.Unmarshal(bytes, &doc); err != nil {
		return "", fmt.Errorf("unable to parse .nuspec file %s: %w", file, err)
	}

	logrus.WithFields(logrus.Fields{
		"name":    doc.Metadata.Name,
		"version": doc.Metadata.Version,
	}).Info("Found .nupsec file")

	nupkgName := fmt.Sprintf("%s.%s.nupkg", doc.Metadata.Name, doc.Metadata.Version)

	return path.Join(path.Dir(file), nupkgName), nil
}

// fileExists determines if the file is present.
func fileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// versionCmd gets the nuget version.
func versionCmd() *exec.Cmd {
	return exec.Command("dotnet", "nuget", "--version")
}

// listSourcesCmd lists the nuget repositories.
func listSourcesCmd() *exec.Cmd {
	return exec.Command("dotnet", "nuget", "list", "source")
}

// addSourceCmd creates a new nuget repository source.
func addSourceCmd(source, name string) *exec.Cmd {
	return exec.Command("dotnet", "nuget", "add", "source", source, "--name", name)
}

// pushPackageCmd pushes a package to the nuget repository.
func pushPackageCmd(path, name, key string) *exec.Cmd {
	return exec.Command("dotnet", "nuget", "push", path, "--source", name, "--api-key", key)
}

// trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}

// runCommands executes the list of cmds in the given directory.
func runCommands(cmds []*exec.Cmd, dir string) error {
	for _, cmd := range cmds {
		err := runCommand(cmd, dir)

		if err != nil {
			return err
		}
	}

	return nil
}

// runCommand executes a cmd in the given directory.
func runCommand(cmd *exec.Cmd, dir string) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	trace(cmd)

	return cmd.Run()
}
