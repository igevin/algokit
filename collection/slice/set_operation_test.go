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
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		name      string
		a         []int
		b         []int
		expectRes []int
		expectErr error
	}{
		{
			name: "nil",
		},
		{
			name:      "a is empty",
			a:         []int{},
			b:         []int{1},
			expectRes: []int{},
		},
		{
			name:      "b is empty",
			a:         []int{1},
			b:         []int{},
			expectRes: []int{},
		},
		{
			name:      "normal",
			a:         []int{1, 2, 3},
			b:         []int{2, 4},
			expectRes: []int{2},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Intersect(tc.a, tc.b)
			assert.Equal(t, tc.expectErr, err)
			if err != nil {
				return
			}
			sort.Ints(res)
			assert.Equal(t, tc.expectRes, res)
		})
	}
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		name      string
		a         []int
		b         []int
		expectRes []int
		expectErr error
	}{
		{
			name: "nil",
		},
		{
			name:      "a is empty",
			a:         []int{},
			b:         []int{1},
			expectRes: []int{1},
		},
		{
			name:      "b is empty",
			a:         []int{1},
			b:         []int{},
			expectRes: []int{1},
		},
		{
			name:      "normal",
			a:         []int{1, 2, 3},
			b:         []int{2, 4},
			expectRes: []int{1, 2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Union(tc.a, tc.b)
			assert.Equal(t, tc.expectErr, err)
			if err != nil {
				return
			}
			sort.Ints(res)
			assert.Equal(t, tc.expectRes, res)
		})
	}
}

func TestDiff(t *testing.T) {
	testCases := []struct {
		name      string
		a         []int
		b         []int
		expectRes []int
		expectErr error
	}{
		{
			name: "nil",
		},
		{
			name:      "a is empty",
			a:         []int{},
			b:         []int{1},
			expectRes: []int{},
		},
		{
			name:      "b is empty",
			a:         []int{1},
			b:         []int{},
			expectRes: []int{1},
		},
		{
			name:      "normal",
			a:         []int{1, 2, 3},
			b:         []int{2, 4},
			expectRes: []int{1, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Diff(tc.a, tc.b)
			assert.Equal(t, tc.expectErr, err)
			if err != nil {
				return
			}
			sort.Ints(res)
			assert.Equal(t, tc.expectRes, res)
		})
	}
}
