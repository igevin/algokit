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

type Set[T any] interface {
	Add(T) error
	AddAll([]T) error
	Contains(T) bool
	IsEmpty() bool
	Remove(T) error
	Size() int
	ToSlice() []T
	Clear()
}

type Operation[T any] interface {
	Intersect(target Set[T]) (Set[T], error)
	Diff(target Set[T]) (Set[T], error)
	Union(target Set[T]) (Set[T], error)
}
