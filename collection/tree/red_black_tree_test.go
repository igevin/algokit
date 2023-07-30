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

package tree

import (
	"github.com/igevin/algokit/comparator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRBTree(t *testing.T) {
	testCases := []struct {
		name      string
		compare   comparator.Compare[int]
		expectRes bool
	}{
		{
			name:      "int",
			compare:   comparator.PrimeComparator[int],
			expectRes: true,
		},
		{
			name:      "nil",
			compare:   nil,
			expectRes: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, string](tc.compare)
			assert.Equal(t, tc.expectRes, IsRedBlackTree[int, string](redBlackTree.root))
		})
	}
}

func TestIsRedBlackTree(t *testing.T) {
	testCases := []struct {
		name string
		node *rbNode[int, string]
		want bool
	}{
		{
			name: "nil",
			node: nil,
			want: true,
		},
		{
			name: "node-nil",
			node: nil,
			want: true,
		},
		{
			name: "root",
			node: &rbNode[int, string]{left: nil, right: nil, color: Black},
			want: true,
		},
		{
			name: "root",
			node: &rbNode[int, string]{left: nil, right: nil, color: Red},
			want: false,
		},
		//			 root(b)
		//			/
		//		   a(b)
		{
			name: "root-oneChild",
			node: &rbNode[int, string]{
				left: &rbNode[int, string]{
					right: nil,
					left:  nil,
					color: Red,
				},
				right: nil,
				color: Black,
			},
			want: true,
		},
		//			 root(b)
		//			/	    \
		//		   a(r)	     b(b)
		{
			name: "root-twoChild",
			node: &rbNode[int, string]{
				left: &rbNode[int, string]{
					right: nil,
					left:  nil,
					color: Red,
				},
				right: &rbNode[int, string]{
					right: nil,
					left:  nil,
					color: Black,
				},
				color: Black,
			},
			want: false,
		},
		//			 root(b)
		//			/	    \
		//		   a(b)	     b(b)
		//		 /  \        /    \
		//      nil  c(r)    d(r)   nil
		//           / \     / \
		//          nil nil nil nil
		{
			name: "blackNodeNotSame",
			node: &rbNode[int, string]{
				left: &rbNode[int, string]{
					right: &rbNode[int, string]{
						right: nil,
						left:  nil,
						color: Red,
					},
					left:  nil,
					color: Black,
				},
				right: &rbNode[int, string]{
					right: nil,
					left: &rbNode[int, string]{
						right: nil,
						left:  nil,
						color: Red,
					},
					color: Black,
				},
				color: Black,
			},
			want: true,
		},
		{
			name: "root-grandson",
			node: &rbNode[int, string]{
				parent: nil,
				key:    7,
				left: &rbNode[int, string]{
					key:   5,
					color: Black,
					left: &rbNode[int, string]{
						key:   4,
						color: Red,
					},
					right: &rbNode[int, string]{
						key:   6,
						color: Red,
					},
				},
				right: &rbNode[int, string]{
					key:   10,
					color: Red,
					left: &rbNode[int, string]{
						key:   9,
						color: Black,
						left: &rbNode[int, string]{
							key:   8,
							color: Red,
						},
					},
					right: &rbNode[int, string]{
						key:   12,
						color: Black,
						left: &rbNode[int, string]{
							key:   11,
							color: Red,
						},
					},
				},
				color: Black,
			},
			want: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsRedBlackTree[int](tc.node)
			assert.Equal(t, tc.want, res)

		})
	}
}

