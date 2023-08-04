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

import "github.com/igevin/algokit/internal/slice"

type SliceStack[T any] struct {
	s []T
}

func NewSliceStack[T any]() *SliceStack[T] {
	s := &SliceStack[T]{
		s: make([]T, 0, 16),
	}
	return s
}

func (s *SliceStack[T]) Pop() (T, error) {
	l := len(s.s)
	if l == 0 {
		var t T
		return t, ErrStackEmpty
	}
	t := s.s[l-1]
	s.s = slice.Shrink(s.s[:l-1])
	return t, nil
}

func (s *SliceStack[T]) Push(t T) (T, error) {
	s.s = append(s.s, t)
	return t, nil
}

func (s *SliceStack[T]) Peek() (T, error) {
	l := len(s.s)
	if l == 0 {
		var t T
		return t, ErrStackEmpty
	}
	t := s.s[l-1]
	return t, nil
}

func (s *SliceStack[T]) Empty() bool {
	c := cap(s.s)
	s.s = make([]T, 0, c)
	return true
}
