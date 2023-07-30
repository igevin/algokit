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

package comparator

type Comparable[T any] interface {
	Compare(t Comparable[T]) int
}

type Compare[T any] func(src T, dst T) int

type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type ComparablePrime interface {
	RealNumber | string
}

func PrimeComparator[T ComparablePrime](src, dst T) int {
	if src > dst {
		return 1
	}
	if src == dst {
		return 0
	}
	return -1
}