func TestRBTree_Add(t *testing.T) {
	testCases := []struct {
		name    string
		k       []int
		want    bool
		wantErr error
		size    int
		wantKey int
	}{
		{
			name: "nil",
			k:    nil,
			want: true,
			size: 0,
		},
		{
			name: "one",
			k:    []int{1},
			want: true,
			size: 1,
		},
		{
			name:    "one",
			k:       []int{1, 2},
			want:    true,
			size:    2,
			wantKey: 1,
		},
		{
			name:    "normal",
			k:       []int{1, 2, 3, 4},
			want:    true,
			size:    4,
			wantKey: 3,
		},
		{
			name:    "same",
			k:       []int{0, 0, 1, 2, 2, 3},
			want:    true,
			size:    0,
			wantErr: ErrRBTreeSameNode,
		},
		{
			name:    "disorder",
			k:       []int{1, 2, 0, 3, 5, 4},
			want:    true,
			wantErr: nil,
			size:    6,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.k); i++ {
				err := redBlackTree.Add(tc.k[i], i)
				if err != nil {
					assert.Equal(t, tc.wantErr, err)
					return
				}
			}
			res := IsRedBlackTree[int, int](redBlackTree.root)
			assert.Equal(t, tc.want, res)
			assert.Equal(t, tc.size, redBlackTree.Size())
		})
	}
}

func TestRBTree_Delete(t *testing.T) {
	testCases := []struct {
		name   string
		delKey int
		key    []int
		size   int
	}{
		{
			name:   "nil",
			delKey: 0,
			key:    nil,
			size:   0,
		},
		{
			name:   "node-empty",
			delKey: 0,
			key:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			size:   9,
		},
		{
			name:   "左右非空子节点,删除节点为黑色",
			delKey: 11,
			key:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			size:   8,
		},
		{
			name:   "左右只有一个非空子节点,删除节点为黑色",
			delKey: 11,
			key:    []int{4, 5, 6, 7, 8, 9, 11, 12},
			//want:   true,
			size: 7,
		},
		{
			name:   "左右均为空节点,删除节点为黑色",
			delKey: 12,
			key:    []int{4, 5, 6, 7, 8, 9, 12},
			size:   6,
		}, {
			name:   "左右非空子节点,删除节点为红色",
			delKey: 5,
			key:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			size:   8,
		},
		// 此状态无法构造出正确的红黑树
		// {
		//	name:   "左右只有一个非空子节点,删除节点为红色",
		//	delKey: 5,
		//	key:    []int{4, 5, 6, 7, 8, 9, 11, 12},
		// },
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.key); i++ {
				err := rbTree.Add(tc.key[i], i)
				if err != nil {
					panic(err)
				}
			}
			assert.True(t, IsRedBlackTree[int](rbTree.root))
			err := rbTree.Delete(tc.delKey)
			require.NoError(t, err)
			assert.True(t, IsRedBlackTree[int](rbTree.root))
			assert.Equal(t, tc.size, rbTree.Size())
		})
	}
}

func TestRBTree_Find(t *testing.T) {
	testCases := []struct {
		name      string
		find      int
		k         []int
		wantKey   int
		wantError error
	}{
		{
			name:      "nil",
			find:      0,
			k:         nil,
			wantError: ErrRBTreeNodeNotFound,
		},
		{
			name:      "node-empty",
			find:      0,
			k:         []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantError: ErrRBTreeNodeNotFound,
		},
		{
			name:    "find",
			find:    11,
			k:       []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 11,
		}, {
			name:    "find",
			find:    12,
			k:       []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 12,
		}, {
			name:    "find",
			find:    7,
			k:       []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 7,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.k); i++ {
				err := rbTree.Add(tc.k[i], tc.k[i])
				require.NoError(t, err)
			}
			assert.True(t, IsRedBlackTree[int](rbTree.root))
			findNode, err := rbTree.Find(tc.find)
			assert.Equal(t, tc.wantError, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.find, findNode)
		})
	}
}

