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

type ArrayList[T any] struct {
	data []T
}

func (a *ArrayList[T]) Get(index int) (t T, e error) {
	l := a.Len()
	if index < 0 || index >= l {
		return t, NewErrIndexOutOfRange(l, index)
	}
	return a.data[index], e
}

func (a *ArrayList[T]) Append(ts ...T) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Add(index int, t T) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Set(index int, t T) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Delete(index int) (T, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Cap() int {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) AsSlice() []T {
	//TODO implement me
	panic("implement me")
}
