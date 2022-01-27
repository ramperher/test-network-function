// Copyright (C) 2020-2022 Red Hat, Inc.
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

package networking

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/test-network-function/test-network-function/pkg/tnf/interactive"
)

func TestParseVariables(t *testing.T) {
	// expected inputs
	testCases := []struct {
		// inputs
		inputRes      string
		declaredPorts map[key]string

		// expected outputs here
		expectedDeclaredPorts map[key]string
		exceptedRes           string
	}{
		{
			inputRes:              `jsonBlobHere`,
			declaredPorts:         map[key]string{},
			expectedDeclaredPorts: map[key]string{},
			exceptedRes:           `jsonBlobHere`,
		},
	}

	for _, tc := range testCases {
		err := parseVariables(tc.inputRes, tc.declaredPorts)
		assert.Nil(t, err)
		assert.Equal(t, tc.declaredPorts, tc.expectedDeclaredPorts)
	}
}

func TestDeclaredPortList(t *testing.T) {
	// expected inputs
	testCases := []struct {
		// inputs
		container     int
		podName       string
		podNamespace  string
		declaredPorts map[key]string

		// expected outputs here
		expectedDeclaredPorts map[key]string
	}{
		{
			container:             12,
			podName:               "xxx",
			podNamespace:          "yy",
			declaredPorts:         map[key]string{},
			expectedDeclaredPorts: map[key]string{},
		},
	}

	for _, tc := range testCases {
		err := declaredPortList(tc.container, tc.podName, tc.podNamespace, tc.declaredPorts)
		assert.Nil(t, err)
		assert.Equal(t, tc.declaredPorts, tc.expectedDeclaredPorts)
	}
}

func TestListeningPortList(t *testing.T) {
	// expected inputs
	testCases := []struct {
		// inputs
		commandlisten  []string
		nodeOc         *interactive.Context
		listeningPorts map[key]string

		// expected outputs here
		expectedlisteningPorts map[key]string
	}{
		{
			commandlisten:          []string{"ss -tulwnH"},
			nodeOc:                 "",
			listeningPorts:         map[key]string{},
			expectedlisteningPorts: map[key]string{},
		},
	}
	for _, tc := range testCases {
		err := listeningPortList(tc.commandlisten, tc.nodeOc, tc.listeningPorts)
		assert.Nil(t, err)
		assert.Equal(t, tc.listeningPorts, tc.expectedlisteningPorts)
	}
}

func TestCheckIfListenIsDeclared(t *testing.T) {
	// expected inputs
	testCases := []struct {
		// inputs
		listeningPorts map[key]string
		declaredPorts  map[key]string

		// expected outputs here
		expectedres map[key]string
	}{
		{
			listeningPorts: map[key]string{},
			declaredPorts:  map[key]string{},
			expectedres:    map[key]string{},
		},
	}
	for _, tc := range testCases {
		err := checkIfListenIsDeclared(tc.listeningPorts, tc.declaredPorts)
		assert.Nil(t, err)
		assert.Equal(t, tc.listeningPorts, tc.expectedres)
	}
}