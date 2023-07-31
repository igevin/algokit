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

package queue

import "github.com/igevin/algokit/collection/list"

type ArrayQueue[T any] struct {
	l list.ArrayList[T]
}

func NewArrayQueueOf[T any](ts []T) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		l: *(list.NewArrayListOf[T](ts)),
	}
}

func (a *ArrayQueue[T]) Enqueue(t T) error {
	return a.l.Append(t)
}

func (a *ArrayQueue[T]) Dequeue() (T, error) {
	return a.l.Delete(0)
}
