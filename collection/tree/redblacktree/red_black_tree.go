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

package redblacktree

import (
	"errors"

	"github.com/igevin/algokit/comparator"
)

type color bool

const (
	Red   color = false
	Black color = true
)

var (
	ErrRBTreeSameNode     = errors.New("algokit: RBTree不能添加重复节点Key")
	ErrRBTreeNodeNotFound = errors.New("algokit: RBTree不存在节点Key")
	// errRBTreeCantRepaceNil = errors.New("algokit: RBTree不能将节点替换为nil")
)

type RBTree[K any, V any] struct {
	root    *rbNode[K, V]
	compare comparator.Compare[K]
	size    int
}

func (rb *RBTree[K, V]) Size() int {
	if rb == nil || rb.root == nil {
		return 0
	}
	return rb.size
}

type rbNode[K any, V any] struct {
	color               color
	key                 K
	value               V
	left, right, parent *rbNode[K, V]
}

func (node *rbNode[K, V]) setNode(v V) {
	if node == nil {
		return
	}
	node.value = v
}

// NewRBTree 构建红黑树
func NewRBTree[K any, V any](compare comparator.Compare[K]) *RBTree[K, V] {
	return &RBTree[K, V]{
		compare: compare,
		root:    nil,
	}
}

func newRBNode[K any, V any](key K, value V) *rbNode[K, V] {
	return &rbNode[K, V]{
		key:   key,
		value: value,
		color: Red,
	}
}

// Add 增加节点
func (rb *RBTree[K, V]) Add(key K, value V) error {
	node := newRBNode(key, value)
	err := rb.addNode(node)
	if err != nil {
		return err
	}
	rb.fixAfterAdd(node)
	return nil
}

// Delete 删除节点
func (rb *RBTree[K, V]) Delete(key K) error {
	if node := rb.findNode(key); node != nil {
		rb.deleteNode(node)
	}
	return nil
}

// Find 查找节点
func (rb *RBTree[K, V]) Find(key K) (V, error) {
	if node := rb.findNode(key); node != nil {
		return node.value, nil
	}
	var v V
	return v, ErrRBTreeNodeNotFound
}

func (rb *RBTree[K, V]) Set(key K, value V) error {
	if node := rb.findNode(key); node != nil {
		node.setNode(value)
		return nil
	}
	return ErrRBTreeNodeNotFound
}

// addNode 插入新节点
func (rb *RBTree[K, V]) addNode(node *rbNode[K, V]) error {
	if rb.root == nil {
		rb.root = node
		rb.size++
		return nil
	}
	t := rb.root
	cmp := 0
	var parent *rbNode[K, V]
	for t != nil {
		parent = t
		cmp = rb.compare(node.key, t.key)
		if cmp < 0 {
			t = t.left
		} else if cmp > 0 {
			t = t.right
		} else if cmp == 0 {
			return ErrRBTreeSameNode
		}
	}
	node.parent = parent
	if cmp < 0 {
		parent.left = node
	} else {
		parent.right = node
	}
	rb.size++
	//rb.fixAfterAdd(node)
	return nil
}

// deleteNode 红黑树删除方法
// 删除分两步,第一步取出后继节点,第二部着色旋转
// 取后继节点
// case1:node左右非空子节点,通过getSuccessor获取后继节点
// case2:node左右只有一个非空子节点
// case3:node左右均为空节点
// 着色旋转
// case1:当删除节点非空且为黑色时,会违反红黑树任何路径黑节点个数相同的约束,所以需要重新平衡
// case2:当删除红色节点时,不会破坏任何约束,所以不需要平衡
func (rb *RBTree[K, V]) deleteNode(node *rbNode[K, V]) {
	// node左右非空,取后继节点
	if node.left != nil && node.right != nil {
		s := rb.findSuccessor(node)
		node.key = s.key
		node.value = s.value
		node = s
	}
	var replacement *rbNode[K, V]
	// node节点只有一个非空子节点
	if node.left != nil {
		replacement = node.left
	} else {
		replacement = node.right
	}
	if replacement != nil {
		replacement.parent = node.parent
		if node.parent == nil {
			rb.root = replacement
		} else if node == node.parent.left {
			node.parent.left = replacement
		} else {
			node.parent.right = replacement
		}
		node.left = nil
		node.right = nil
		node.parent = nil
		if node.isBlack() {
			rb.fixAfterDelete(replacement)
		}
	} else if node.parent == nil {
		// 如果node节点无父节点,说明node为root节点
		rb.root = nil
	} else {
		// node子节点均为空
		if node.isBlack() {
			rb.fixAfterDelete(node)
		}
		if node.parent != nil {
			if node == node.parent.left {
				node.parent.left = nil
			} else if node == node.parent.right {
				node.parent.right = nil
			}
			node.parent = nil
		}
	}
	rb.size--
}

