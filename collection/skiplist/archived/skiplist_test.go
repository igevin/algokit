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

package archived

import (
	"fmt"
	"github.com/igevin/algokit/comparator"
	"math/rand"
	"testing"
)

func TestInt(t *testing.T) {
	sl := New[int](comparator.PrimeComparator[int])
	if sl.Len() != 0 || sl.Front() != nil && sl.Back() != nil {
		t.Fatal()
	}

	testData := []int{1, 2, 3}

	sl.Insert(testData[0])
	if sl.Len() != 1 || sl.Front().Value != testData[0] || sl.Back().Value != testData[0] {
		t.Fatal()
	}

	sl.Insert(testData[2])
	if sl.Len() != 2 || sl.Front().Value != testData[0] || sl.Back().Value != testData[2] {
		t.Fatal()
	}

	sl.Insert(testData[1])
	if sl.Len() != 3 || sl.Front().Value != testData[0] || sl.Back().Value != testData[2] {
		t.Fatal()
	}

	sl.Insert(-999)
	sl.Insert(-888)
	sl.Insert(888)
	sl.Insert(999)
	sl.Insert(1000)

	expect := []int{-999, -888, 1, 2, 3, 888, 999, 1000}
	ret := make([]int, 0)

	for e := sl.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value)
	}
	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[i] {
			t.Fatal()
		}
	}

	e := sl.Find(2)
	if e == nil || e.Value != 2 {
		t.Fatal()
	}

	ret = make([]int, 0)
	for ; e != nil; e = e.Next() {
		ret = append(ret, e.Value)
	}
	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[i+3] {
			t.Fatal()
		}
	}

	sl.Remove(sl.Find(2))
	sl.Delete(888)
	sl.Delete(1000)

	expect = []int{-999, -888, 1, 3, 999}
	ret = make([]int, 0)

	for e := sl.Back(); e != nil; e = e.Prev() {
		ret = append(ret, e.Value)
	}

	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[len(ret)-i-1] {
			t.Fatal()
		}
	}

	if sl.Front().Value != -999 {
		t.Fatal()
	}

	sl.Remove(sl.Front())
	if sl.Front().Value != -888 || sl.Back().Value != 999 {
		t.Fatal()
	}

	sl.Remove(sl.Back())
	if sl.Front().Value != -888 || sl.Back().Value != 3 {
		t.Fatal()
	}

	if e = sl.Insert(int(2)); e.Value != 2 {
		t.Fatal()
	}
	sl.Delete(int(-888))

	//if r := sl.Delete(int(123)); r != nil {
	//	t.Fatal()
	//}

	if sl.Len() != 3 {
		t.Fatal()
	}

	sl.Insert(int(2))
	sl.Insert(int(2))
	sl.Insert(int(1))

	if e = sl.Find(int(2)); e == nil {
		t.Fatal()
	}

	expect = []int{int(2), int(2), int(2), int(3)}
	ret = make([]int, 0)
	for ; e != nil; e = e.Next() {
		ret = append(ret, e.Value)
	}
	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[i] {
			t.Fatal()
		}
	}

	sl2 := sl.Init()
	if sl.Len() != 0 || sl.Front() != nil || sl.Back() != nil ||
		sl2.Len() != 0 || sl2.Front() != nil || sl2.Back() != nil {
		t.Fatal()
	}

	// for i := 0; i < 100; i++ {
	// 	sl.Insert(int(rand.Intn(200)))
	// }
	// output(sl)
}

func BenchmarkIntInsertOrder(b *testing.B) {
	b.StopTimer()
	sl := New[int](comparator.PrimeComparator[int])
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Insert(int(i))
	}
}

func BenchmarkIntInsertRandom(b *testing.B) {
	b.StopTimer()
	sl := New[int](comparator.PrimeComparator[int])
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Insert(int(rand.Int()))
	}
}

func BenchmarkIntDeleteOrder(b *testing.B) {
	b.StopTimer()
	sl := New[int](comparator.PrimeComparator[int])
	for i := 0; i < 1000000; i++ {
		sl.Insert(int(i))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Delete(int(i))
	}
}

func BenchmarkIntDeleteRandom(b *testing.B) {
	b.StopTimer()
	sl := New[int](comparator.PrimeComparator[int])
	for i := 0; i < 1000000; i++ {
		sl.Insert(int(rand.Int()))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Delete(int(rand.Int()))
	}
}

func BenchmarkIntFindOrder(b *testing.B) {
	b.StopTimer()
	sl := New[int](comparator.PrimeComparator[int])
	for i := 0; i < 1000000; i++ {
		sl.Insert(int(i))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Find(int(i))
	}
}

func BenchmarkIntFindRandom(b *testing.B) {
	b.StopTimer()
	sl := New[int](comparator.PrimeComparator[int])
	for i := 0; i < 1000000; i++ {
		sl.Insert(int(rand.Int()))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Find(int(rand.Int()))
	}
}

func output(sl *SkipList[int]) {
	var x *Element[int]
	for i := 0; i < MaxLevel; i++ {
		fmt.Printf("LEVEL[%v]: ", i)
		count := 0
		x = sl.header.levelForwards[i].forward
		for x != nil {
			// fmt.Printf("%v -> ", x.Value)
			count++
			x = x.levelForwards[i].forward
		}
		// fmt.Println("NIL")
		fmt.Println("count==", count)
	}
}
