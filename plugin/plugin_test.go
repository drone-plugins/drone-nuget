// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateArgs(t *testing.T) {
	keyExists := Args{ApiKey: "someKey"}
	noKey := Args{ApiKey: ""}

	// assert equality
	assert.Nil(t, validateAndSetArgs(&keyExists), "result is nil")
	assert.NotNil(t, noKey)
}

func TestValidationArgsPackageLocation(t *testing.T) {
	packageExists := Args{ApiKey: "someKey", PackageLocation: "testdata"}
	noPackage := Args{ApiKey: "someKey", PackageLocation: "/notExist"}

	// assert equality
	assert.Nil(t, validateAndSetArgs(&packageExists))
	assert.Error(t, validateAndSetArgs(&noPackage))
}

func TestExecNoKeySupplied(t *testing.T) {
	noKey := Args{ApiKey: ""}

	result := Exec(context.Background(), noKey)

	assert.Contains(t, result.Error(), "issues with the parameters passed")
}
