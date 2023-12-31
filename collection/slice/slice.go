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

package slice

import "github.com/igevin/algokit/internal/slice"

func DeleteV0[T any](src []T, index int) ([]T, error) {
	if src == nil || index < 0 || index >= len(src) {
		return nil, ErrSliceWrapped(slice.ErrOutOfRange)
	}
	res := src[:index]
	res = append(res, src[index+1:]...)
	return slice.Shrink(res), nil
}

func Delete[T any](src []T, index int) ([]T, error) {
	res, _, err := slice.Delete(src, index)
	return res, ErrSliceWrapped(err)
}

func Index[T comparable](src []T, target T) int {
	if src == nil {
		return -1
	}
	for i, el := range src {
		if el == target {
			return i
		}
	}
	return -1
}

func LastIndexOf[T comparable](src []T, target T) int {
	if src == nil {
		return -1
	}
	for i := len(src) - 1; i >= 0; i-- {
		if src[i] == target {
			return i
		}
	}
	return -1
}

func IndexAll[T comparable](src []T, target T) []int {
	if src == nil {
		return make([]int, 0, 0)
	}
	res := make([]int, 0, len(src))
	for i, el := range src {
		if el == target {
			res = append(res, i)
		}
	}
	return res
}
