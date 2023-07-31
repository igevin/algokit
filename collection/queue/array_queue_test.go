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

package queue

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewArrayQueueOf(t *testing.T) {
	testCases := []struct {
		name      string
		val       []int
		expectRes []int
	}{
		{
			name:      "add to empty queue",
			val:       []int{},
			expectRes: []int{},
		},
		{
			name:      "add to not empty queue",
			val:       []int{1, 2},
			expectRes: []int{1, 2},
		},
	}
	for _, tc := range testCases {
		q := NewArrayQueueOf(tc.val)
		require.NotNil(t, q)
		assert.Equal(t, tc.expectRes, q.l.AsSlice())
	}
}

func TestArrayQueue_Enqueue(t *testing.T) {
	testCases := []struct {
		name        string
		q           *ArrayQueue[int]
		val         int
		expectRes   []int
		expectedErr error
	}{
		{
			name:      "add to empty queue",
			q:         NewArrayQueueOf([]int{}),
			val:       1,
			expectRes: []int{1},
		},
		{
			name:      "add to not empty queue",
			q:         NewArrayQueueOf([]int{1, 2}),
			val:       1,
			expectRes: []int{1, 2, 1},
		},
	}

	for _, tc := range testCases {
		err := tc.q.Enqueue(tc.val)
		assert.Equal(t, tc.expectedErr, err)
		if err != nil {
			return
		}
		assert.Equal(t, tc.expectRes, tc.q.l.AsSlice())
	}
}

func TestArrayQueue_Dequeue(t *testing.T) {
	testCases := []struct {
		name        string
		q           *ArrayQueue[int]
		expectRes   int
		expectedErr error
	}{
		{
			name:        "dequeue from empty queue",
			q:           NewArrayQueueOf([]int{}),
			expectedErr: ErrDequeueFromEmptyQueue,
		},
		{
			name:      "dequeue from not empty queue",
			q:         NewArrayQueueOf([]int{1, 2}),
			expectRes: 1,
		},
	}

	for _, tc := range testCases {
		res, err := tc.q.Dequeue()
		assert.Equal(t, tc.expectedErr, err)
		if err != nil {
			return
		}
		assert.Equal(t, tc.expectRes, res)
	}
}
