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

func Map[S any, T any](src []S, m func(S) T) []T {
	if src == nil {
		return nil
	}
	res := make([]T, len(src))
	for i, el := range src {
		res[i] = m(el)
	}
	return res
}

func Filter[T any](src []T, filter func(T) bool) []T {
	if src == nil {
		return nil
	}
	res := make([]T, 0, len(src))
	for _, el := range src {
		if filter(el) {
			res = append(res, el)
		}
	}
	return res
}

func Reduce[T any](src []T, f func(a, b T) T) T {
	if src == nil || len(src) == 0 {
		var t T
		return t
	}
	res := src[0]
	for i := 1; i < len(src); i++ {
		res = f(res, src[i])
	}
	return res
}
