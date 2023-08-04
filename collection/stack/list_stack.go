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

import "github.com/igevin/algokit/collection/list"

type ListStack[T any] struct {
	l list.List[T]
}

func NewListStack[T any](l list.List[T]) *ListStack[T] {
	return &ListStack[T]{
		l: l,
	}
}

func (l *ListStack[T]) Pop() (T, error) {
	s := l.Len()
	if s == 0 {
		var t T
		return t, ErrStackEmpty
	}
	return l.l.Delete(s - 1)
}

func (l *ListStack[T]) Push(t T) (T, error) {
	err := l.l.Append(t)
	return t, err
}

func (l *ListStack[T]) Peek() (T, error) {
	s := l.Len()
	if s == 0 {
		var t T
		return t, ErrStackEmpty
	}
	return l.l.Get(s - 1)
}

func (l *ListStack[T]) Len() int {
	return l.l.Len()
}
