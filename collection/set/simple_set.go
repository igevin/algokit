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

type SimpleSet[T any] struct {
}

func (s *SimpleSet[T]) Add(t T) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleSet[T]) AddAll(ts []T) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleSet[T]) Contains(t T) bool {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleSet[T]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleSet[T]) Remove(t T) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleSet[T]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleSet[T]) Intersect(target Set[T]) (Set[T], error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleSet[T]) Diff(target Set[T]) (Set[T], error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleSet[T]) Union(target Set[T]) (Set[T], error) {
	//TODO implement me
	panic("implement me")
}
