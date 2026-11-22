package purescript

import "gopurs/output/gopurs_runtime"

type CompareFn func(a, b interface{}) int

const maxDegree = 6
const minItems = (maxDegree - 1) / 2

type Item struct {
	Key   interface{}
	Value interface{}
}

type Node struct {
	items    []Item
	children []*Node
}

func (n *Node) clone() *Node {
	newNode := &Node{
		items:    make([]Item, len(n.items), maxDegree),
		children: make([]*Node, len(n.children), maxDegree+1),
	}
	copy(newNode.items, n.items)
	copy(newNode.children, n.children)
	return newNode
}

func (n *Node) find(key interface{}, cmp CompareFn) (int, bool) {
	for i, item := range n.items {
		c := cmp(key, item.Key)
		if c == 0 {
			return i, true
		} else if c < 0 {
			return i, false
		}
	}
	return len(n.items), false
}

type BTree struct {
	root    *Node
	size    int
	compare CompareFn
}
type anyValuer interface {
	AnyVal() any
}

func asTree(m interface{}) *BTree {
	if val, ok := m.(anyValuer); ok {
		return val.AnyVal().(*BTree)
	}
	return m.(*BTree)
}

// A comparator belongs to this operation; never mutate a shared persistent tree.
func (t *BTree) withCompare(cmp CompareFn) *BTree {
	copied := *t
	copied.compare = cmp
	return &copied
}

func newBTree(cmp CompareFn) *BTree {
	return &BTree{compare: cmp}
}

func (t *BTree) Size() int { return t.size }

func (t *BTree) Lookup(key interface{}) (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	return t.root.lookup(key, t.compare)
}

func (n *Node) lookup(key interface{}, cmp CompareFn) (interface{}, bool) {
	i, found := n.find(key, cmp)
	if found {
		return n.items[i].Value, true
	}
	if len(n.children) == 0 {
		return nil, false
	}
	return n.children[i].lookup(key, cmp)
}

func (t *BTree) Insert(key, value interface{}) *BTree {
	if t.root == nil {
		return &BTree{
			root:    &Node{items: []Item{{key, value}}},
			size:    1,
			compare: t.compare,
		}
	}

	newRoot, splitItem, rightNode, replaced := t.root.insert(key, value, t.compare)
	newSize := t.size
	if !replaced {
		newSize++
	}

	if rightNode != nil {
		// Root split
		return &BTree{
			root: &Node{
				items:    []Item{splitItem},
				children: []*Node{newRoot, rightNode},
			},
			size:    newSize,
			compare: t.compare,
		}
	}

	return &BTree{
		root:    newRoot,
		size:    newSize,
		compare: t.compare,
	}
}

func (n *Node) insert(key, value interface{}, cmp CompareFn) (*Node, Item, *Node, bool) {
	i, found := n.find(key, cmp)
	// The item list always changes, so it is copied with room for the single
	// item the insert path may append before a split. The child list is copied
	// only when this node actually descends; a value replacement can share it
	// because published nodes are never mutated.
	newNode := &Node{
		items: make([]Item, len(n.items), len(n.items)+1),
	}
	copy(newNode.items, n.items)

	if found {
		newNode.children = n.children
		newNode.items[i].Value = value
		return newNode, Item{}, nil, true
	}

	var replaced bool
	var splitItem Item
	var rightChild *Node

	if len(n.children) == 0 {
		newNode.children = nil
		newNode.items = append(newNode.items, Item{})
		copy(newNode.items[i+1:], newNode.items[i:])
		newNode.items[i] = Item{key, value}
	} else {
		newNode.children = make([]*Node, len(n.children), len(n.children)+1)
		copy(newNode.children, n.children)
		var childNew *Node
		childNew, splitItem, rightChild, replaced = n.children[i].insert(key, value, cmp)
		newNode.children[i] = childNew
		if rightChild != nil {
			newNode.items = append(newNode.items, Item{})
			copy(newNode.items[i+1:], newNode.items[i:])
			newNode.items[i] = splitItem

			newNode.children = append(newNode.children, nil)
			copy(newNode.children[i+2:], newNode.children[i+1:])
			newNode.children[i+1] = rightChild
		}
	}

	if len(newNode.items) >= maxDegree {
		mid := maxDegree / 2
		upItem := newNode.items[mid]

		rightNode := &Node{
			items: make([]Item, len(newNode.items[mid+1:]), maxDegree),
		}
		copy(rightNode.items, newNode.items[mid+1:])
		if len(newNode.children) > 0 {
			rightNode.children = make([]*Node, len(newNode.children[mid+1:]), maxDegree+1)
			copy(rightNode.children, newNode.children[mid+1:])
			newNode.children = newNode.children[:mid+1]
		}
		newNode.items = newNode.items[:mid]

		return newNode, upItem, rightNode, replaced
	}

	return newNode, Item{}, nil, replaced
}