func TestRBTree_deleteNode(t *testing.T) {
	testCases := []struct {
		name      string
		delKey    int
		key       []int
		want      bool
		wantError error
	}{
		{
			name:   "本身为右节点,左右非空子节点,删除节点为黑色",
			delKey: 11,
			key:    []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:   true,
		},
		{
			name:   "本身为右节点,右只有一个非空子节点,删除节点为黑色",
			delKey: 11,
			key:    []int{2, 3, 4, 5, 6, 7, 8, 9, 11, 12},
			want:   true,
		},
		{
			name:   "本身为右节点,左右均为空节点,删除节点为黑色",
			delKey: 6,
			key:    []int{2, 3, 4, 5, 6, 7, 8, 9, 11, 12},
			want:   true,
		},
		{
			name:   "本身为左节点,左右非空子节点,删除节点为黑色",
			delKey: 3,
			key:    []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:   true,
		},
		{
			name:   "本身为左节点,左右只有一个非空子节点,删除节点为黑色",
			delKey: 3,
			key:    []int{2, 3, 5, 6, 7, 8, 9, 11, 12},
			want:   true,
		},
		{
			name:   "本身为左节点,左右均为空节点,删除节点为黑色",
			delKey: 8,
			key:    []int{2, 3, 5, 6, 7, 8, 9, 11, 12},
			want:   true,
		},
		{
			name:   "本身是左节点,只有左边子节点,删除节点为黑色",
			delKey: 3,
			key:    []int{5, 3, 4, 6, 2},
			want:   true,
		},
		// name: "本身为左节点,左右只有一个非空子节点,删除节点为红色(无法正确构造)"
		{
			name:   "本身为左节点,左右均为空节点,删除节点为红色",
			delKey: 2,
			key:    []int{2, 3, 5, 6, 7, 8, 9, 11, 12},
			want:   true,
		},
		{
			name:   "删除root节点",
			delKey: 7,
			key:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:   true,
		},
		{
			name:   "删除root节点",
			delKey: 7,
			key:    []int{7},
			want:   true,
		},
		{
			name:   "root",
			delKey: 2,
			key:    []int{2, 1},
			want:   true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.key); i++ {
				err := rbTree.Add(tc.key[i], i)
				if err != nil {
					panic(err)
				}
			}
			delNode := rbTree.findNode(tc.delKey)
			if delNode == nil {
				return
			}
			rbTree.deleteNode(delNode)
			rbTree.deleteNode(delNode)
			assert.Equal(t, tc.want, IsRedBlackTree[int](rbTree.root))
		})
	}
}

func TestRBTree_findNode(t *testing.T) {
	testCases := []struct {
		name    string
		findKey int
		key     []int
		wantKey int
	}{
		{
			name:    "nil",
			findKey: 0,
			key:     nil,
			wantKey: -1,
		},
		{
			name:    "node-empty",
			findKey: 0,
			key:     []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: -1,
		},
		{
			name:    "find",
			findKey: 11,
			key:     []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 11,
		}, {
			name:    "find",
			findKey: 12,
			key:     []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 12,
		}, {
			name:    "find",
			findKey: 7,
			key:     []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 7,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.key); i++ {
				err := rbTree.Add(tc.key[i], i)
				if err != nil {
					panic(err)
				}
			}
			assert.Equal(t, true, IsRedBlackTree[int](rbTree.root))
			findNode := rbTree.findNode(tc.findKey)
			res := -1
			if findNode != nil {
				res = findNode.key
			}
			assert.Equal(t, tc.wantKey, res)
		})
	}
}

