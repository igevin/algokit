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

package list

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewNode(t *testing.T) {
	n := newNode(1)
	assert.NotNil(t, n)
	assert.Equal(t, 1, n.val)
	assert.Nil(t, n.prev)
	assert.Nil(t, n.next)
}

func TestNewLinkedList(t *testing.T) {
	res := NewLinkedList[int]()
	assert.NotNil(t, res)
	assert.NotNil(t, res.head)
	assert.NotNil(t, res.tail)
	assert.Equal(t, res.tail, res.head.next)
	assert.Equal(t, res.head, res.tail.prev)
	assert.Equal(t, 0, res.Len())
}

func TestLinkedList_Add(t *testing.T) {
	testCases := []struct {
		name      string
		l         *LinkedList[int]
		index     int
		val       int
		expectErr error
	}{
		{
			name:      "empty list, out of range",
			l:         NewLinkedList[int](),
			index:     1,
			expectErr: ErrIndexOutOfRange,
		},
		{
			name:  "empty list, normal",
			l:     NewLinkedList[int](),
			index: 0,
			val:   1,
		},
		{
			name:      "normal list, out of range",
			l:         newLinkedListOf([]int{1, 2, 3, 4}),
			index:     4,
			expectErr: ErrIndexOutOfRange,
		},
		{
			name:  "normal list, add first",
			l:     newLinkedListOf([]int{1, 2, 3, 4}),
			index: 0,
			val:   0,
		},
		{
			name:  "normal list, add normal",
			l:     newLinkedListOf([]int{1, 2, 3, 4}),
			index: 2,
			val:   0,
		},
		{
			name:  "normal list, add last",
			l:     newLinkedListOf([]int{1, 2, 3, 4}),
			index: 3,
			val:   0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			length := tc.l.length
			err := tc.l.Add(tc.index, tc.val)
			assert.Equal(t, tc.expectErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, 1, tc.l.Len()-length)
			res, err := tc.l.Get(tc.index)
			require.NoError(t, err)
			assert.Equal(t, tc.val, res)
		})
	}
}

func TestLinkedList_Append(t *testing.T) {

}

func TestLinkedList_AsSlice(t *testing.T) {

}

func TestLinkedList_Cap(t *testing.T) {

}

func TestLinkedList_Delete(t *testing.T) {

}

func TestLinkedList_Get(t *testing.T) {

}

func TestLinkedList_Len(t *testing.T) {

}

func TestLinkedList_Set(t *testing.T) {

}

func newLinkedListOf[T any](ts []T) *LinkedList[T] {
	res := NewLinkedList[T]()
	_ = res.Append(ts...)
	return res
}