func (t *BTree) Delete(key interface{}) *BTree {
	if t.root == nil {
		return t
	}

	newRoot, deleted := t.root.delete(key, t.compare)
	if !deleted {
		return t // unchanged
	}

	// If root has no items and has one child, the child becomes the new root
	if len(newRoot.items) == 0 && len(newRoot.children) == 1 {
		newRoot = newRoot.children[0]
	} else if len(newRoot.items) == 0 && len(newRoot.children) == 0 {
		newRoot = nil
	}

	return &BTree{
		root:    newRoot,
		size:    t.size - 1,
		compare: t.compare,
	}
}

func (n *Node) delete(key interface{}, cmp CompareFn) (*Node, bool) {
	i, found := n.find(key, cmp)
	newNode := n.clone()

	if len(n.children) == 0 {
		if found {
			// Leaf node, just remove the item
			newNode.items = append(newNode.items[:i], newNode.items[i+1:]...)
			return newNode, true
		}
		return nil, false
	}

	if found {
		// Internal node, replace with predecessor
		// Predecessor is the largest item in children[i]
		predChild, predItem := n.children[i].deleteMax()
		newNode.items[i] = predItem
		newNode.children[i] = predChild

		newNode.rebalanceChild(i)
		return newNode, true
	}

	// Not found, go down to children[i]
	childNew, deleted := n.children[i].delete(key, cmp)
	if !deleted {
		return nil, false
	}
	newNode.children[i] = childNew
	newNode.rebalanceChild(i)
	return newNode, true
}

func (n *Node) deleteMax() (*Node, Item) {
	newNode := n.clone()
	if len(n.children) == 0 {
		item := newNode.items[len(newNode.items)-1]
		newNode.items = newNode.items[:len(newNode.items)-1]
		return newNode, item
	}

	lastIdx := len(n.children) - 1
	childNew, item := n.children[lastIdx].deleteMax()
	newNode.children[lastIdx] = childNew
	newNode.rebalanceChild(lastIdx)
	return newNode, item
}

