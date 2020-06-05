// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package cli

import (
	"os/exec"
)

type nuget struct {
	path string
}

func (n *nuget) VersionCmd() *exec.Cmd {
	return exec.Command("nuget")
}

func (n *nuget) ListSourcesCmd() *exec.Cmd {
	return exec.Command("nuget", "sources")
}
