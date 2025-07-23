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
	"github.com/igevin/algokit/comparator"
)

// skipMapNode represents a node in the SkipMapList.
// @internal
type skipMapNode[K, V any] struct {
	key     K
	val     V
	forward []*skipMapNode[K, V]
}

// SkipMapList is a map-like data structure implemented with a skip list.
// It stores key-value pairs, sorted by key.
// Level 0 contains all key-value pairs, while upper levels only store keys as indices for fast searching.
// For simplicity, this implementation uses the same node structure for all levels,
// but logically the value is only relevant at level 0.
type SkipMapList[K, V any] struct {
	header  *skipMapNode[K, V]
	levels  int
	length  int
	compare comparator.Compare[K]
}

// NewSkipMapList creates and initializes a new SkipMapList.
// It requires a comparator function for the key type K.
func NewSkipMapList[K, V any](c comparator.Compare[K]) *SkipMapList[K, V] {
	var k K
	var v V
	return &SkipMapList[K, V]{
		header:  &skipMapNode[K, V]{key: k, val: v, forward: make([]*skipMapNode[K, V], maxLevel)},
		levels:  1,
		length:  0,
		compare: c,
	}
}

// Get retrieves the value associated with the given key.
// It returns the value and true if the key is found, otherwise it returns the zero value for V and false.
func (s *SkipMapList[K, V]) Get(key K) (V, bool) {
	p := s.header
	for i := s.levels - 1; i >= 0; i-- {
		for p.forward[i] != nil && s.compare(p.forward[i].key, key) < 0 {
			p = p.forward[i]
		}
	}
	p = p.forward[0]
	if p != nil && s.compare(p.key, key) == 0 {
		return p.val, true
	}
	var v V
	return v, false
}

// Put inserts or updates a key-value pair in the skip list.
// If the key already exists, its value is updated.
func (s *SkipMapList[K, V]) Put(key K, val V) {
	update := make([]*skipMapNode[K, V], maxLevel)
	p := s.header
	for i := s.levels - 1; i >= 0; i-- {
		for p.forward[i] != nil && s.compare(p.forward[i].key, key) < 0 {
			p = p.forward[i]
		}
		update[i] = p
	}
	p = p.forward[0]

	// If key already exists, update the value
	if p != nil && s.compare(p.key, key) == 0 {
		p.val = val
		return
	}

	// If key does not exist, insert a new node
	level := randomLevel()
	if level > s.levels {
		for i := s.levels; i < level; i++ {
			update[i] = s.header
		}
		s.levels = level
	}

	newNode := &skipMapNode[K, V]{key: key, val: val, forward: make([]*skipMapNode[K, V], level)}
	for i := 0; i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
	s.length++
}

// Delete removes the key-value pair associated with the given key.
func (s *SkipMapList[K, V]) Delete(key K) {
	update := make([]*skipMapNode[K, V], maxLevel)
	p := s.header
	for i := s.levels - 1; i >= 0; i-- {
		for p.forward[i] != nil && s.compare(p.forward[i].key, key) < 0 {
			p = p.forward[i]
		}
		update[i] = p
	}
	p = p.forward[0]

	if p == nil || s.compare(p.key, key) != 0 {
		return // Key not found
	}

	for i := 0; i < s.levels; i++ {
		if update[i].forward[i] != p {
			break
		}
		update[i].forward[i] = p.forward[i]
	}

	// Adjust levels if the top levels become empty
	for s.levels > 1 && s.header.forward[s.levels-1] == nil {
		s.levels--
	}
	s.length--
}

// Len returns the number of key-value pairs in the skip list.
func (s *SkipMapList[K, V]) Len() int {
	return s.length
}