// rebalanceChild ensures children[i] has at least minItems items.
// If it doesn't, it borrows from a sibling or merges.
func (n *Node) rebalanceChild(i int) {
	child := n.children[i]
	if len(child.items) >= minItems {
		return
	}

	// Try borrowing from left sibling
	if i > 0 && len(n.children[i-1].items) > minItems {
		leftSib := n.children[i-1].clone()
		newChild := child.clone()

		// Move n.items[i-1] down to child
		newChild.items = append([]Item{n.items[i-1]}, newChild.items...)
		if len(leftSib.children) > 0 {
			newChild.children = append([]*Node{leftSib.children[len(leftSib.children)-1]}, newChild.children...)
			leftSib.children = leftSib.children[:len(leftSib.children)-1]
		}

		// Move leftSib's last item up to n
		n.items[i-1] = leftSib.items[len(leftSib.items)-1]
		leftSib.items = leftSib.items[:len(leftSib.items)-1]

		n.children[i-1] = leftSib
		n.children[i] = newChild
		return
	}

	// Try borrowing from right sibling
	if i < len(n.children)-1 && len(n.children[i+1].items) > minItems {
		rightSib := n.children[i+1].clone()
		newChild := child.clone()

		// Move n.items[i] down to child
		newChild.items = append(newChild.items, n.items[i])
		if len(rightSib.children) > 0 {
			newChild.children = append(newChild.children, rightSib.children[0])
			rightSib.children = rightSib.children[1:]
		}

		// Move rightSib's first item up to n
		n.items[i] = rightSib.items[0]
		rightSib.items = rightSib.items[1:]

		n.children[i+1] = rightSib
		n.children[i] = newChild
		return
	}

	// Must merge. We merge child and a sibling.
	if i > 0 {
		// Merge children[i-1] and children[i]
		leftSib := n.children[i-1].clone()
		// Bring n.items[i-1] down
		leftSib.items = append(leftSib.items, n.items[i-1])
		leftSib.items = append(leftSib.items, child.items...)
		if len(child.children) > 0 {
			leftSib.children = append(leftSib.children, child.children...)
		}

		// Remove n.items[i-1] and n.children[i]
		n.items = append(n.items[:i-1], n.items[i:]...)
		n.children = append(n.children[:i], n.children[i+1:]...)
		n.children[i-1] = leftSib
	} else {
		// Merge children[i] and children[i+1]
		newChild := child.clone()
		rightSib := n.children[i+1]

		newChild.items = append(newChild.items, n.items[i])
		newChild.items = append(newChild.items, rightSib.items...)
		if len(rightSib.children) > 0 {
			newChild.children = append(newChild.children, rightSib.children...)
		}

		n.items = append(n.items[:i], n.items[i+1:]...)
		n.children = append(n.children[:i+1], n.children[i+2:]...)
		n.children[i] = newChild
	}
}

type FoldFn func(acc, key, value interface{}) interface{}

func (t *BTree) Foldl(f FoldFn, acc interface{}) interface{} {
	if t.root == nil {
		return acc
	}
	return t.root.foldl(f, acc)
}

func (n *Node) foldl(f FoldFn, acc interface{}) interface{} {
	if len(n.children) == 0 {
		for _, item := range n.items {
			acc = f(acc, item.Key, item.Value)
		}
		return acc
	}

	for i, item := range n.items {
		acc = n.children[i].foldl(f, acc)
		acc = f(acc, item.Key, item.Value)
	}
	acc = n.children[len(n.items)].foldl(f, acc)
	return acc
}

func (t *BTree) Foldr(f FoldFn, acc interface{}) interface{} {
	if t.root == nil {
		return acc
	}
	return t.root.foldr(f, acc)
}

func (n *Node) foldr(f FoldFn, acc interface{}) interface{} {
	if len(n.children) == 0 {
		for i := len(n.items) - 1; i >= 0; i-- {
			item := n.items[i]
			acc = f(acc, item.Key, item.Value)
		}
		return acc
	}

	acc = n.children[len(n.items)].foldr(f, acc)
	for i := len(n.items) - 1; i >= 0; i-- {
		item := n.items[i]
		acc = f(acc, item.Key, item.Value)
		acc = n.children[i].foldr(f, acc)
	}
	return acc
}

func (t *BTree) Keys() []interface{} {
	keys := make([]interface{}, 0, t.size)
	res := t.Foldl(func(acc, key, value interface{}) interface{} {
		return append(acc.([]interface{}), key)
	}, keys)
	return res.([]interface{})
}

func (t *BTree) Values() []interface{} {
	values := make([]interface{}, 0, t.size)
	res := t.Foldl(func(acc, key, value interface{}) interface{} {
		return append(acc.([]interface{}), value)
	}, values)
	return res.([]interface{})
}

type CombineFn func(v1, v2 interface{}) interface{}

func (t *BTree) UnionWith(other *BTree, f CombineFn) *BTree {
	if t.size == 0 {
		return other
	}
	if other.size == 0 {
		return t
	}

	res := t
	finalRes := other.Foldl(func(acc, key, value interface{}) interface{} {
		tree := asTree(acc)
		existing, ok := tree.Lookup(key)
		if ok {
			newVal := f(existing, value)
			return tree.Insert(key, newVal)
		}
		return tree.Insert(key, value)
	}, res)

	return asTree(finalRes)
}

