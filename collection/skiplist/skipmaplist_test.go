package skiplist

import (
	"testing"

	"github.com/igevin/algokit/comparator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkipMapList_Get(t *testing.T) {
	testCases := []struct {
		name    string
		sml     *SkipMapList[string, int]
		key     string
		wantVal int
		wantOk  bool
	}{
		{
			name:   "empty list",
			sml:    NewSkipMapList[string, int](comparator.PrimeComparator[string]),
			key:    "a",
			wantOk: false,
		},
		{
			name: "get existing key",
			sml: func() *SkipMapList[string, int] {
				sml := NewSkipMapList[string, int](comparator.PrimeComparator[string])
				sml.Put("apple", 10)
				sml.Put("banana", 20)
				return sml
			}(),
			key:     "apple",
			wantVal: 10,
			wantOk:  true,
		},
		{
			name: "get non-existing key",
			sml: func() *SkipMapList[string, int] {
				sml := NewSkipMapList[string, int](comparator.PrimeComparator[string])
				sml.Put("apple", 10)
				sml.Put("banana", 20)
				return sml
			}(),
			key:    "cherry",
			wantOk: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, ok := tc.sml.Get(tc.key)
			assert.Equal(t, tc.wantOk, ok)
			if ok {
				assert.Equal(t, tc.wantVal, val)
			}
		})
	}
}

func TestSkipMapList_Put(t *testing.T) {
	testCases := []struct {
		name    string
		sml     *SkipMapList[string, int]
		key     string
		val     int
		wantLen int
	}{
		{
			name:    "put to empty",
			sml:     NewSkipMapList[string, int](comparator.PrimeComparator[string]),
			key:     "apple",
			val:     10,
			wantLen: 1,
		},
		{
			name: "put new key",
			sml: func() *SkipMapList[string, int] {
				sml := NewSkipMapList[string, int](comparator.PrimeComparator[string])
				sml.Put("apple", 10)
				return sml
			}(),
			key:     "banana",
			val:     20,
			wantLen: 2,
		},
		{
			name: "update existing key",
			sml: func() *SkipMapList[string, int] {
				sml := NewSkipMapList[string, int](comparator.PrimeComparator[string])
				sml.Put("apple", 10)
				return sml
			}(),
			key:     "apple",
			val:     15,
			wantLen: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.sml.Put(tc.key, tc.val)
			assert.Equal(t, tc.wantLen, tc.sml.Len())
			val, ok := tc.sml.Get(tc.key)
			require.True(t, ok)
			assert.Equal(t, tc.val, val)
		})
	}
}

func TestSkipMapList_Delete(t *testing.T) {
	testCases := []struct {
		name    string
		sml     *SkipMapList[int, string]
		key     int
		wantLen int
	}{
		{
			name:    "delete from empty",
			sml:     NewSkipMapList[int, string](comparator.PrimeComparator[int]),
			key:     1,
			wantLen: 0,
		},
		{
			name: "delete existing",
			sml: func() *SkipMapList[int, string] {
				sml := NewSkipMapList[int, string](comparator.PrimeComparator[int])
				sml.Put(1, "one")
				sml.Put(2, "two")
				return sml
			}(),
			key:     1,
			wantLen: 1,
		},
		{
			name: "delete non-existing",
			sml: func() *SkipMapList[int, string] {
				sml := NewSkipMapList[int, string](comparator.PrimeComparator[int])
				sml.Put(1, "one")
				return sml
			}(),
			key:     2,
			wantLen: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.sml.Delete(tc.key)
			assert.Equal(t, tc.wantLen, tc.sml.Len())
			_, ok := tc.sml.Get(tc.key)
			assert.False(t, ok)
		})
	}
}

func TestSkipMapList_Len(t *testing.T) {
	sml := NewSkipMapList[string, bool](comparator.PrimeComparator[string])
	assert.Equal(t, 0, sml.Len())

	sml.Put("a", true)
	assert.Equal(t, 1, sml.Len())

	sml.Put("b", true)
	assert.Equal(t, 2, sml.Len())

	sml.Delete("a")
	assert.Equal(t, 1, sml.Len())

	sml.Delete("b")
	assert.Equal(t, 0, sml.Len())
}