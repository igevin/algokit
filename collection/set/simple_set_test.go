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

package set

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSimpleSet(t *testing.T) {
	s := NewSimpleSet[int]()
	require.NotNil(t, s)

	s = NewSimpleSet[int](WithNewSizeOption[int](10))
	require.NotNil(t, s)
}

func TestSimpleSet_Add(t *testing.T) {
	testCases := []struct {
		name        string
		set         SimpleSet[int]
		input       []int
		expectRes   map[int]struct{}
		expectError error
	}{
		{
			name:      "add one",
			set:       *(NewSimpleSet[int]()),
			input:     []int{1},
			expectRes: map[int]struct{}{1: {}},
		},
		{
			name:      "add two same",
			set:       *(NewSimpleSet[int]()),
			input:     []int{1, 1},
			expectRes: map[int]struct{}{1: {}},
		},
		{
			name:      "add a lot",
			set:       *(NewSimpleSet[int]()),
			input:     []int{1, 2, 1, 3, 4, 3, 2, 5, 4, 6},
			expectRes: map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, el := range tc.input {
				err := tc.set.Add(el)
				require.NoError(t, err)
			}
			assert.Equal(t, tc.expectRes, tc.set.m)
		})
	}
}

func TestSimpleSet_AddAll(t *testing.T) {
	testCases := []struct {
		name        string
		set         SimpleSet[int]
		input       []int
		expectRes   map[int]struct{}
		expectError error
	}{
		{
			name:      "add nil",
			set:       *(NewSimpleSet[int]()),
			input:     nil,
			expectRes: map[int]struct{}{},
		},
		{
			name:      "add one",
			set:       *(NewSimpleSet[int]()),
			input:     []int{1},
			expectRes: map[int]struct{}{1: {}},
		},
		{
			name:      "add a lot",
			set:       *(NewSimpleSet[int]()),
			input:     []int{1, 2, 1, 3, 4, 3, 2, 5, 4, 6},
			expectRes: map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.set.AddAll(tc.input)
			require.NoError(t, err)
			assert.Equal(t, tc.expectRes, tc.set.m)
		})
	}
}

func TestSimpleSet_Contains(t *testing.T) {
	testCases := []struct {
		name        string
		set         SimpleSet[int]
		input       int
		expectRes   bool
		expectError error
	}{
		{
			name: "contains",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2, 3})
				return *s
			}(),
			input:     1,
			expectRes: true,
		},
		{
			name: "not contains",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2, 3})
				return *s
			}(),
			input:     5,
			expectRes: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.set.Contains(tc.input)
			assert.Equal(t, tc.expectRes, res)
		})
	}
}

func TestSimpleSet_IsEmpty(t *testing.T) {
	testCases := []struct {
		name      string
		set       SimpleSet[int]
		expectRes bool
	}{
		{
			name: "empty",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				return *s
			}(),
			expectRes: true,
		},
		{
			name: "not empty",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2, 3})
				return *s
			}(),
			expectRes: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.set.IsEmpty()
			assert.Equal(t, tc.expectRes, res)
		})
	}
}

func TestSimpleSet_Size(t *testing.T) {
	testCases := []struct {
		name      string
		set       SimpleSet[int]
		expectRes int
	}{
		{
			name: "empty",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				return *s
			}(),
			expectRes: 0,
		},
		{
			name: "not empty",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2, 3})
				return *s
			}(),
			expectRes: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.set.Size()
			assert.Equal(t, tc.expectRes, res)
		})
	}
}

func TestSimpleSet_ToSlice(t *testing.T) {
	testCases := []struct {
		name      string
		set       SimpleSet[int]
		expectRes []int
	}{
		{
			name: "empty",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				return *s
			}(),
			expectRes: []int{},
		},
		{
			name: "not empty",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2, 3})
				return *s
			}(),
			expectRes: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.set.ToSlice()
			// todo: 实现slice的排序功能后，改为比较排序后的切片
			resSet := NewSimpleSet[int]()
			_ = resSet.AddAll(res)
			expectSet := NewSimpleSet[int]()
			_ = expectSet.AddAll(tc.expectRes)
			assert.Equal(t, expectSet, resSet)
		})
	}
}

func TestSimpleSet_Clear(t *testing.T) {
	testCases := []struct {
		name      string
		set       SimpleSet[int]
		expectRes map[int]struct{}
	}{
		{
			name: "empty",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				return *s
			}(),
			expectRes: map[int]struct{}{},
		},
		{
			name: "not empty",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2, 3})
				return *s
			}(),
			expectRes: map[int]struct{}{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.set.Clear()
			assert.Equal(t, tc.expectRes, tc.set.m)
		})
	}
}

func TestSimpleSet_Intersect(t *testing.T) {
	testCases := []struct {
		name        string
		set         SimpleSet[int]
		target      Set[int]
		expectRes   Set[int]
		expectError error
	}{
		// todo 增加不同类型的Set交叉Intersect的实现
		{
			name: "target is nil",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2, 3})
				return *s
			}(),
			target: nil,
			expectRes: func() *SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2, 3})
				return s
			}(),
		},
		{
			name: "target is empty",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2, 3})
				return *s
			}(),
			target: func() *SimpleSet[int] {
				s := NewSimpleSet[int]()
				return s
			}(),
			expectRes: func() *SimpleSet[int] {
				s := NewSimpleSet[int]()
				return s
			}(),
		},
		{
			name: "set is empty",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				return *s
			}(),
			target: func() *SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2})
				return s
			}(),
			expectRes: func() *SimpleSet[int] {
				s := NewSimpleSet[int]()
				return s
			}(),
		},
		{
			name: "normal",
			set: func() SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2, 3})
				return *s
			}(),
			target: func() *SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2})
				return s
			}(),
			expectRes: func() *SimpleSet[int] {
				s := NewSimpleSet[int]()
				_ = s.AddAll([]int{1, 2})
				return s
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.set.Intersect(tc.target)
			assert.Equal(t, tc.expectError, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.expectRes, res)
		})
	}
}
