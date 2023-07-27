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

// todo 初步框架，有待调整

type Node[K comparable, V any] struct {
	Key    K
	Value  V
	Red    bool
	Left   *Node[K, V]
	Right  *Node[K, V]
	Parent *Node[K, V]
}

type RedBlackTree[K comparable, V any] struct {
	Root *Node[K, V]
}

func NewNode[K comparable, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{Key: key, Value: value, Red: true}
}

func NewRedBlackTree[K comparable, V any]() *RedBlackTree[K, V] {
	return &RedBlackTree[K, V]{}
}

func (t *RedBlackTree[K, V]) Insert(key K, value V) error {
	newNode := NewNode(key, value)
	if t.Root == nil {
		t.Root = newNode
	} else {
		t.insertRec(t.Root, newNode)
	}
	t.Root.Red = false
	return nil
}

func (t *RedBlackTree[K, V]) insertRec(root, newNode *Node[K, V]) {
	if root == nil {
		return
	}
	if lessThan(newNode.Key, root.Key) {
		if root.Left == nil {
			root.Left = newNode
			newNode.Parent = root
		} else {
			t.insertRec(root.Left, newNode)
		}
	} else if lessThan(root.Key, newNode.Key) {
		if root.Right == nil {
			root.Right = newNode
			newNode.Parent = root
		} else {
			t.insertRec(root.Right, newNode)
		}
	}
	t.fixViolation(newNode)
}

func (t *RedBlackTree[K, V]) fixViolation(node *Node[K, V]) {
	parentNode := node.Parent
	if parentNode == nil {
		node.Red = false
	} else if parentNode.Red {
		grandParentNode := parentNode.Parent
		if grandParentNode == nil {
			return
		}
		if node == parentNode.Right && parentNode == grandParentNode.Left {
			t.rotateLeft(parentNode)
			node = parentNode
			parentNode = node.Parent
		} else if node == parentNode.Left && parentNode == grandParentNode.Right {
			t.rotateRight(parentNode)
			node = parentNode
			parentNode = node.Parent
		}
		if parentNode == grandParentNode.Left {
			t.rotateRight(grandParentNode)
		} else {
			t.rotateLeft(grandParentNode)
		}
		parentNode.Red = false
		grandParentNode.Red = true
		if grandParentNode == t.Root {
			grandParentNode.Red = false
		}
	}
}

func lessThan[K comparable](a, b K) bool {
	panic("implement me")
}

func (t *RedBlackTree[K, V]) rotateLeft(node *Node[K, V]) {
	rightNode := node.Right
	node.Right = rightNode.Left
	if node.Right != nil {
		node.Right.Parent = node
	}
	rightNode.Parent = node.Parent
	if node.Parent == nil {
		t.Root = rightNode
	} else if node == node.Parent.Left {
		node.Parent.Left = rightNode
	} else {
		node.Parent.Right = rightNode
	}
	rightNode.Left = node
	node.Parent = rightNode
}

func (t *RedBlackTree[K, V]) rotateRight(node *Node[K, V]) {
	leftNode := node.Left
	node.Left = leftNode.Right
	if node.Left != nil {
		node.Left.Parent = node
	}
	leftNode.Parent = node.Parent
	if node.Parent == nil {
		t.Root = leftNode
	} else if node == node.Parent.Right {
		node.Parent.Right = leftNode
	} else {
		node.Parent.Left = leftNode
	}
	leftNode.Right = node
	node.Parent = leftNode
}

func (t *RedBlackTree[K, V]) Delete(key K) error {
	node := t.searchRec(t.Root, key)
	if node != nil {
		t.deleteNode(node)
	}
	return nil
}

func (t *RedBlackTree[K, V]) searchRec(root *Node[K, V], key K) *Node[K, V] {
	if root == nil || root.Key == key {
		return root
	}
	if lessThan(key, root.Key) {
		return t.searchRec(root.Left, key)
	} else {
		return t.searchRec(root.Right, key)
	}
}

func (t *RedBlackTree[K, V]) deleteNode(node *Node[K, V]) {
	y := node
	color := y.Red
	var x, xParent *Node[K, V]

	if node.Left == nil {
		x = node.Right
		t.transplant(node, node.Right)
	} else if node.Right == nil {
		x = node.Left
		t.transplant(node, node.Left)
	} else {
		y = t.minimum(node.Right)
		color = y.Red
		x = y.Right

		if y.Parent == node {
			xParent = y
		} else {
			t.transplant(y, y.Right)
			y.Right = node.Right
			y.Right.Parent = y
		}

		t.transplant(node, y)
		y.Left = node.Left
		y.Left.Parent = y
		y.Red = node.Red
	}

	if !color {
		t.fixDeleteViolation(x, xParent)
	}
}

func (t *RedBlackTree[K, V]) minimum(node *Node[K, V]) *Node[K, V] {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (t *RedBlackTree[K, V]) transplant(u, v *Node[K, V]) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

func (t *RedBlackTree[K, V]) fixDeleteViolation(node, parentNode *Node[K, V]) {
	for node != t.Root && (node == nil || !node.Red) {
		if parentNode.Left == node {
			siblingNode := parentNode.Right

			if siblingNode.Red {
				siblingNode.Red = false
				parentNode.Red = true
				t.rotateLeft(parentNode)
				siblingNode = parentNode.Right
			}

			if (siblingNode.Left == nil || !siblingNode.Left.Red) && (siblingNode.Right == nil || !siblingNode.Right.Red) {
				siblingNode.Red = true
				node = parentNode
				parentNode = node.Parent
			} else {
				if siblingNode.Right == nil || !siblingNode.Right.Red {
					siblingNode.Left.Red = false
					siblingNode.Red = true
					t.rotateRight(siblingNode)
					siblingNode = parentNode.Right
				}
				siblingNode.Red = parentNode.Red
				parentNode.Red = false
				if siblingNode.Right != nil {
					siblingNode.Right.Red = false
				}
				t.rotateLeft(parentNode)
				node = t.Root
			}
		} else {
			siblingNode := parentNode.Left

			if siblingNode.Red {
				siblingNode.Red = false
				parentNode.Red = true
				t.rotateRight(parentNode)
				siblingNode = parentNode.Left
			}

			if (siblingNode.Right == nil || !siblingNode.Right.Red) && (siblingNode.Left == nil || !siblingNode.Left.Red) {
				siblingNode.Red = true
				node = parentNode
				parentNode = node.Parent
			} else {
				if siblingNode.Left == nil || !siblingNode.Left.Red {
					siblingNode.Right.Red = false
					siblingNode.Red = true
					t.rotateLeft(siblingNode)
					siblingNode = parentNode.Left
				}
				siblingNode.Red = parentNode.Red
				parentNode.Red = false
				if siblingNode.Left != nil {
					siblingNode.Left.Red = false
				}
				t.rotateRight(parentNode)
				node = t.Root
			}
		}
	}

	if node != nil {
		node.Red = false
	}
}

func (t *RedBlackTree[K, V]) Search(key K) (V, bool) {
	node := t.searchRec(t.Root, key)
	if node != nil {
		return node.Value, true
	}
	var v V
	return v, false
}

func (t *RedBlackTree[K, V]) Update(key K, value V) {
	node := t.searchRec(t.Root, key)
	if node != nil {
		node.Value = value
	}
}