// findSuccessor 寻找后继节点
// case1: node节点存在右子节点,则右子树的最小节点是node的后继节点
// case2: node节点不存在右子节点,则其第一个为左节点的祖先的父节点为node的后继节点
func (rb *RBTree[K, V]) findSuccessor(node *rbNode[K, V]) *rbNode[K, V] {
	if node == nil {
		return nil
	}
	if node.right != nil {
		p := node.right
		for p.left != nil {
			p = p.left
		}
		return p
	}
	p := node.parent
	ch := node
	for p != nil && ch == p.right {
		ch = p
		p = p.parent
	}
	return p

}

func (rb *RBTree[K, V]) findNode(key K) *rbNode[K, V] {
	node := rb.root
	for node != nil {
		cmp := rb.compare(key, node.key)
		if cmp < 0 {
			node = node.left
		} else if cmp > 0 {
			node = node.right
		} else {
			return node
		}
	}
	return nil
}

// fixAfterAdd 插入时着色旋转
// 如果是空节点、root节点、父节点是黑无需构建
// 可分为3种情况
// fixUncleRed 叔叔节点是红色右节点
// fixAddLeftBlack 叔叔节点是黑色右节点
// fixAddRightBlack 叔叔节点是黑色左节点
func (rb *RBTree[K, V]) fixAfterAdd(x *rbNode[K, V]) {
	x.color = Red
	for x != nil && x != rb.root && x.getParent().isRed() {
		uncle := x.getUncle()
		if uncle.isRed() {
			x = rb.fixUncleRed(x, uncle)
			continue
		}
		if x.getParent() == x.getGrandParent().getLeft() {
			x = rb.fixAddLeftBlack(x)
			continue
		}
		x = rb.fixAddRightBlack(x)
	}
	rb.root.setColor(Black)
}

// fixAddLeftRed 叔叔节点是红色右节点，由于不能存在连续红色节点,此时祖父节点x.getParent().getParent()必为黑。另x为红所以叔父节点需要变黑，祖父变红，此时红黑树完成
//
//							  b(b)                    b(r)
//							/		\				/		\
//						  a(r)        y(r)  ->   a(b)        y(b)
//						/   \       /  \         /   \       /  \
//		            x(r)    nil   nil  nil    x (r) nil   nil  nil
//	             	/  \                      /  \
//	            	nil nil                   nil nil
func (rb *RBTree[K, V]) fixUncleRed(x *rbNode[K, V], y *rbNode[K, V]) *rbNode[K, V] {
	x.getParent().setColor(Black)
	y.setColor(Black)
	x.getGrandParent().setColor(Red)
	x = x.getGrandParent()
	return x
}

// fixAddLeftBlack 叔叔节点是黑色右节点.x节点是父节点左节点,执行左旋，此时x节点变为原x节点的父节点a,也就是左子节点。的接着将x的父节点和爷爷节点的颜色对换。然后对爷爷节点进行右旋转,此时红黑树完成
// 如果x为左节点则跳过左旋操作
//
//							  b(b)                    b(b)                b(r)
//							/		\				/		\            /   \
//						  a(r)        y(b)  ->   a(r)        y(b)  ->  a(b)   y(b)
//						/   \       /  \         /   \       /  \      /  \    /  \
//		               nil   x (r) nil  nil      x(r) nil  nil  nil   x(r) nil nil nil
//	           		 		 /  \               /  \                  / \
//	           		 		nil nil             nil nil              nil nil
func (rb *RBTree[K, V]) fixAddLeftBlack(x *rbNode[K, V]) *rbNode[K, V] {
	if x == x.getParent().getRight() {
		x = x.getParent()
		rb.rotateLeft(x)
	}
	x.getParent().setColor(Black)
	x.getGrandParent().setColor(Red)
	rb.rotateRight(x.getGrandParent())
	return x
}

// fixAddRightBlack 叔叔节点是黑色左节点.x节点是父节点右节点,执行右旋，此时x节点变为原x节点的父节点a,也就是右子节点。接着将x的父节点和爷爷节点的颜色对换。然后对爷爷节点进行右旋转,此时红黑树完成
// 如果x为右节点则跳过右旋操作
//
//							  b(b)                    b(b)                b(r)
//							/		\				/		\            /   \
//						  y(b)       a(r)  ->   y(b)        a(r)  ->  y(b)     a(b)
//						/   \       /  \         /   \       /  \      /  \    /  \
//		               nil   nil x(r)  nil      nil nil  nil  x(r)   nil nil  nil  x(r)
//	           		 		      /  \                         /  \               /  \
//	           		 		      nil nil                    nil nil              nil nil
func (rb *RBTree[K, V]) fixAddRightBlack(x *rbNode[K, V]) *rbNode[K, V] {
	if x == x.getParent().getLeft() {
		x = x.getParent()
		rb.rotateRight(x)
	}
	x.getParent().setColor(Black)
	x.getGrandParent().setColor(Red)
	rb.rotateLeft(x.getGrandParent())
	return x
}

