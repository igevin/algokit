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

package mapx

type Hashable interface {
	// Code 返回该元素的哈希值
	// 注意：哈希值应该尽可能的均匀以避免冲突
	Code() uint64
	// Equals 比较两个元素是否相等。如果返回 true，那么我们会认为两个键是一样的。
	Equals(key any) bool
}

type HashMap[K Hashable, V any] struct {
}

func (h *HashMap[K, V]) Keys() []K {
	//TODO implement me
	panic("implement me")
}

func (h *HashMap[K, V]) Values() []V {
	//TODO implement me
	panic("implement me")
}

func (h *HashMap[K, V]) KeysValues() []V {
	//TODO implement me
	panic("implement me")
}

func (h *HashMap[K, V]) Get(key K) (V, bool) {
	//TODO implement me
	panic("implement me")
}

func (h *HashMap[K, V]) GetOrDefault(key K, value V) V {
	//TODO implement me
	panic("implement me")
}

func (h *HashMap[K, V]) Put(key K, value V) (V, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HashMap[K, V]) PutIfAbsent(key K, value V) (V, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HashMap[K, V]) Delete(key K) (V, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HashMap[K, V]) DeleteIf(key K, value V) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HashMap[K, V]) Len() int {
	//TODO implement me
	panic("implement me")
}
