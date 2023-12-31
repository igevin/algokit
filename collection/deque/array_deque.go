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

type ArrayDequeue[T any] struct {
	l list.ArrayList[T]
}

func (q *ArrayDequeue[T]) AddFirst(t T) error {
	return q.l.Add(0, t)
}

func (q *ArrayDequeue[T]) AddLast(t T) error {
	return q.l.Append(t)
}

func (q *ArrayDequeue[T]) RemoveFirst() (T, error) {
	return q.l.Delete(0)
}

func (q *ArrayDequeue[T]) RemoveLast() (T, error) {
	i := q.l.Len() - 1
	return q.l.Delete(i)
}

func (q *ArrayDequeue[T]) PeekFirst() (T, error) {
	return q.l.Get(0)
}

func (q *ArrayDequeue[T]) PeekLast() (T, error) {
	i := q.l.Len() - 1
	return q.l.Get(i)
}
