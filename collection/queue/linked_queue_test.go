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

func TestNewLinkedQueue(t *testing.T) {
	q := NewLinkedQueue[int]()
	require.NotNil(t, q)
}

func TestLinkedQueue_Enqueue(t *testing.T) {
	q := NewLinkedQueue[int]()
	require.NotNil(t, q)
	vals := []int{1, 2, 3}
	for _, val := range vals {
		require.NoError(t, q.Enqueue(val))
	}

	s := q.l.AsSlice()
	assert.Equal(t, vals, s)
}

func TestLinkedQueue_Dequeue(t *testing.T) {
	testCases := []struct {
		name        string
		q           *LinkedQueue[int]
		expectRes   int
		expectedErr error
	}{
		{
			name:        "dequeue from empty queue",
			q:           newLinkedQueueOf([]int{}),
			expectedErr: ErrDequeueFromEmptyQueue,
		},
		{
			name:      "dequeue from not empty queue",
			q:         newLinkedQueueOf([]int{1, 2}),
			expectRes: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.q.Dequeue()
			assert.Equal(t, tc.expectedErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.expectRes, res)
		})

	}
}

func newLinkedQueueOf(vals []int) *LinkedQueue[int] {
	q := NewLinkedQueue[int]()
	for _, val := range vals {
		_ = q.Enqueue(val)
	}
	return q
}
