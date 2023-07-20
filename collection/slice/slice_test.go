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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type deleteFunc[T any] func(src []T, index int) ([]T, error)

func TestDelete(t *testing.T) {
	testDelete(t, Delete[int])
}

func TestDeleteV0(t *testing.T) {
	testDelete(t, DeleteV0[int])
}

func testDelete(t *testing.T, delete deleteFunc[int]) {
	testCases := []struct {
		name        string
		src         []int
		index       int
		expectedRes []int
		expectedErr error
	}{
		{
			name:        "nil",
			src:         nil,
			expectedErr: ErrOutOfRange,
		},
		{
			name:        "empty",
			src:         []int{},
			index:       0,
			expectedErr: ErrOutOfRange,
		},
		{
			name:        "boundary",
			src:         []int{1, 2, 3},
			index:       3,
			expectedErr: ErrOutOfRange,
		},
		{
			name:        "out of range",
			src:         []int{1, 2, 3},
			index:       4,
			expectedErr: ErrOutOfRange,
		},
		{
			name:        "first",
			src:         []int{1, 2, 3},
			index:       0,
			expectedRes: []int{2, 3},
		},
		{
			name:        "last",
			src:         []int{1, 2, 3},
			index:       2,
			expectedRes: []int{1, 2},
		},
		{
			name:        "normal",
			src:         []int{1, 2, 3},
			index:       1,
			expectedRes: []int{1, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := delete(tc.src, tc.index)
			require.Equal(t, tc.expectedErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func TestIndex(t *testing.T) {
	testCases := []struct {
		name        string
		src         []int
		target      int
		expectedRes int
	}{
		{
			name:        "nil",
			src:         nil,
			target:      1,
			expectedRes: -1,
		},
		{
			name:        "empty",
			src:         []int{},
			target:      1,
			expectedRes: -1,
		},
		{
			name:        "not found",
			src:         []int{1, 2, 3},
			target:      4,
			expectedRes: -1,
		},
		{
			name:        "found",
			src:         []int{1, 2, 3},
			target:      3,
			expectedRes: 2,
		},
		{
			name:        "duplicated target",
			src:         []int{1, 2, 3, 4, 3},
			target:      3,
			expectedRes: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Index(tc.src, tc.target)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	testCases := []struct {
		name        string
		src         []int
		target      int
		expectedRes int
	}{
		{
			name:        "nil",
			src:         nil,
			target:      1,
			expectedRes: -1,
		},
		{
			name:        "empty",
			src:         []int{},
			target:      1,
			expectedRes: -1,
		},
		{
			name:        "not found",
			src:         []int{1, 2, 3},
			target:      4,
			expectedRes: -1,
		},
		{
			name:        "found",
			src:         []int{1, 2, 3},
			target:      3,
			expectedRes: 2,
		},
		{
			name:        "duplicated target",
			src:         []int{1, 2, 3, 4, 3},
			target:      3,
			expectedRes: 4,
		},
		{
			name:        "duplicated target2",
			src:         []int{1, 2, 3, 4, 3, 5},
			target:      3,
			expectedRes: 4,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := LastIndexOf(tc.src, tc.target)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func TestIndexAll(t *testing.T) {
	testCases := []struct {
		name        string
		src         []int
		target      int
		expectedRes []int
	}{
		{
			name:        "nil",
			src:         nil,
			target:      1,
			expectedRes: make([]int, 0, 0),
		},
		{
			name:        "empty",
			src:         []int{},
			target:      1,
			expectedRes: make([]int, 0, 0),
		},
		{
			name:        "not found",
			src:         []int{1, 2, 3},
			target:      4,
			expectedRes: make([]int, 0, 3),
		},
		{
			name:        "found 1",
			src:         []int{1, 2, 3},
			target:      3,
			expectedRes: []int{2},
		},
		{
			name:        "found 2",
			src:         []int{1, 2, 3, 4, 3},
			target:      3,
			expectedRes: []int{2, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IndexAll(tc.src, tc.target)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
