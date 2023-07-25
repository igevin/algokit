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

func TestDelete(t *testing.T) {
	testCases := []struct {
		name          string
		src           []int
		index         int
		expectedSlice []int
		expectedVal   int
		expectedErr   error
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
			name:          "first",
			src:           []int{1, 2, 3},
			index:         0,
			expectedSlice: []int{2, 3},
			expectedVal:   1,
		},
		{
			name:          "last",
			src:           []int{1, 2, 3},
			index:         2,
			expectedSlice: []int{1, 2},
			expectedVal:   3,
		},
		{
			name:          "normal",
			src:           []int{1, 2, 3},
			index:         1,
			expectedSlice: []int{1, 3},
			expectedVal:   2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.src, tc.index)
			require.Equal(t, tc.expectedErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.expectedSlice, res)
			assert.Equal(t, tc.expectedVal, val)
		})
	}
}
