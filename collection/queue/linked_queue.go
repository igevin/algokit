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

import (
	"errors"
	"github.com/igevin/algokit/collection/list"
)

type LinkedQueue[T any] struct {
	l list.LinkedList[T]
}

func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{
		l: *(list.NewLinkedList[T]()),
	}
}

func (l *LinkedQueue[T]) Enqueue(t T) error {
	return l.l.Append(t)
}

func (l *LinkedQueue[T]) Dequeue() (T, error) {
	t, err := l.l.Delete(0)
	if errors.Is(err, list.ErrIndexOutOfRange) {
		err = ErrDequeueFromEmptyQueue
	}
	return t, err
}
