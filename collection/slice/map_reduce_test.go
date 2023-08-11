// Copyright 2023 igevin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	f := func(a int) int { return a * a }
	testCases := []struct {
		name        string
		input       []int
		f           func(int) int
		expectedRes []int
	}{
		{
			name:        "nil",
			input:       nil,
			expectedRes: nil,
		},
		{
			name:        "zero",
			input:       []int{},
			expectedRes: []int{},
		},
		{
			name:        "one element",
			input:       []int{2},
			f:           f,
			expectedRes: []int{4},
		},
		{
			name:        "normal",
			input:       []int{1, 2, 3, 4},
			f:           f,
			expectedRes: []int{1, 4, 9, 16},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Map(tc.input, tc.f)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func TestFilter(t *testing.T) {
	f := func(a int) bool { return a%2 == 0 }
	testCases := []struct {
		name        string
		input       []int
		f           func(int) bool
		expectedRes []int
	}{
		{
			name:        "nil",
			input:       nil,
			expectedRes: nil,
		},
		{
			name:        "zero",
			input:       []int{},
			expectedRes: []int{},
		},
		{
			name:        "one and kept",
			input:       []int{2},
			f:           f,
			expectedRes: []int{2},
		},
		{
			name:        "one and filtered",
			input:       []int{1},
			f:           f,
			expectedRes: []int{},
		},
		{
			name:        "normal",
			input:       []int{1, 2, 3, 4},
			f:           f,
			expectedRes: []int{2, 4},
		},
		{
			name:        "normal, all filtered",
			input:       []int{1, 3, 5, 7},
			f:           f,
			expectedRes: []int{},
		},
		{
			name:        "normal, all kept",
			input:       []int{2, 4, 6, 8},
			f:           f,
			expectedRes: []int{2, 4, 6, 8},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Filter(tc.input, tc.f)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func TestReduce(t *testing.T) {
	f := func(a, b int) int { return a + b }
	testCases := []struct {
		name        string
		input       []int
		f           func(int, int) int
		expectedRes int
	}{
		{
			name:        "nil",
			input:       nil,
			expectedRes: 0,
		},
		{
			name:        "zero",
			input:       []int{},
			expectedRes: 0,
		},
		{
			name:        "one",
			input:       []int{1},
			f:           f,
			expectedRes: 1,
		},
		{
			name:        "normal",
			input:       []int{1, 2, 3, 4},
			f:           f,
			expectedRes: 10,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Reduce(tc.input, tc.f)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
