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
	"errors"
	"github.com/igevin/algokit/comparator"
	"math/rand"
)

const (
	maxLevel = 32 // 跳表的最大层级为 32
	branch   = 4  // 每隔几个节点，建一级索引；随机代码层数的算法中，表示概率 p = 1/4
)

var ErrNodeNotFound = errors.New("algokit: 节点未找到")

// SkipList 是一个跳表，层高为levels, 自底向上为第0层、1层…… levels-1 层
// 第0层是跳表全部数据，排成一个有序链表，之上逐层创建了索引，以加速0层链表的crud，整个跳表可以视为一个索引
// 每层均是一个有序链表，上一层的节点指向下一层的相同节点，从而整个跳表的数据，前后上下均可以联通
// header 是头指针，指向整个跳表，表示为一个Node数据结构
// node 结构中，val存储本身数据，forward 数组，存储每层上，该节点的下一个节点
// 故head的forward 数组，指向了每层的头指针
type SkipList[T any] struct {
	header  *node[T]
	levels  int
	length  int // 0层一共多少节点
	compare comparator.Compare[T]
}

func NewSkipList[T any](c comparator.Compare[T]) *SkipList[T] {
	return &SkipList[T]{
		header:  newNode[T](),
		levels:  1,
		length:  0,
		compare: c,
	}
}

func (s *SkipList[T]) Front() (T, error) {
	if s.header.forward[0] == nil {
		var t T
		return t, ErrNodeNotFound
	}
	return s.header.forward[0].val, nil
}

func (s *SkipList[T]) Len() int {
	return s.length
}

// Find 查找跳表是是否存在val
func (s *SkipList[T]) Find(val T) (T, error) {
	n := s.find(val)
	if n == nil {
		var t T
		return t, ErrNodeNotFound
	}
	return n.val, nil
}

// traverse 从第level层开始遍历跳表，找到val 的索引 p ，
// 并记录从level开始到p所在的层为止，每层上的索引位置
func (s *SkipList[T]) traverse(val T, level int) (*node[T], []*node[T]) {
	update := make([]*node[T], level)
	p := s.header
	for i := level - 1; i >= 0; i-- {
		for p.forward[i] != nil && s.compare(p.forward[i].val, val) < 0 {
			p = p.forward[i]
		}
		update[i] = p
	}
	return p, update
}

// find 查找跳表中，是否存在值为val的索引项（节点），并返回该索引
func (s *SkipList[T]) find(val T) *node[T] {
	p, _ := s.traverse(val, s.levels)
	// 由于上层的数据，下层一定有，故不管当前在第几层，直接在第0层确认到底有没有该数据即可
	if p.forward[0] != nil && s.compare(p.forward[0].val, val) == 0 {
		return p.forward[0]
	} else {
		return nil
	}
}

// Insert 插入一个新的值到跳表
func (s *SkipList[T]) Insert(val T) {
	level := randomLevel()
	n := newNode(withVal[T](val), withLevel[T](level))
	_, update := s.traverse(val, level)
	for i := 0; i < level; i++ {
		n.forward[i] = update[i].forward[i]
		update[i].forward[i] = n
	}
	if s.levels < level {
		s.levels = level
	}
	s.length++
}

func (s *SkipList[T]) Delete(val T) {
	p, update := s.traverse(val, s.levels)
	if p.forward[0] == nil || s.compare(p.forward[0].val, val) != 0 {
		return
	}
	for i := s.levels - 1; i >= 0; i-- {
		if update[i].forward[i] != nil && s.compare(update[i].forward[i].val, val) == 0 {
			update[i].forward[i] = update[i].forward[i].forward[i]
		}
	}
	for s.levels > 1 && s.header.forward[s.levels] == nil {
		s.levels--
	}
	s.length--

}

type node[T any] struct {
	val       T
	forward   []*node[T] // 存储该节点在每层索引上的后继节点
	nMaxLevel int
}

type nodeOption[T any] func(n *node[T])

func newNode[T any](opts ...nodeOption[T]) *node[T] {
	var t T
	n := &node[T]{
		val:       t,
		forward:   make([]*node[T], maxLevel),
		nMaxLevel: 0,
	}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

func withVal[T any](val T) nodeOption[T] {
	return func(n *node[T]) {
		n.val = val
	}
}

func withLevel[T any](level int) nodeOption[T] {
	return func(n *node[T]) {
		n.nMaxLevel = level
	}
}

// randomLevel 返回一个随机的层高.
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
	for level < maxLevel && (rand.Int31()&0xFFFF)%branch == 0 {
		level += 1
	}
	return level
}
