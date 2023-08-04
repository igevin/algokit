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
	"github.com/igevin/algokit/collection/list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewListStack(t *testing.T) {
	ls := []list.List[int]{
		list.NewArrayListOf([]int{1, 2, 3}),
		list.NewLinkedList[int](),
	}
	for _, l := range ls {
		s := NewListStack[int](l)
		require.NotNil(t, s)
		require.NotNil(t, s.l)
	}

}

func TestListStack_Push(t *testing.T) {
	ls := []list.List[int]{
		list.NewArrayListOf([]int{}),
		list.NewLinkedList[int](),
	}
	for _, l := range ls {
		s := NewListStack[int](l)
		vals := []int{1, 2, 3, 4, 5}
		for _, val := range vals {
			res, err := s.Push(val)
			require.NoError(t, err)
			require.Equal(t, val, res)
		}
		require.Equal(t, vals, s.l.AsSlice())
	}
}
func TestListStack_Len(t *testing.T) {
	ls := []list.List[int]{
		list.NewArrayListOf([]int{}),
		list.NewLinkedList[int](),
	}
	for _, l := range ls {
		s := NewListStack[int](l)
		require.Equal(t, 0, s.Len())
		vals := []int{1, 2, 3, 4, 5}
		for i, val := range vals {
			_, _ = s.Push(val)
			require.Equal(t, i+1, s.Len())
		}
	}
}

func TestListStack_Pop(t *testing.T) {
	testCases := []struct {
		name      string
		stacks    []*ListStack[int]
		expectRes int
		expectErr error
	}{
		{
			name: "pop from empty stacks",
			stacks: []*ListStack[int]{
				NewListStack[int](list.NewArrayListOf([]int{})),
				NewListStack[int](list.NewLinkedList[int]()),
			},
			expectErr: ErrStackEmpty,
		},
		{
			name: "pop from not empty stacks, 1 element",
			stacks: []*ListStack[int]{
				func() *ListStack[int] {
					res := NewListStack[int](list.NewArrayListOf([]int{}))
					_, _ = res.Push(1)
					return res
				}(),
				func() *ListStack[int] {
					res := NewListStack[int](list.NewLinkedList[int]())
					_, _ = res.Push(1)
					return res
				}(),
			},
			expectRes: 1,
		},
		{
			name: "pop from not empty stacks, more elements",
			stacks: []*ListStack[int]{
				func() *ListStack[int] {
					res := NewListStack[int](list.NewArrayListOf([]int{}))
					_, _ = res.Push(1)
					_, _ = res.Push(2)
					_, _ = res.Push(3)
					return res
				}(),
				func() *ListStack[int] {
					res := NewListStack[int](list.NewLinkedList[int]())
					_, _ = res.Push(1)
					_, _ = res.Push(2)
					_, _ = res.Push(3)
					return res
				}(),
			},
			expectRes: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, stack := range tc.stacks {
				l := stack.Len()
				res, err := stack.Pop()
				assert.Equal(t, tc.expectErr, err)
				if err != nil {
					return
				}
				assert.Equal(t, tc.expectRes, res)
				assert.Equal(t, 1, l-stack.Len())
			}

		})
	}
}

func TestListStack_Peek(t *testing.T) {
	testCases := []struct {
		name      string
		stacks    []*ListStack[int]
		expectRes int
		expectErr error
	}{
		{
			name: "peek from empty stacks",
			stacks: []*ListStack[int]{
				NewListStack[int](list.NewArrayListOf([]int{})),
				NewListStack[int](list.NewLinkedList[int]()),
			},
			expectErr: ErrStackEmpty,
		},
		{
			name: "peek from not empty stacks, 1 element",
			stacks: []*ListStack[int]{
				func() *ListStack[int] {
					res := NewListStack[int](list.NewArrayListOf([]int{}))
					_, _ = res.Push(1)
					return res
				}(),
				func() *ListStack[int] {
					res := NewListStack[int](list.NewLinkedList[int]())
					_, _ = res.Push(1)
					return res
				}(),
			},
			expectRes: 1,
		},
		{
			name: "peek from not empty stacks, more elements",
			stacks: []*ListStack[int]{
				func() *ListStack[int] {
					res := NewListStack[int](list.NewArrayListOf([]int{}))
					_, _ = res.Push(1)
					_, _ = res.Push(2)
					_, _ = res.Push(3)
					return res
				}(),
				func() *ListStack[int] {
					res := NewListStack[int](list.NewLinkedList[int]())
					_, _ = res.Push(1)
					_, _ = res.Push(2)
					_, _ = res.Push(3)
					return res
				}(),
			},
			expectRes: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, stack := range tc.stacks {
				l := stack.Len()
				res, err := stack.Peek()
				assert.Equal(t, tc.expectErr, err)
				if err != nil {
					return
				}
				assert.Equal(t, tc.expectRes, res)
				assert.Equal(t, l, stack.Len())
			}

		})
	}
}
