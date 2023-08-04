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

// SimpleHashMap 是一个哈希表，采用一次性扩容
type SimpleHashMap[K Hashable, V any] struct {
}

func (s *SimpleHashMap[K, V]) Keys() []K {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleHashMap[K, V]) Values() []V {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleHashMap[K, V]) KeysValues() []V {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleHashMap[K, V]) Get(key K) (V, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleHashMap[K, V]) GetOrDefault(key K, value V) V {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleHashMap[K, V]) Put(key K, value V) (V, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleHashMap[K, V]) PutIfAbsent(key K, value V) (V, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleHashMap[K, V]) Delete(key K) (V, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleHashMap[K, V]) DeleteIf(key K, value V) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleHashMap[K, V]) Len() int {
	//TODO implement me
	panic("implement me")
}
