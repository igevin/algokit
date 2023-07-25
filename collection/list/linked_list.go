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

import "math"

type node[T any] struct {
	val  T
	next *node[T]
	prev *node[T]
}

func newNode[T any](val T) *node[T] {
	return &node[T]{
		val: val,
	}
}

func (n *node[T]) append(m *node[T]) error {
	n.next = m
	m.prev = n
	return nil
}

type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func NewLinkedList[T any]() *LinkedList[T] {
	res := &LinkedList[T]{
		head:   newNode[T](nil),
		tail:   newNode[T](nil),
		length: 0,
	}
	res.head.next = res.tail
	res.tail.prev = res.head
	return res
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	if n, err := l.getNode(index); err != nil {
		var t T
		return t, err
	} else {
		return n.val, nil
	}
}

func (l *LinkedList[T]) getNode(index int) (n *node[T], err error) {
	if index < 0 || index >= l.Len() {
		err = ErrIndexOutOfRange
		return
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.next
	}
	return cur, nil
}

func (l *LinkedList[T]) Append(ts ...T) error {
	end := l.tail.prev
	for _, t := range ts {
		n := newNode(t)
		err := end.append(n)
		if err != nil {
			return err
		}
		l.length++
		end = n
	}
	return end.append(l.tail)
}

func (l *LinkedList[T]) Add(index int, t T) error {
	cur, err := l.getNode(index)
	if err != nil {
		return err
	}
	n := newNode(t)
	prev := cur.prev
	prev.next = n
	n.prev = prev
	n.next = cur
	l.length++
	return nil
}

func (l *LinkedList[T]) Set(index int, t T) error {
	cur, err := l.getNode(index)
	if err != nil {
		return err
	}
	cur.val = t
	return nil
}

func (l *LinkedList[T]) Delete(index int) (T, error) {
	cur, err := l.getNode(index)
	if err != nil {
		return cur.val, err
	}
	prev := cur.prev
	next := cur.next
	prev.next = next
	next.prev = prev
	l.length--
	return cur.val, nil
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) Cap() int {
	return math.MaxInt
}

func (l *LinkedList[T]) AsSlice() []T {
	t := make([]T, 0, l.Len())
	for cur := l.head.next; cur != l.tail; cur = cur.next {
		t = append(t, cur.val)
	}
	return t
}