func TestRBTree_rotateLeft(t *testing.T) {
	testCases := []struct {
		name           string
		key            []int
		wantParent     int
		wantLeftChild  int
		wantRightChild int
		rotaNode       int
	}{
		{
			name:       "only-root",
			key:        []int{1},
			wantParent: 1,
			rotaNode:   1,
		},
		{
			name:           "节点有2个子节点，并且本身是右节点",
			key:            []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			rotaNode:       9,
			wantParent:     11,
			wantLeftChild:  8,
			wantRightChild: 10,
		},
		{
			name:          "节点有2个子节点，并且本身是左节点",
			key:           []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			rotaNode:      5,
			wantParent:    6,
			wantLeftChild: 4,
		},
		{
			name:          "节点有1个子节点",
			key:           []int{1, 2, 3, 4},
			rotaNode:      2,
			wantParent:    3,
			wantLeftChild: 1,
		},
		{
			name:          "节点没有子节点",
			key:           []int{1, 2, 3},
			rotaNode:      2,
			wantParent:    3,
			wantLeftChild: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.key); i++ {
				err := rbTree.Add(tc.key[i], i)
				if err != nil {
					panic(err)
				}
			}
			rotaNode := rbTree.findNode(tc.rotaNode)
			rbTree.rotateLeft(rotaNode)
			if rotaNode.getParent() != nil {
				assert.Equal(t, tc.wantParent, rotaNode.getParent().key)
				if rotaNode.getLeft() != nil {
					assert.Equal(t, tc.wantLeftChild, rotaNode.getLeft().key)
				}
				if rotaNode.getRight() != nil {
					assert.Equal(t, tc.wantRightChild, rotaNode.getRight().key)
				}
			}
		})
	}
}

func TestRBTree_rotateRight(t *testing.T) {
	testCases := []struct {
		name           string
		key            []int
		wantParent     int
		wantLeftChild  int
		wantRightChild int
		rotaNode       int
	}{
		{
			name:       "only-root",
			key:        []int{1},
			wantParent: 1,
			rotaNode:   1,
		},
		{
			name:           "节点有2个子节点，并且本身是右节点",
			key:            []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			rotaNode:       9,
			wantParent:     8,
			wantRightChild: 11,
		},
		{
			name:           "节点有2个子节点，并且本身是左节点",
			key:            []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			rotaNode:       5,
			wantParent:     4,
			wantRightChild: 6,
		},
		{
			name:           "有一个子节点",
			key:            []int{4, 5, 3, 2},
			rotaNode:       4,
			wantParent:     3,
			wantRightChild: 5,
		},
		{
			name:           "没有子节点",
			key:            []int{4, 5, 3},
			rotaNode:       4,
			wantParent:     3,
			wantRightChild: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.key); i++ {
				err := rbTree.Add(tc.key[i], i)
				if err != nil {
					panic(err)
				}
			}
			rotaNode := rbTree.findNode(tc.rotaNode)
			rbTree.rotateRight(rotaNode)
			if rotaNode.getParent() != nil {
				assert.Equal(t, tc.wantParent, rotaNode.getParent().key)
				if rotaNode.getLeft() != nil {
					assert.Equal(t, tc.wantLeftChild, rotaNode.getLeft().key)
				}
				if rotaNode.getRight() != nil {
					assert.Equal(t, tc.wantRightChild, rotaNode.getRight().key)
				}
			}
		})
	}
}

func TestRBNode_getColor(t *testing.T) {
	testCases := []struct {
		name      string
		node      *rbNode[int, int]
		wantColor color
	}{
		{
			name:      "nod-nil",
			node:      nil,
			wantColor: true,
		},
		{
			name:      "new node",
			node:      newRBNode[int, int](1, 1),
			wantColor: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantColor, tc.node.getColor())
		})
	}
}

func TestRBNode_getLeft(t *testing.T) {
	testCases := []struct {
		name     string
		node     *rbNode[int, int]
		wantNode *rbNode[int, int]
	}{
		{
			name:     "nod-nil",
			node:     nil,
			wantNode: nil,
		},
		{
			name:     "new node",
			node:     newRBNode[int, int](1, 1),
			wantNode: nil,
		},
		{
			name: "new node have left-child",
			node: &rbNode[int, int]{
				key: 2,
				left: &rbNode[int, int]{
					key: 1,
				},
			},
			wantNode: &rbNode[int, int]{
				key: 1,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantNode, tc.node.getLeft())
		})
	}
}

