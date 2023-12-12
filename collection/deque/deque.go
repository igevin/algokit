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

package deque

import "github.com/igevin/algokit/collection/list"

type baseDeque[T any] struct {
	l list.List[T]
}

func newBaseDeque[T any](l list.List[T]) baseDeque[T] {
	return baseDeque[T]{
		l: l,
	}
}

func (q *baseDeque[T]) AddFirst(t T) error {
	return q.l.Add(0, t)
}

func (q *baseDeque[T]) AddLast(t T) error {
	return q.l.Append(t)
}

func (q *baseDeque[T]) RemoveFirst() (T, error) {
	return q.l.Delete(0)
}

func (q *baseDeque[T]) RemoveLast() (T, error) {
	i := q.l.Len() - 1
	return q.l.Delete(i)
}
