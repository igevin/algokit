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

package skiplist

import (
	"math"
	"testing"

	"github.com/igevin/algokit/comparator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList[int](comparator.PrimeComparator[int])
	require.NotNil(t, sl)
	require.Equal(t, 1, sl.levels)
	require.Equal(t, 0, sl.length)
	header := sl.header
	require.Equal(t, maxLevel, header.level)
}

func TestSkipListNormal(t *testing.T) {
	sl := NewSkipList[int](comparator.PrimeComparator[int])
	require.Equal(t, 0, sl.Len())
	vals := []int{3, 2, 5, 1, 4}
	m := math.MaxInt
	for i, val := range vals {
		if val < m {
			m = val
		}
		err := sl.Insert(val)
		assert.NoError(t, err)
		assert.Equal(t, i+1, sl.Len())
		f, err := sl.Front()
		require.NoError(t, err)
		assert.Equal(t, m, f)
	}
	res := asSlice(sl)
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, res)

	err := sl.Insert(4)
	assert.Equal(t, ErrSameNode, err)

	val, err := sl.Find(5)
	require.NoError(t, err)
	require.Equal(t, 5, val)
	_, err = sl.Find(6)
	require.Equal(t, ErrNodeNotFound, err)

	sl.Delete(6)
	res = asSlice(sl)
	require.Equal(t, expected, res)

	sl.Delete(3)
	res = asSlice(sl)
	expected = []int{1, 2, 4, 5}
	require.Equal(t, expected, res)
	require.Equal(t, len(expected), sl.Len())
}

func TestSkipList_Insert(t *testing.T) {
	testCases := []struct {
		name        string
		sl          *SkipList[int]
		val         int
		expectRes   []int
		expectError error
	}{
		{
			name:      "insert to empty skip list",
			sl:        NewSkipList[int](comparator.PrimeComparator[int]),
			val:       1,
			expectRes: []int{1},
		},
		{
			name: "insert to not empty skip list",
			sl: func() *SkipList[int] {
				sl := NewSkipList[int](comparator.PrimeComparator[int])
				_ = sl.Insert(1)
				return sl
			}(),
			val:       2,
			expectRes: []int{1, 2},
		},
		{
			name: "insert same value to not empty skip list",
			sl: func() *SkipList[int] {
				sl := NewSkipList[int](comparator.PrimeComparator[int])
				_ = sl.Insert(1)
				return sl
			}(),
			val:         1,
			expectError: ErrSameNode,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := tc.sl.Len()
			err := tc.sl.Insert(tc.val)
			assert.Equal(t, tc.expectError, err)
			if err != nil {
				return
			}
			assert.Equal(t, l+1, tc.sl.Len())
			assert.Equal(t, tc.expectRes, asSlice(tc.sl))
		})
	}
}

func TestSkipList_Find(t *testing.T) {
	testCases := []struct {
		name        string
		sl          *SkipList[int]
		val         int
		expectRes   int
		expectError error
	}{
		{
			name:        "empty skip list",
			sl:          NewSkipList[int](comparator.PrimeComparator[int]),
			val:         1,
			expectError: ErrNodeNotFound,
		},
		{
			name: " skip list, 1 element, not found",
			sl: func() *SkipList[int] {
				sl := NewSkipList[int](comparator.PrimeComparator[int])
				_ = sl.Insert(1)
				return sl
			}(),
			val:         2,
			expectError: ErrNodeNotFound,
		},
		{
			name: " skip list, 1 element, found",
			sl: func() *SkipList[int] {
				sl := NewSkipList[int](comparator.PrimeComparator[int])
				_ = sl.Insert(1)
				return sl
			}(),
			val:       1,
			expectRes: 1,
		},
		{
			name: "skip list, more elements, not found",
			sl: func() *SkipList[int] {
				sl := NewSkipList[int](comparator.PrimeComparator[int])
				for _, el := range []int{1, 2, 3, 4, 5} {
					_ = sl.Insert(el)
				}

				return sl
			}(),
			val:         6,
			expectError: ErrNodeNotFound,
		},
		{
			name: "skip list, more elements,  found",
			sl: func() *SkipList[int] {
				sl := NewSkipList[int](comparator.PrimeComparator[int])
				for _, el := range []int{1, 2, 3, 4, 5} {
					_ = sl.Insert(el)
				}

				return sl
			}(),
			val:       5,
			expectRes: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.sl.Find(tc.val)
			assert.Equal(t, tc.expectError, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.expectRes, res)
		})
	}
}

func TestSkipList_Delete(t *testing.T) {
	testCases := []struct {
		name        string
		sl          *SkipList[int]
		val         int
		expectRes   []int
		expectError error
	}{
		{
			name:      "empty skip list",
			sl:        NewSkipList[int](comparator.PrimeComparator[int]),
			val:       1,
			expectRes: []int{},
		},
		{
			name: " skip list, 1 element, not found",
			sl: func() *SkipList[int] {
				sl := NewSkipList[int](comparator.PrimeComparator[int])
				_ = sl.Insert(1)
				return sl
			}(),
			val:       2,
			expectRes: []int{1},
		},
		{
			name: " skip list, 1 element, found",
			sl: func() *SkipList[int] {
				sl := NewSkipList[int](comparator.PrimeComparator[int])
				_ = sl.Insert(1)
				return sl
			}(),
			val:       1,
			expectRes: []int{},
		},
		{
			name: "skip list, more elements, not found",
			sl: func() *SkipList[int] {
				sl := NewSkipList[int](comparator.PrimeComparator[int])
				for _, el := range []int{1, 2, 3, 4, 5} {
					_ = sl.Insert(el)
				}

				return sl
			}(),
			val:       6,
			expectRes: []int{1, 2, 3, 4, 5},
		},
		{
			name: "skip list, more elements,  found",
			sl: func() *SkipList[int] {
				sl := NewSkipList[int](comparator.PrimeComparator[int])
				for _, el := range []int{1, 2, 3, 4, 5} {
					_ = sl.Insert(el)
				}

				return sl
			}(),
			val:       5,
			expectRes: []int{1, 2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.sl.Delete(tc.val)
			assert.Equal(t, tc.expectRes, asSlice(tc.sl))
		})
	}
}

func asSlice[T any](sl *SkipList[T]) []T {
	res := make([]T, 0, 10)
	p := sl.header.forward[0]
	for p != nil {
		res = append(res, p.val)
		p = p.forward[0]
	}
	return res
}
