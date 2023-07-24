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

import "github.com/igevin/algokit/collection/set"

func Intersect[T comparable](a, b []T) ([]T, error) {
	if a == nil || b == nil {
		return nil, nil
	}
	a1 := set.NewSimpleSet[T](set.WithSliceOption[T](a))
	b1 := set.NewSimpleSet[T](set.WithSliceOption[T](b))
	if res, err := a1.Intersect(b1); err != nil {
		return nil, err
	} else {
		return res.ToSlice(), nil
	}
}

func Union[T comparable](a, b []T) ([]T, error) {
	if a == nil || b == nil {
		return nil, nil
	}
	a1 := set.NewSimpleSet[T](set.WithSliceOption[T](a))
	b1 := set.NewSimpleSet[T](set.WithSliceOption[T](b))
	if res, err := a1.Union(b1); err != nil {
		return nil, err
	} else {
		return res.ToSlice(), nil
	}
}

func Diff[T comparable](a, b []T) ([]T, error) {
	if a == nil || b == nil {
		return nil, nil
	}
	a1 := set.NewSimpleSet[T](set.WithSliceOption[T](a))
	b1 := set.NewSimpleSet[T](set.WithSliceOption[T](b))
	if res, err := a1.Diff(b1); err != nil {
		return nil, err
	} else {
		return res.ToSlice(), nil
	}
}
