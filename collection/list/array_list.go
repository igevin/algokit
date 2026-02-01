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

import "github.com/igevin/algokit/internal/slice"

type ArrayList[T any] struct {
	data []T
}

func NewArrayListOf[T any](ts []T) *ArrayList[T] {
	return &ArrayList[T]{
		data: ts,
	}
}

func (a *ArrayList[T]) Get(index int) (t T, e error) {
	l := a.Len()
	if index < 0 || index >= l {
		return t, NewErrIndexOutOfRange(l, index)
	}
	return a.data[index], e
}

func (a *ArrayList[T]) Append(ts ...T) error {
	a.data = append(a.data, ts...)
	return nil
}

func (a *ArrayList[T]) Add(index int, t T) error {
	if index < 0 || index > len(a.data) {
		return NewErrIndexOutOfRange(a.Len(), index)
	}
	// 先扩容（如果需要），确保有足够的空间
	a.data = append(a.data, t)
	// 将 index 位置及之后的元素向后移动一位
	copy(a.data[index+1:], a.data[index:])
	// 将新元素放入 index 位置
	a.data[index] = t
	return nil
}

func (a *ArrayList[T]) Set(index int, t T) error {
	length := len(a.data)
	if index >= length || index < 0 {
		return NewErrIndexOutOfRange(length, index)
	}
	a.data[index] = t
	return nil
}

func (a *ArrayList[T]) Delete(index int) (T, error) {
	data, t, err := slice.Delete(a.data, index)
	if err != nil {
		return t, err
	}
	a.data = data
	return t, nil
}

func (a *ArrayList[T]) Len() int {
	return len(a.data)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.data)
}

func (a *ArrayList[T]) AsSlice() []T {
	res := make([]T, len(a.data))
	copy(res, a.data)
	return res
}
