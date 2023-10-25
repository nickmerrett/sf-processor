//
// Copyright (C) 2023 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package falco implements a frontend for (extended) Falco rules engine.
package falco_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy/falco"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/flatrecord"
)

var rulesPath string = "../../../../resources/policies/runtimeintegrity"

func TestMain(m *testing.M) {
	logger.InitLoggers(logger.TRACE)
	os.Exit(m.Run())
}

func TestCompiler(t *testing.T) {
	pc := falco.NewPolicyCompiler(flatrecord.NewOperations())
	paths, err := ioutils.ListRecursiveFilePaths(rulesPath, ".yaml")
	assert.NoError(t, err)
	_, _, err = pc.Compile(paths...)
	assert.NoError(t, err)
}
