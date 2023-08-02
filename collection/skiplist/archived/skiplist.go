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

package archived

import (
	"github.com/igevin/algokit/comparator"
	"math/rand"
)

const (
	MaxLevel = 32 // 跳表的最大层级为 32
	Branch   = 4  // 每隔几个节点，建一级索引；随机代码层数的算法中，表示概率 p = 1/4
)

type SkipList[T any] struct {
	header  *Element[T] // 头指针，其中header.levelForwards 存储了每层的头指针
	tail    *Element[T]
	update  []*Element[T]
	rank    []int
	length  int // 0层一共多少节点
	levels  int // SkipList 总的层数
	compare comparator.Compare[T]
}

// New returns an initialized skip list.
func New[T any](compare comparator.Compare[T]) *SkipList[T] {
	var t T
	return &SkipList[T]{
		header:  newElement[T](MaxLevel, t),
		tail:    nil,
		update:  make([]*Element[T], MaxLevel),
		rank:    make([]int, MaxLevel),
		length:  0,
		levels:  1,
		compare: compare,
	}
}

// Init initializes or clears skip list sl.
func (sl *SkipList[T]) Init() *SkipList[T] {
	var t T
	sl.header = newElement[T](MaxLevel, t)
	sl.tail = nil
	sl.update = make([]*Element[T], MaxLevel)
	sl.rank = make([]int, MaxLevel)
	sl.length = 0
	sl.levels = 1
	return sl
}

// Front returns the first elements of skip list sl or nil.
func (sl *SkipList[T]) Front() *Element[T] {
	return sl.header.levelForwards[0].forward
}

// Back returns the last elements of skip list sl or nil.
func (sl *SkipList[T]) Back() *Element[T] {
	return sl.tail
}

// Len returns the length of skip list sl.
func (sl *SkipList[T]) Len() int {
	return sl.length
}

// Find finds an element e that e.Value == v, and returns e or nil.
func (sl *SkipList[T]) Find(v T) *Element[T] {
	x := sl.find(v) // x.Value >= v
	if x != nil && sl.compare(v, x.Value) == 0 {
		return x
	}

	return nil
}

// find finds the first element e that e.Value >= v, and returns e or nil.
func (sl *SkipList[T]) find(v T) *Element[T] {
	x := sl.header
	for i := sl.levels - 1; i >= 0; i-- {
		for x.levelForwards[i].forward != nil && sl.compare(x.levelForwards[i].forward.Value, v) < 0 {
			x = x.levelForwards[i].forward
		}
		sl.update[i] = x
	}

	return x.levelForwards[0].forward
}

// Insert inserts v, increments sl.length, and returns a new element of wrap v.
func (sl *SkipList[T]) Insert(v T) *Element[T] {
	x := sl.header
	for i := sl.levels - 1; i >= 0; i-- {
		// store rank that is crossed to reach the insert position
		if i == sl.levels-1 {
			sl.rank[i] = 0
		} else {
			sl.rank[i] = sl.rank[i+1]
		}
		for x.levelForwards[i].forward != nil && sl.compare(x.levelForwards[i].forward.Value, v) < 0 {
			sl.rank[i] += x.levelForwards[i].span
			x = x.levelForwards[i].forward
		}
		sl.update[i] = x
	}

	// ensure that the v is unique, the re-insertion of v should never happen since the
	// caller of sl.Insert() should test in the hash table if the element is already inside or not.
	level := randomLevel()
	if level > sl.levels {
		for i := sl.levels; i < level; i++ {
			sl.rank[i] = 0
			sl.update[i] = sl.header
			sl.update[i].levelForwards[i].span = sl.length
		}
		sl.levels = level
	}

	x = newElement[T](level, v)
	for i := 0; i < level; i++ {
		x.levelForwards[i].forward = sl.update[i].levelForwards[i].forward
		sl.update[i].levelForwards[i].forward = x

		// update span covered by update[i] as x is inserted here
		x.levelForwards[i].span = sl.update[i].levelForwards[i].span - sl.rank[0] + sl.rank[i]
		sl.update[i].levelForwards[i].span = sl.rank[0] - sl.rank[i] + 1
	}

	// increment span for untouched levels
	for i := level; i < sl.levels; i++ {
		sl.update[i].levelForwards[i].span++
	}

	if sl.update[0] == sl.header {
		x.backward = nil
	} else {
		x.backward = sl.update[0]
	}
	if x.levelForwards[0].forward != nil {
		x.levelForwards[0].forward.backward = x
	} else {
		sl.tail = x
	}
	sl.length++

	return x
}

// deleteElement deletes e from its skip list, and decrements sl.length.
func (sl *SkipList[T]) deleteElement(e *Element[T], update []*Element[T]) {
	for i := 0; i < sl.levels; i++ {
		if update[i].levelForwards[i].forward == e {
			update[i].levelForwards[i].span += e.levelForwards[i].span - 1
			update[i].levelForwards[i].forward = e.levelForwards[i].forward
		} else {
			update[i].levelForwards[i].span -= 1
		}
	}

	if e.levelForwards[0].forward != nil {
		e.levelForwards[0].forward.backward = e.backward
	} else {
		sl.tail = e.backward
	}

	for sl.levels > 1 && sl.header.levelForwards[sl.levels-1].forward == nil {
		sl.levels--
	}
	sl.length--
}

// Remove removes e from sl if e is an element of skip list sl.
// It returns the element value e.Value.
func (sl *SkipList[T]) Remove(e *Element[T]) interface{} {
	x := sl.find(e.Value)                            // x.Value >= e.Value
	if x == e && sl.compare(e.Value, x.Value) >= 0 { // e.Value >= x.Value
		sl.deleteElement(x, sl.update)
		return x.Value
	}

	return nil
}

// Delete deletes an element e that e.Value == v, and returns e.Value or nil.
func (sl *SkipList[T]) Delete(v T) (T, bool) {
	x := sl.find(v)                              // x.Value >= v
	if x != nil && sl.compare(v, x.Value) >= 0 { // v >= x.Value
		sl.deleteElement(x, sl.update)
		return x.Value, true
	}
	var t T
	return t, false
}

type levelForward[T any] struct {
	forward *Element[T]
	span    int
}

type Element[T any] struct {
	Value         T
	backward      *Element[T]
	levelForwards []*levelForward[T]
}

// Next returns the next skip list element or nil.
func (e *Element[T]) Next() *Element[T] {
	return e.levelForwards[0].forward
}

// Prev returns the previous skip list element or nil.
func (e *Element[T]) Prev() *Element[T] {
	return e.backward
}

// newElement returns an initialized element.
func newElement[T any](level int, v T) *Element[T] {
	levelForwards := make([]*levelForward[T], level)
	for i := 0; i < level; i++ {
		levelForwards[i] = new(levelForward[T])
	}

	return &Element[T]{
		Value:         v,
		backward:      nil,
		levelForwards: levelForwards,
	}
}

// randomLevel returns a random levelForward.
// Skip List的论文中，随机层数的伪代码为：
// RandomLevel():
//
//	lvl = 1
//	while random() < p and lvl < MaxLevel do
//		lvl+=1
//	return lvl
//
// 其中，p的意思为：节点有第i层指针，那么第i+1层出现的概率为p
// redis 中，p = 1/4, maxLevel = 64
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