func TestRBNode_getRight(t *testing.T) {
	testCases := []struct {
		name     string
		node     *rbNode[int, int]
		wantNode *rbNode[int, int]
	}{
		{
			name:     "nod-nil",
			node:     nil,
			wantNode: nil,
		},
		{
			name:     "new node",
			node:     newRBNode[int, int](1, 1),
			wantNode: nil,
		},
		{
			name: "new node have right-child",
			node: &rbNode[int, int]{
				key: 1,
				right: &rbNode[int, int]{
					key: 2,
				},
			},
			wantNode: &rbNode[int, int]{
				key: 2,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantNode, tc.node.getRight())
		})
	}
}

func TestRBNode_getParent(t *testing.T) {
	testCases := []struct {
		name     string
		node     *rbNode[int, int]
		wantNode *rbNode[int, int]
	}{
		{
			name:     "nod-nil",
			node:     nil,
			wantNode: nil,
		},
		{
			name:     "new node",
			node:     newRBNode[int, int](1, 1),
			wantNode: nil,
		},
		{
			name: "new node have parent",
			node: &rbNode[int, int]{
				key: 2,
				parent: &rbNode[int, int]{
					key: 3,
				},
			},
			wantNode: &rbNode[int, int]{
				key: 3,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantNode, tc.node.getParent())
		})
	}
}

func TestRBNode_setColor(t *testing.T) {
	testCases := []struct {
		name      string
		node      *rbNode[int, int]
		color     color
		wantColor color
	}{
		{
			name:      "nod-nil",
			node:      nil,
			color:     false,
			wantColor: Black,
		},
		{
			name:      "new node",
			node:      newRBNode[int, int](1, 1),
			color:     true,
			wantColor: Black,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.node.setColor(tc.color)
			assert.Equal(t, tc.wantColor, tc.node.getColor())
		})
	}
}

func TestNewRBNode(t *testing.T) {
	testCases := []struct {
		name     string
		key      int
		value    int
		wantNode *rbNode[int, int]
	}{
		{
			name:  "new node",
			key:   1,
			value: 1,
			wantNode: &rbNode[int, int]{
				key:    1,
				value:  1,
				left:   nil,
				right:  nil,
				parent: nil,
				color:  Red,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			node := newRBNode[int, int](tc.key, tc.value)
			assert.Equal(t, tc.wantNode, node)
		})
	}
}

func TestRBNode_getBrother(t *testing.T) {
	testCases := []struct {
		name    string
		k       []int
		nodeKye int
		want    int
	}{
		{
			name: "nil",
			k:    nil,
		},
		{
			name:    "no-brother",
			nodeKye: 1,
			k:       []int{1},
		},
		{
			name:    "no-brother",
			nodeKye: 1,
			k:       []int{1, 2},
		},
		{
			name:    "have brother",
			k:       []int{1, 2, 3, 4},
			nodeKye: 1,
			want:    3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.k); i++ {
				err := redBlackTree.Add(tc.k[i], i)
				if err != nil {
					panic(err)
				}
			}
			tagNode := redBlackTree.findNode(tc.nodeKye)
			brNode := tagNode.getBrother()
			if brNode == nil {
				return
			}
			assert.Equal(t, tc.want, brNode.key)

		})
	}
}

func TestRBNode_getGrandParent(t *testing.T) {
	testCases := []struct {
		name    string
		k       []int
		nodeKye int
		want    int
	}{
		{
			name: "nil",
			k:    nil,
		},
		{
			name:    "no-grandpa",
			nodeKye: 1,
			k:       []int{1},
		},
		{
			name:    "no-grandpa",
			nodeKye: 1,
			k:       []int{1, 2},
		},
		{
			name:    "have grandpa",
			k:       []int{1, 2, 3, 4},
			nodeKye: 4,
			want:    2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.k); i++ {
				err := redBlackTree.Add(tc.k[i], i)
				if err != nil {
					panic(err)
				}
			}
			tagNode := redBlackTree.findNode(tc.nodeKye)
			brNode := tagNode.getGrandParent()
			if brNode == nil {
				return
			}
			assert.Equal(t, tc.want, brNode.key)

		})
	}
}

