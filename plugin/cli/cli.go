// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// NuGet is an interface for the command line interface for NuGet.
type NuGet interface {
	// VersionCmd creates a command that will display the version of NuGet.
	VersionCmd() *exec.Cmd

	// ListSources creates a command that will display the NuGet package sources.
	ListSourcesCmd() *exec.Cmd

	// AddSourceCmd creates a command that will add a NuGet package source.
	AddSourceCmd(source, name string) *exec.Cmd

	// PushPackageCmd creates a command that will push a package to a NuGet package source.
	PushPackageCmd(path, name, key string) *exec.Cmd
}

func New() (NuGet, error) {
	if n, err := NewDotNet(); err == nil {
		return n, err
	}

	return NewNuGet()
}

func NewDotNet() (NuGet, error) {
	path, err := exec.LookPath("dotnet")
	if err != nil {
		return nil, fmt.Errorf("could not find dotnet in path")
	}

	return &dotnet{
		path: path,
	}, nil
}

func NewNuGet() (NuGet, error) {
	path, err := exec.LookPath("nuget")
	if err != nil {
		return nil, fmt.Errorf("could not find nuget in path")
	}

	return &nuget{
		path: path,
	}, nil
}

// runCommands executes the list of cmds in the given directory.
func RunCommands(cmds []*exec.Cmd, dir string) error {
	for _, cmd := range cmds {
		err := RunCommand(cmd, dir)

		if err != nil {
			return err
		}
	}

	return nil
}

func RunCommand(cmd *exec.Cmd, dir string) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	trace(cmd)

	return cmd.Run()
}

// trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}
