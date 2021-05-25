// Copyright (C) 2021 Red Hat, Inc.
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

package junit_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/test-network-function/test-network-function/pkg/junit"
)

const (
	testJunitXMLFilename = "success.junit.xml"
	testKey              = "cnf-certification-tests_junit"
)

func TestExtractTestSuiteResults(t *testing.T) {
	junitResults, err := junit.ExportJUnitAsJSON(path.Join("testdata", testJunitXMLFilename))
	claim := make(map[string]interface{})
	claim[testKey] = junitResults
	assert.Nil(t, err)
	assert.NotNil(t, junitResults)
	results, err := junit.ExtractTestSuiteResults(claim, testKey)
	assert.Nil(t, err)
	// positive test
	assert.Equal(t, true, results["generic when Reading namespace of test/test Should not be 'default' and should not begin with 'openshift-'"].Passed)
	// negative test
	assert.Equal(t, false, results["generic when Testing owners of CNF pod Should contain at least one of kind DaemonSet/ReplicaSet"].Passed, false)
}