// PureScript maps of the same key type use a coherent Ord instance. At that
// boundary, reuse the larger tree and insert only the smaller side. Keep the
// raw UnionWith method above for trees that may have different orderings.
func unionWithSameOrdering(left, right *BTree, f CombineFn) *BTree {
	// Overlapping keys need two updates below. Keep the original path for
	// similarly sized maps; a fourfold imbalance also wins with full overlap.
	if left.size == 0 || left.size > right.size/4 {
		return left.UnionWith(right, f)
	}

	return asTree(left.Foldl(func(acc, key, value interface{}) interface{} {
		tree := asTree(acc)
		if existing, ok := tree.Lookup(key); ok {
			// Insert preserves an existing key. Replace it explicitly so an
			// Ord-equivalent key still has the left map's representative.
			return tree.Delete(key).Insert(key, f(value, existing))
		}
		return tree.Insert(key, value)
	}, right))
}

func (t *BTree) IntersectionWith(other *BTree, f CombineFn) *BTree {
	res := newBTree(t.compare)
	if t.size == 0 || other.size == 0 {
		return res
	}

	smaller, larger := t, other
	if t.size > other.size {
		smaller, larger = other, t
	}

	finalRes := smaller.Foldl(func(acc, key, value interface{}) interface{} {
		tree := asTree(acc)
		if otherVal, ok := larger.Lookup(key); ok {
			var newVal interface{}
			if smaller == t {
				newVal = f(value, otherVal)
			} else {
				newVal = f(otherVal, value)
			}
			return tree.Insert(key, newVal)
		}
		return tree
	}, res)

	return asTree(finalRes)
}

func (t *BTree) Difference(other *BTree) *BTree {
	res := t
	if t.size == 0 || other.size == 0 {
		return res
	}

	finalRes := other.Foldl(func(acc, key, _ interface{}) interface{} {
		tree := asTree(acc)
		return tree.Delete(key)
	}, res)

	return asTree(finalRes)
}

// FFI Wrappers

var Data_Map_Internal_Empty = func() interface{} {
	// A dummy cmp function because the BTree struct requires it,
	// but it will be replaced on the first insert since insert creates a new tree.
	// Actually, an empty tree does not know its cmp until the first insert!
	// In our FFI, we pass cmp on EVERY operation.
	// So newBTree(nil) is fine.
	return newBTree(nil)
}()

func Data_Map_Internal_IsEmpty(m interface{}) bool {
	return asTree(m).Size() == 0
}

func Data_Map_Internal_Singleton(k interface{}) func(interface{}) interface{} {
	return func(v interface{}) interface{} {
		return &BTree{
			root:    &Node{items: []Item{{k, v}}},
			size:    1,
			compare: nil, // compare will be injected on next insert
		}
	}
}

func Data_Map_Internal_InsertImpl(compareInt func(interface{}, interface{}) int, k interface{}, v interface{}, m interface{}) interface{} {
	tree := asTree(m)
	cmp := compareInt
	if tree.Size() == 0 {
		return newBTree(cmp).Insert(k, v)
	}
	// We must ensure the tree uses the current cmp
	tree = tree.withCompare(cmp)
	return tree.Insert(k, v)
}

func Data_Map_Internal_InsertWithImpl(compareInt func(interface{}, interface{}) int, f func(interface{}) func(interface{}) interface{}, k interface{}, v interface{}, m interface{}) interface{} {
	tree := asTree(m)
	cmp := compareInt
	if tree.Size() == 0 {
		return newBTree(cmp).Insert(k, v)
	}
	tree = tree.withCompare(cmp)

	existing, ok := tree.Lookup(k)
	if ok {
		newVal := f(existing)(v)
		return tree.Insert(k, newVal)
	}
	return tree.Insert(k, v)
}

func Data_Map_Internal_LookupImpl(just func(interface{}) interface{}, nothing interface{}, compareInt func(interface{}, interface{}) int, k interface{}, m interface{}) interface{} {
	tree := asTree(m)
	if tree.Size() == 0 {
		return nothing
	}
	cmp := compareInt
	tree = tree.withCompare(cmp)
	val, ok := tree.Lookup(k)
	if ok {
		return just(val)
	}
	return nothing
}