// fixAfterDelete 删除时着色旋转
// 根据x是节点位置分为fixAfterDeleteLeft,fixAfterDeleteRight两种情况
func (rb *RBTree[K, V]) fixAfterDelete(x *rbNode[K, V]) {
	for x != rb.root && x.isBlack() {
		if x == x.parent.getLeft() {
			x = rb.fixAfterDeleteLeft(x)
		} else {
			x = rb.fixAfterDeleteRight(x)
		}
	}
	x.setColor(Black)
}

// fixAfterDeleteLeft 处理x为左子节点时的平衡处理
func (rb *RBTree[K, V]) fixAfterDeleteLeft(x *rbNode[K, V]) *rbNode[K, V] {
	sib := x.getParent().getRight()
	if sib.isRed() {
		sib.setColor(Black)
		sib.getParent().setColor(Red)
		rb.rotateLeft(x.getParent())
		sib = x.getParent().getRight()
	}
	if sib.getLeft().isBlack() && sib.getRight().isBlack() {
		sib.setColor(Red)
		x = x.getParent()
	} else {
		if sib.getRight().getColor() == Black {
			sib.getLeft().setColor(Black)
			sib.setColor(Red)
			rb.rotateRight(sib)
			sib = x.getParent().getRight()
		}
		sib.setColor(x.getParent().getColor())
		x.getParent().setColor(Black)
		sib.getRight().setColor(Black)
		rb.rotateLeft(x.getParent())
		x = rb.root
	}
	return x
}

// fixAfterDeleteRight 处理x为右子节点时的平衡处理
func (rb *RBTree[K, V]) fixAfterDeleteRight(x *rbNode[K, V]) *rbNode[K, V] {
	sib := x.getParent().getLeft()
	if sib.isRed() {
		sib.setColor(Black)
		x.getParent().setColor(Red)
		rb.rotateRight(x.getParent())
		sib = x.getBrother()
	}
	if sib.getRight().isBlack() && sib.getLeft().isBlack() {
		sib.setColor(Red)
		x = x.getParent()
	} else {
		if sib.getLeft().isBlack() {
			sib.getRight().setColor(Black)
			sib.setColor(Red)
			rb.rotateLeft(sib)
			sib = x.getParent().getLeft()
		}
		sib.setColor(x.getParent().getColor())
		x.getParent().setColor(Black)
		sib.getLeft().setColor(Black)
		rb.rotateRight(x.getParent())
		x = rb.root
	}
	return x
}

// rotateLeft 左旋转
//
//							  b                    a
//							/	\				  /	  \
//						  c       a  ->    		 b     y
//								 / \            /  \
//		                     	x    y     		c	x

func (rb *RBTree[K, V]) rotateLeft(node *rbNode[K, V]) {
	if node == nil || node.getRight() == nil {
		return
	}
	r := node.right
	node.right = r.left
	if r.left != nil {
		r.left.parent = node
	}
	r.parent = node.parent
	if node.parent == nil {
		rb.root = r
	} else if node.parent.left == node {
		node.parent.left = r
	} else {
		node.parent.right = r
	}
	r.left = node
	node.parent = r

}

// rotateRight 右旋转
//
//						  b                    c
//						/	\				  /	  \
//					  c       a  ->    		 x     b
//					 /	\	                       / \
//	                 x  y  	     	 	  	       y  a
func (rb *RBTree[K, V]) rotateRight(node *rbNode[K, V]) {
	if node == nil || node.getLeft() == nil {
		return
	}
	l := node.left
	node.left = l.right
	if l.right != nil {
		l.right.parent = node
	}
	l.parent = node.parent
	if node.parent == nil {
		rb.root = l
	} else if node.parent.right == node {
		node.parent.right = l
	} else {
		node.parent.left = l
	}
	l.right = node
	node.parent = l

}

func (node *rbNode[K, V]) getColor() color {
	if node == nil {
		return Black
	}
	return node.color
}

func (node *rbNode[K, V]) isBlack() bool {
	return node == nil || node.color == Black
}

func (node *rbNode[K, V]) isRed() bool {
	return node != nil && node.color == Red
}

func (node *rbNode[K, V]) setColor(color color) {
	if node == nil {
		return
	}
	node.color = color
}

func (node *rbNode[K, V]) getParent() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	return node.parent
}

func (node *rbNode[K, V]) getLeft() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	return node.left
}

func (node *rbNode[K, V]) getRight() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	return node.right
}

func (node *rbNode[K, V]) getUncle() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	return node.getParent().getBrother()
}
func (node *rbNode[K, V]) getGrandParent() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	return node.getParent().getParent()
}
func (node *rbNode[K, V]) getBrother() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	if node == node.getParent().getLeft() {
		return node.getParent().getRight()
	}
	return node.getParent().getLeft()
}
