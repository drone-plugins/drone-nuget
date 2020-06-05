// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package main

import (
	"github.com/drone-plugins/drone-nuget/plugin"
	"github.com/urfave/cli/v2"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "source",
			Usage:       "nuget package repository source url",
			EnvVars:     []string{"PLUGIN_SOURCE"},
			Destination: &settings.Source,
		},
		&cli.StringFlag{
			Name:        "name",
			Usage:       "nuget package repository name",
			EnvVars:     []string{"PLUGIN_NAME"},
			Destination: &settings.Name,
		},
	}
}