// LookupNative is the raw native-comparator lookup: the comparator is only
// visible to Go, and absence is reported with a boolean instead of a boxed
// `Maybe` (whose representation belongs to the caller's backend).
func Data_Map_Internal_LookupNative(compare CompareFn, k interface{}, m interface{}) (interface{}, bool) {
	tree := asTree(m)
	if tree.Size() == 0 {
		return nil, false
	}
	return tree.withCompare(compare).Lookup(k)
}

// InsertNative is the raw native-comparator insert. As with every other
// operation the comparator belongs to this call; the shared tree is untouched.
func Data_Map_Internal_InsertNative(compare CompareFn, k interface{}, v interface{}, m interface{}) interface{} {
	tree := asTree(m)
	if tree.Size() == 0 {
		return newBTree(compare).Insert(k, v)
	}
	return tree.withCompare(compare).Insert(k, v)
}

func Data_Map_Internal_LookupNativeImpl(compare CompareFn, just func(interface{}) interface{}, nothing interface{}, k interface{}, m interface{}) interface{} {
	if val, ok := Data_Map_Internal_LookupNative(compare, k, m); ok {
		return just(val)
	}
	return nothing
}

func Data_Map_Internal_InsertNativeImpl(compare CompareFn, k interface{}, v interface{}, m interface{}) interface{} {
	return Data_Map_Internal_InsertNative(compare, k, v, m)
}

func Data_Map_Internal_DeleteImpl(compareInt func(interface{}, interface{}) int, k interface{}, m interface{}) interface{} {
	tree := asTree(m)
	if tree.Size() == 0 {
		return m
	}
	cmp := compareInt
	tree = tree.withCompare(cmp)
	return tree.Delete(k)
}

func Data_Map_Internal_KeysImpl(m interface{}) []interface{} {
	return asTree(m).Keys()
}

func Data_Map_Internal_ValuesImpl(m interface{}) []interface{} {
	return asTree(m).Values()
}

func Data_Map_Internal_ToArrayImpl(tuple func(interface{}) func(interface{}) interface{}, m interface{}) []interface{} {
	tree := asTree(m)
	items := make([]interface{}, 0, tree.Size())
	tree.Foldl(func(_ interface{}, key, value interface{}) interface{} {
		items = append(items, tuple(key)(value))
		return nil
	}, nil)
	return items
}

func Data_Map_Internal_UnionWithImpl(compareInt func(interface{}, interface{}) int, f func(interface{}) func(interface{}) interface{}, m1 interface{}, m2 interface{}) interface{} {
	t1 := asTree(m1)
	t2 := asTree(m2)
	cmp := compareInt
	t1 = t1.withCompare(cmp)
	t2 = t2.withCompare(cmp)
	return unionWithSameOrdering(t1, t2, func(v1, v2 interface{}) interface{} {
		return f(v1)(v2)
	})
}

// UnionWithNative is UnionWithImpl with a native comparator: the combine
// callback still crosses the backend boundary, but comparisons never do.
func Data_Map_Internal_UnionWithNative(compare CompareFn, f CombineFn, m1 interface{}, m2 interface{}) interface{} {
	t1 := asTree(m1).withCompare(compare)
	t2 := asTree(m2).withCompare(compare)
	return unionWithSameOrdering(t1, t2, f)
}

func Data_Map_Internal_IntersectionWithImpl(compareInt func(interface{}, interface{}) int, f func(interface{}) func(interface{}) interface{}, m1 interface{}, m2 interface{}) interface{} {
	t1 := asTree(m1)
	t2 := asTree(m2)
	cmp := compareInt
	t1 = t1.withCompare(cmp)
	t2 = t2.withCompare(cmp)
	return t1.IntersectionWith(t2, func(v1, v2 interface{}) interface{} {
		return f(v1)(v2)
	})
}

