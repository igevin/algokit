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
	"math/rand"
)

const MaxLevel = 32
const Branch = 4

type skipListLevel[T any] struct {
	forward *Element[T]
	span    int
}

type Element[T any] struct {
	Value    T
	backward *Element[T]
	level    []*skipListLevel[T]
}

// Next returns the next skiplist element or nil.
func (e *Element[T]) Next() *Element[T] {
	return e.level[0].forward
}

// Prev returns the previous skiplist element or nil.
func (e *Element[T]) Prev() *Element[T] {
	return e.backward
}

// newElement returns an initialized element.
func newElement[T any](level int, v T) *Element[T] {
	slLevels := make([]*skipListLevel[T], level)
	for i := 0; i < level; i++ {
		slLevels[i] = new(skipListLevel[T])
	}

	return &Element[T]{
		Value:    v,
		backward: nil,
		level:    slLevels,
	}
}

// randomLevel returns a random level.
func randomLevel() int {
	level := 1
	for (rand.Int31()&0xFFFF)%Branch == 0 {
		level += 1
	}

	if level < MaxLevel {
		return level
	} else {
		return MaxLevel
	}
}
