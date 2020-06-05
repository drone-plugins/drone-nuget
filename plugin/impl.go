// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"

	"github.com/drone-plugins/drone-nuget/plugin/cli"
)

// Settings for the plugin.
type Settings struct {
	APIKey string
	Source string
	Name   string
	File   string
}

const (
	nugetOrgName   = "nuget.org"
	nugetOrgSource = "https://api.nuget.org/v3/index.json"
)

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	if p.settings.APIKey == "" {
		return fmt.Errorf("no api key provided")
	}

	// Verify specified file
	if p.settings.File == "" {
		return fmt.Errorf("no package specified")
	}
	info, err := os.Stat(p.settings.File)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %w", err)
	}
	if info.IsDir() {
		return fmt.Errorf("file is a directory")
	}

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

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	nuget, err := cli.New()
	if err != nil {
		return err
	}

	cmds := []*exec.Cmd{
		nuget.VersionCmd(),
	}

	if p.settings.Name != nugetOrgName {
		cmds = append(cmds, nuget.AddSourceCmd(p.settings.Source, p.settings.Name))
	}

	cmds = append(cmds, nuget.ListSourcesCmd())
	cmds = append(cmds, nuget.PushPackageCmd(p.settings.File, p.settings.Name, p.settings.APIKey))

	return cli.RunCommands(cmds, "")
}