func Data_Map_Internal_DifferenceImpl(compareInt func(interface{}, interface{}) int, m1 interface{}, m2 interface{}) interface{} {
	t1 := asTree(m1)
	t2 := asTree(m2)
	cmp := compareInt
	t1 = t1.withCompare(cmp)
	t2 = t2.withCompare(cmp)
	return t1.Difference(t2)
}

func Data_Map_Internal_SizeImpl(m interface{}) int {
	return asTree(m).Size()
}

func (n *Node) mapValues(f func(interface{}) interface{}) *Node {
	if n == nil {
		return nil
	}
	newNode := &Node{
		items:    make([]Item, len(n.items), maxDegree),
		children: make([]*Node, len(n.children), maxDegree+1),
	}
	for i, item := range n.items {
		newNode.items[i] = Item{Key: item.Key, Value: f(item.Value)}
	}
	for i, child := range n.children {
		newNode.children[i] = child.mapValues(f)
	}
	return newNode
}

func (t *BTree) MapValues(f func(interface{}) interface{}) *BTree {
	return &BTree{
		root:    t.root.mapValues(f),
		size:    t.size,
		compare: t.compare,
	}
}

func Data_Map_Internal_MapImpl(f func(interface{}) interface{}, m interface{}) interface{} {
	return asTree(m).MapValues(f)
}

func Data_Map_Internal_FoldlImpl(f func(interface{}) func(interface{}) func(interface{}) interface{}, z interface{}, m interface{}) interface{} {
	tree := asTree(m)
	return tree.Foldl(func(acc, key, value interface{}) interface{} {
		return f(acc)(key)(value)
	}, z)
}

func Data_Map_Internal_FoldrImpl(f func(interface{}) func(interface{}) func(interface{}) interface{}, z interface{}, m interface{}) interface{} {
	tree := asTree(m)
	return tree.Foldr(func(acc, key, value interface{}) interface{} {
		return f(acc)(key)(value)
	}, z)
}

