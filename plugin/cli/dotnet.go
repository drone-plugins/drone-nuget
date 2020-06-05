// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package cli

import (
	"os/exec"
)

type dotnet struct {
	path string
}

func (n *dotnet) VersionCmd() *exec.Cmd {
	return exec.Command("dotnet", "nuget", "--version")
}

func (n *dotnet) ListSourcesCmd() *exec.Cmd {
	return exec.Command("dotnet", "nuget", "list", "source")
}