func TestRBNode_getUncle(t *testing.T) {
	testCases := []struct {
		name    string
		k       []int
		nodeKye int
		want    int
	}{
		{
			name: "nil",
			k:    nil,
		},
		{
			name:    "no-uncle",
			nodeKye: 1,
			k:       []int{1},
		},
		{
			name:    "no-uncle",
			nodeKye: 1,
			k:       []int{1, 2},
		},
		{
			name:    "have uncle",
			k:       []int{1, 2, 3, 4},
			nodeKye: 4,
			want:    1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.k); i++ {
				err := redBlackTree.Add(tc.k[i], i)
				if err != nil {
					panic(err)
				}
			}
			tagNode := redBlackTree.findNode(tc.nodeKye)
			brNode := tagNode.getUncle()
			if brNode == nil {
				return
			}
			assert.Equal(t, tc.want, brNode.key)

		})
	}
}

func TestRBNode_set(t *testing.T) {
	testCases := []struct {
		name     string
		node     *rbNode[int, int]
		value    int
		wantNode *rbNode[int, int]
	}{
		{
			name:     "nil",
			node:     nil,
			value:    1,
			wantNode: nil,
		},
		{
			name: "new node",
			node: &rbNode[int, int]{
				key:    1,
				value:  0,
				left:   nil,
				right:  nil,
				parent: nil,
				color:  Red,
			},
			value: 1,
			wantNode: &rbNode[int, int]{
				key:    1,
				value:  1,
				left:   nil,
				right:  nil,
				parent: nil,
				color:  Red,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.node.setNode(tc.value)
			assert.Equal(t, tc.wantNode, tc.node)
		})
	}
}

func TestRBTree_findSuccessor(t *testing.T) {
	testCases := []struct {
		name      string
		k         []int
		successor int
		wantKey   int
	}{
		{
			name:      "nil-successor",
			k:         nil,
			successor: 8,
		},
		{
			name:      "have no successor",
			k:         []int{2},
			successor: 2,
		},
		{
			name:      "have right successor",
			k:         []int{5, 4, 6, 3, 2},
			successor: 3,
			wantKey:   4,
		},
		{
			name:      "have right successor",
			k:         []int{5, 4, 7, 6, 3, 2},
			successor: 5,
			wantKey:   6,
		},
		{
			name:      "have no-right successor",
			k:         []int{5, 4, 6, 3, 2},
			successor: 4,
			wantKey:   5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.k); i++ {
				err := redBlackTree.Add(tc.k[i], i)
				if err != nil {
					return
				}
			}
			tagNode := redBlackTree.findNode(tc.successor)
			successorNode := redBlackTree.findSuccessor(tagNode)
			if successorNode == nil {
				return
			}
			assert.Equal(t, tc.wantKey, successorNode.key)
		})
	}
}

func TestRBTree_fixAddLeftBlack(t *testing.T) {
	testCases := []struct {
		name    string
		k       []int
		addNode int
		want    int
	}{
		{
			name:    "nod is right",
			k:       []int{2, 1, 3},
			addNode: 3,
			want:    2,
		},
		{
			name:    "node is left",
			k:       []int{2, 1},
			addNode: 1,
			want:    1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.k); i++ {
				err := redBlackTree.Add(tc.k[i], i)
				if err != nil {
					return
				}
			}
			node := redBlackTree.findNode(tc.addNode)
			x := redBlackTree.fixAddLeftBlack(node)
			assert.Equal(t, tc.want, x.key)
		})

	}
}