func Data_Map_Internal_FilterKeysImpl(p func(interface{}) interface{}, m interface{}) interface{} {
	tree := asTree(m)
	res := newBTree(tree.compare)
	finalRes := tree.Foldl(func(acc, key, value interface{}) interface{} {
		accTree := asTree(acc)
		pv := p(key)
		var isTrue bool
		if val, ok := pv.(gopurs_runtime.Value); ok {
			isTrue = val.IntVal != 0
		} else {
			panic("Expected Value from p(key)")
		}
		if isTrue {
			return accTree.Insert(key, value)
		}
		return accTree
	}, res)
	return asTree(finalRes)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Map_Internal_DeleteImpl = // TAST: (ForAll [k, v] (Func [(Func [(TypeVar k), (TypeVar k)] Int), (TypeVar k), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) int {
		inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		return gopurs_runtime.Unbox[int](inner_res0)
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Map_Internal_DeleteImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_DifferenceImpl = // TAST: (ForAll [k, v, w] (Func [(Func [(TypeVar k), (TypeVar k)] Int), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)]), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar w)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) int {
		inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		return gopurs_runtime.Unbox[int](inner_res0)
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Map_Internal_DifferenceImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_Empty = // TAST: (ForAll [k, v] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)]))
gopurs_runtime.Box(Data_Map_Internal_Empty)
var _Gopurs_Data_Map_Internal_FilterKeysImpl = // TAST: (ForAll [k, v] (Func [(Func [(TypeVar k)] Boolean), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_res := Data_Map_Internal_FilterKeysImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_FoldlImpl = // TAST: (ForAll [k, v, z] (Func [(Func [(TypeVar z), (TypeVar k), (TypeVar v)] (TypeVar z)), (TypeVar z), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (TypeVar z)))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) any {
				return gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			}
		}
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Map_Internal_FoldlImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_FoldrImpl = // TAST: (ForAll [k, v, z] (Func [(Func [(TypeVar z), (TypeVar k), (TypeVar v)] (TypeVar z)), (TypeVar z), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (TypeVar z)))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) any {
				return gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			}
		}
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Map_Internal_FoldrImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_InsertImpl = // TAST: (ForAll [k, v] (Func [(Func [(TypeVar k), (TypeVar k)] Int), (TypeVar k), (TypeVar v), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) int {
		inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		return gopurs_runtime.Unbox[int](inner_res0)
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Data_Map_Internal_InsertImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_InsertNativeImpl = // TAST: (ForAll [k, v] (Func [(ADT ["Data","Map","Internal","Comparator"] [(TypeVar k)]), (TypeVar k), (TypeVar v), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[CompareFn](arg0)
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Data_Map_Internal_InsertNativeImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_InsertWithImpl = // TAST: (ForAll [k, v] (Func [(Func [(TypeVar k), (TypeVar k)] Int), (Func [(TypeVar v), (TypeVar v)] (TypeVar v)), (TypeVar k), (TypeVar v), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) int {
		inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		return gopurs_runtime.Unbox[int](inner_res0)
	}
	go_arg1 := func(p0_0 any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Data_Map_Internal_InsertWithImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_IntersectionWithImpl = // TAST: (ForAll [k, a, b, c] (Func [(Func [(TypeVar k), (TypeVar k)] Int), (Func [(TypeVar a), (TypeVar b)] (TypeVar c)), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar a)]), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar b)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar c)])))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) int {
		inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		return gopurs_runtime.Unbox[int](inner_res0)
	}
	go_arg1 := func(p0_0 any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Data_Map_Internal_IntersectionWithImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_IsEmpty = // TAST: (ForAll [k, v] (Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] Boolean))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Map_Internal_IsEmpty(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_KeysImpl = // TAST: (ForAll [k, v] (Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (Array (TypeVar k))))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Map_Internal_KeysImpl(go_arg0)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
var _Gopurs_Data_Map_Internal_LookupImpl = // TAST: (ForAll [k, v] (Func [(Func [(TypeVar v)] (ADT ["Data","Maybe","Maybe"] [(TypeVar v)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar v)]), (Func [(TypeVar k), (TypeVar k)] Int), (TypeVar k), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Maybe","Maybe"] [(TypeVar v)])))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := func(p0_0 any, p0_1 any) int {
		inner_res0 := gopurs_runtime.Apply2(arg2, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		return gopurs_runtime.Unbox[int](inner_res0)
	}
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Data_Map_Internal_LookupImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_LookupNativeImpl = // TAST: (ForAll [k, v] (Func [(ADT ["Data","Map","Internal","Comparator"] [(TypeVar k)]), (Func [(TypeVar v)] (ADT ["Data","Maybe","Maybe"] [(TypeVar v)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar v)]), (TypeVar k), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Maybe","Maybe"] [(TypeVar v)])))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[CompareFn](arg0)
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Data_Map_Internal_LookupNativeImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_MapImpl = // TAST: (ForAll [k, a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar a)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar b)])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_res := Data_Map_Internal_MapImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_Singleton = // TAST: (ForAll [k, v] (Func [(TypeVar k), (TypeVar v)] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Map_Internal_Singleton(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
		inner_res := go_res(arg)
		return gopurs_runtime.Box(inner_res)
	})
})
var _Gopurs_Data_Map_Internal_SizeImpl = // TAST: (ForAll [k, v] (Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] Int))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Map_Internal_SizeImpl(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_Map_Internal_ToArrayImpl = // TAST: (ForAll [k, v, a] (Func [(Func [(TypeVar k), (TypeVar v)] (TypeVar a)), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (Array (TypeVar a))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_arg1 := arg1
	go_res := Data_Map_Internal_ToArrayImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
var _Gopurs_Data_Map_Internal_UnionWithImpl = // TAST: (ForAll [k, v] (Func [(Func [(TypeVar k), (TypeVar k)] Int), (Func [(TypeVar v), (TypeVar v)] (TypeVar v)), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)]), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) int {
		inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		return gopurs_runtime.Unbox[int](inner_res0)
	}
	go_arg1 := func(p0_0 any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Data_Map_Internal_UnionWithImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Map_Internal_ValuesImpl = // TAST: (ForAll [k, v] (Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (Array (TypeVar v))))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Map_Internal_ValuesImpl(go_arg0)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
