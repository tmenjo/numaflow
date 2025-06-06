/*
Copyright 2022 The Numaproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"strings"
	"testing"

	"github.com/numaproj/numaflow/pkg/apis/numaflow/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	str := RandomString(20)
	assert.Equal(t, 20, len(str))
}

func TestRandomLowercaseString(t *testing.T) {
	str := RandomLowerCaseString(20)
	assert.Equal(t, 20, len(str))
	assert.Equal(t, str, strings.ToLower(str))
}

func TestCompareSlice(t *testing.T) {
	tests := []struct {
		name     string
		sliceA   []string
		sliceB   []string
		operator v1alpha1.LogicOperator
		expected bool
	}{
		{
			name: "or_true",
			sliceA: []string{
				"even",
				"odd",
				"prime",
			},
			sliceB: []string{
				"even",
				"ab",
				"bc",
			},
			operator: v1alpha1.LogicOperatorOr,
			expected: true,
		},
		{
			name: "and_true",
			sliceA: []string{
				"even",
				"odd",
				"prime",
			},
			sliceB: []string{
				"even",
				"odd",
				"prime",
				"abc",
			},
			operator: v1alpha1.LogicOperatorAnd,
			expected: true,
		},
		{
			name: "not_true",
			sliceA: []string{
				"even",
				"odd",
				"prime",
			},
			sliceB: []string{
				"abc",
				"bca",
				"cab",
			},
			operator: v1alpha1.LogicOperatorNot,
			expected: true,
		},
		{
			name: "or_false",
			sliceA: []string{
				"even",
				"odd",
				"prime",
			},
			sliceB: []string{
				"abc",
				"bca",
				"cab",
			},
			operator: v1alpha1.LogicOperatorOr,
			expected: false,
		},
		{
			name: "not_false",
			sliceA: []string{
				"even",
				"odd",
				"prime",
			},
			sliceB: []string{
				"even",
				"bca",
				"cab",
			},
			operator: v1alpha1.LogicOperatorNot,
			expected: false,
		},
		{
			name: "and_false",
			sliceA: []string{
				"even",
				"odd",
				"prime",
			},
			sliceB: []string{
				"even",
				"odd",
				"abc",
			},
			operator: v1alpha1.LogicOperatorAnd,
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, CompareSlice(tt.operator, tt.sliceA, tt.sliceB), tt.expected)
		})
	}
}

func TestDNS1035(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple lowercase conversion",
			input:    "HELLO",
			expected: "hello",
		},
		{
			name:     "replace special characters",
			input:    "hello@world!123",
			expected: "hello-world-123",
		},
		{
			name:     "multiple consecutive special chars",
			input:    "hello!!!world###123",
			expected: "hello-world-123",
		},
		{
			name:     "spaces and underscores",
			input:    "hello_world space test",
			expected: "hello-world-space-test",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only special characters",
			input:    "@#$%^&*",
			expected: "-",
		},
		{
			name:     "mixed case with numbers and hyphens",
			input:    "My-Cool-Service123",
			expected: "my-cool-service123",
		},
		{
			name:     "unicode characters",
			input:    "héllo→wörld",
			expected: "h-llo-w-rld",
		},
		{
			name:     "empty",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DNS1035(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHashcode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "a",
			input:    "1",
			expected: 2212294583,
		},
		{
			name:     "b",
			input:    "hello world",
			expected: 222957957,
		},
		{
			name:     "c",
			input:    "This is a test",
			expected: 3229261618,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Hashcode(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