func TestRBTree_fixAddRightBlack(t *testing.T) {
	testCases := []struct {
		name    string
		k       []int
		addNode int
		want    int
	}{
		{
			name:    "nod is left",
			k:       []int{2, 1},
			addNode: 1,
			want:    2,
		},
		{
			name:    "node is right",
			k:       []int{2, 1, 3},
			addNode: 3,
			want:    3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.k); i++ {
				err := redBlackTree.Add(tc.k[i], i)
				if err != nil {
					return
				}
			}
			node := redBlackTree.findNode(tc.addNode)
			x := redBlackTree.fixAddRightBlack(node)
			assert.Equal(t, tc.want, x.key)
		})

	}
}

func TestRBTree_fixAfterDeleteLeft(t *testing.T) {
	testCases := []struct {
		name   string
		delKey int
		key    []int
		want   int
	}{
		{
			name:   "兄弟节点是红色并且兄弟左节点左侧是黑色",
			delKey: 10,
			key:    []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:   11,
		},
		{
			name:   "兄弟节点是黑色,兄弟节点左测是黑色",
			delKey: 2,
			key:    []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:   3,
		},
		{
			name:   "兄弟节点是黑色,兄弟节点左测不是黑色",
			delKey: 8,
			key:    []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:   5,
		},
		{
			name:   "兄弟节点是红色,兄弟节点左测是黑色",
			delKey: 1,
			key:    []int{2, 1, 3},
			want:   2,
		},
		{
			name:   "节点左旋之后兄弟节点是红色",
			delKey: 21,
			key:    []int{15, 20, 10, 16, 21, 8, 14, 7},
			want:   15,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.key); i++ {
				err := rbTree.Add(tc.key[i], i)
				if err != nil {
					panic(err)
				}

			}
			delNode := rbTree.findNode(tc.delKey)
			if delNode == nil {
				return
			} else {
				x := rbTree.fixAfterDeleteLeft(delNode)
				assert.Equal(t, tc.want, x.key)
			}
		})
	}
}

func TestRBTree_fixAfterDeleteRight(t *testing.T) {
	testCases := []struct {
		name   string
		delKey int
		key    []int
		want   int
	}{
		{
			name:   "兄弟节点是红色并且兄弟左节点左侧是黑色",
			delKey: 12,
			key:    []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:   11,
		},
		{
			name:   "兄弟节点是黑色,兄弟节点左测是黑色",
			delKey: 11,
			key:    []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 0},
			want:   9,
		},
		{
			name:   "兄弟节点是黑色,兄弟节点左测不是黑色",
			delKey: 4,
			key:    []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 0},
			want:   5,
		},
		{
			name:   "兄弟节点是红色,兄弟节点左测是黑色",
			delKey: 3,
			key:    []int{2, 1, 3},
			want:   2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](comparator.PrimeComparator[int])
			for i := 0; i < len(tc.key); i++ {
				err := rbTree.Add(tc.key[i], i)
				if err != nil {
					panic(err)
				}

			}
			delNode := rbTree.findNode(tc.delKey)
			if delNode == nil {
				return
			} else {
				x := rbTree.fixAfterDeleteRight(delNode)
				assert.Equal(t, tc.want, x.key)
			}
		})
	}
}

// IsRedBlackTree 检测是否满足红黑树
func IsRedBlackTree[K any, V any](root *rbNode[K, V]) bool {
	// 检测节点是否黑色
	if root.isRed() {
		return false
	}
	// count 取最左树的黑色节点作为对照
	count := 0
	num := 0
	node := root
	for node != nil {
		if node.isBlack() {
			count++
		}
		node = node.getLeft()
	}
	return nodeCheck[K](root, count, num)
}

// nodeCheck 节点检测
// 1、是否有连续的红色节点
// 2、每条路径的黑色节点是否一致
func nodeCheck[K any, V any](node *rbNode[K, V], count int, num int) bool {
	if node == nil {
		return true
	}
	if node.isRed() && node.parent.isRed() {
		return false
	}
	if node.getColor() {
		num++
	}
	if node.getLeft() == nil && node.getRight() == nil {
		if num != count {
			return false
		}
	}
	return nodeCheck(node.left, count, num) && nodeCheck(node.right, count, num)
}
