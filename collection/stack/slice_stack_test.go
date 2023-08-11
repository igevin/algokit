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

package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSliceStack(t *testing.T) {
	s := NewSliceStack[int]()
	require.NotNil(t, s)
	require.NotNil(t, s.s)
}

func TestSliceStack_Push(t *testing.T) {
	s := NewSliceStack[int]()
	vals := []int{1, 2, 3, 4, 5}
	for _, val := range vals {
		res, err := s.Push(val)
		require.NoError(t, err)
		require.Equal(t, val, res)
	}
	require.Equal(t, vals, s.s)
}
func TestSliceStack_Len(t *testing.T) {
	s := NewSliceStack[int]()
	require.Equal(t, 0, s.Len())
	vals := []int{1, 2, 3, 4, 5}
	for i, val := range vals {
		_, _ = s.Push(val)
		require.Equal(t, i+1, s.Len())
	}
}

func TestSliceStack_Pop(t *testing.T) {
	testCases := []struct {
		name      string
		stack     *SliceStack[int]
		expectRes int
		expectErr error
	}{
		{
			name:      "pop from empty stack",
			stack:     NewSliceStack[int](),
			expectErr: ErrStackEmpty,
		},
		{
			name: "pop from not empty stack, 1 element",
			stack: func() *SliceStack[int] {
				s := NewSliceStack[int]()
				_, _ = s.Push(1)
				return s
			}(),
			expectRes: 1,
		},
		{
			name: "pop from not empty stack, more elements",
			stack: func() *SliceStack[int] {
				s := NewSliceStack[int]()
				_, _ = s.Push(1)
				_, _ = s.Push(2)
				_, _ = s.Push(3)
				return s
			}(),
			expectRes: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := tc.stack.Len()
			res, err := tc.stack.Pop()
			assert.Equal(t, tc.expectErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.expectRes, res)
			assert.Equal(t, 1, l-tc.stack.Len())
		})
	}
}

func TestSliceStack_Peek(t *testing.T) {
	testCases := []struct {
		name      string
		stack     *SliceStack[int]
		expectRes int
		expectErr error
	}{
		{
			name:      "peek from empty stack",
			stack:     NewSliceStack[int](),
			expectErr: ErrStackEmpty,
		},
		{
			name: "pop from not empty stack, 1 element",
			stack: func() *SliceStack[int] {
				s := NewSliceStack[int]()
				_, _ = s.Push(1)
				return s
			}(),
			expectRes: 1,
		},
		{
			name: "pop from not empty stack, more elements",
			stack: func() *SliceStack[int] {
				s := NewSliceStack[int]()
				_, _ = s.Push(1)
				_, _ = s.Push(2)
				_, _ = s.Push(3)
				return s
			}(),
			expectRes: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := tc.stack.Len()
			res, err := tc.stack.Peek()
			assert.Equal(t, tc.expectErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.expectRes, res)
			assert.Equal(t, l, tc.stack.Len())
		})
	}
}
