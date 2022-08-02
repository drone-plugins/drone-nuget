// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	ApiKey          string `envconfig:"PLUGIN_NUGET_APIKEY"`
	NugetUri        string `envconfig:"PLUGIN_NUGET_URI"`
	PackageLocation string `envconfig:"PLUGIN_PACKAGE_LOCATION"`
}

const globalNugetUri = "https://api.nuget.org/v3/index.json"

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func validateAndSetArgs(args *Args) error {
	if args.ApiKey == "" {
		return fmt.Errorf("nuget api key must be set in settings")
	}
	if args.NugetUri == "" {
		args.NugetUri = globalNugetUri
	}
	if args.PackageLocation != "" && !fileExists(args.PackageLocation) {
		return fmt.Errorf("the package location: %s does not exist", args.PackageLocation)
	}
	return nil
}

func walkPath(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.WithError(err).Errorln("error in walk path")
			return nil
		}
		if filepath.Ext(path) == ".nupkg" {
			*files = append(*files, path)
		}
		return nil
	}
}

func pushToNuget(file string, args Args) *exec.Cmd {
	return exec.Command("dotnet", "nuget", "push", file, "--api-key", args.ApiKey, "--source", args.NugetUri, "--skip-duplicate")
}

func Exec(_ context.Context, args Args) error {
	logrus.Debug("Starting ...")

	err := validateAndSetArgs(&args)
	if err != nil {
		return fmt.Errorf("issues with the parameters passed: %w", err)
	}

	var files []string
	// checks if single package location was provided, if not push all.
	if args.PackageLocation == "" {
		err = filepath.Walk(".", walkPath(&files))
		if err != nil {
			logrus.WithError(err).Errorln("Exec plugin, filepath.walk")
			return err
		}
	} else {
		files = append(files, args.PackageLocation)
	}

	if len(files) == 0 {
		logrus.Errorln("No packages to publish ...")
		return nil
	}

	for _, file := range files {
		if file != "" {
			logrus.Debugf("Pushing package: %s ", file)
			cmd := pushToNuget(file, args)
			output, err := cmd.Output()
			if err != nil {
				return err
			}
			logrus.Infof(string(output))
		}
	}
	logrus.Debugln("Finished ...")
	return nil
}
