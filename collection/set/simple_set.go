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

package set

type SimpleSet[T comparable] struct {
	m map[T]struct{}
}

type SimpleSetOption[T comparable] func(s *SimpleSet[T])

func NewSimpleSet[T comparable](opts ...SimpleSetOption[T]) *SimpleSet[T] {
	s := &SimpleSet[T]{
		m: make(map[T]struct{}, 16),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithNewSizeOption[T comparable](size int) SimpleSetOption[T] {
	return func(s *SimpleSet[T]) {
		s.m = make(map[T]struct{}, size)
	}
}

func (s *SimpleSet[T]) Add(t T) error {
	s.m[t] = struct{}{}
	return nil
}

func (s *SimpleSet[T]) AddAll(ts []T) error {
	for _, t := range ts {
		err := s.Add(t)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SimpleSet[T]) Contains(t T) bool {
	_, c := s.m[t]
	return c
}

func (s *SimpleSet[T]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *SimpleSet[T]) Remove(t T) error {
	delete(s.m, t)
	return nil
}

func (s *SimpleSet[T]) Size() int {
	return len(s.m)
}

func (s *SimpleSet[T]) ToSlice() []T {
	res := make([]T, 0, s.Size())
	for k := range s.m {
		res = append(res, k)
	}
	return res
}

func (s *SimpleSet[T]) Clear() {
	s.m = make(map[T]struct{}, len(s.m))
}

func (s *SimpleSet[T]) Intersect(target Set[T]) (Set[T], error) {
	if target == nil {
		return s, nil
	}
	res := NewSimpleSet[T](WithNewSizeOption[T](func(a, b Set[T]) int {
		size := a.Size()
		if size > b.Size() {
			size = b.Size()
		}
		return size
	}(s, target)))
	for k := range s.m {
		if target.Contains(k) {
			err := res.Add(k)
			if err != nil {
				return nil, err
			}
		}
	}
	return res, nil
}

func (s *SimpleSet[T]) Diff(target Set[T]) (Set[T], error) {
	if target == nil {
		return s, nil
	}
	res := NewSimpleSet[T](WithNewSizeOption[T](s.Size()))
	for k := range s.m {
		if !target.Contains(k) {
			err := res.Add(k)
			if err != nil {
				return nil, err
			}
		}
	}
	return res, nil
}

func (s *SimpleSet[T]) Union(target Set[T]) (Set[T], error) {
	if target == nil {
		return s, nil
	}
	res := NewSimpleSet[T](WithNewSizeOption[T](s.Size() + target.Size()))
	err := res.AddAll(s.ToSlice())
	if err != nil {
		return nil, err
	}
	err = res.AddAll(target.ToSlice())
	if err != nil {
		return nil, err
	}
	return res, nil
}
