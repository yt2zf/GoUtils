package list

import "github.com/yt2zf/GoUtils/internal/errs"

var (
	_ List[any] = &LinkedList[any]{} // LinkedList 实现List接口
)

type node[T any] struct {
	prev *node[T]
	next *node[T]
	val  T
}

// 双向循环链表
type LinkedList[T any] struct {
	head   *node[T] // 虚拟头节点
	tail   *node[T] // 虚拟尾节点
	length int
}

// 创建双向循环链表，当成constrcutor
func NewLinkedList[T any]() *LinkedList[T] {
	head := &node[T]{}
	tail := &node[T]{next: head, prev: head}
	head.next, head.prev = tail, tail
	return &LinkedList[T]{
		head: head,
		tail: tail,
	}
}

// 将切片转化为双向循环链表，切片的元素值直接使用 没有复制
func NewLinkedListOf[T any](ts []T) *LinkedList[T] {
	list := NewLinkedList[T]()
	if err := list.Append(ts...); err != nil {
		panic(err)
	}
	return list
}

func (l *LinkedList[T]) findNode(index int) *node[T] {
	var cur *node[T]
	if index <= l.Len()/2 {
		cur = l.head
		for i := -1; i < index; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail
		for i := l.Len(); i > index; i-- {
			cur = cur.prev
		}
	}
	return cur
}

// 在下标为index的位置插入一个元素
// 当index等于 LinkedList长度时，效果等同于append
func (l *LinkedList[T]) Add(index int, t T) error {
	if index < 0 || index > l.Len() {
		return errs.NewErrIndexOutOfRange(l.Len(), index)
	}
	if index == l.length {
		return l.Append(t)
	}
	next := l.findNode(index)
	node := &node[T]{prev: next.prev, next: next, val: t}
	node.prev.next, node.next.prev = node, node
	l.length++
	return nil

}

// Append implements List.
func (l *LinkedList[T]) Append(ts ...T) error {
	for _, t := range ts {
		node := &node[T]{prev: l.tail.prev, next: l.tail, val: t}
		node.prev.next, node.next.prev = node, node
		l.length++
	}
	return nil
}

// Set implements List.
func (l *LinkedList[T]) Set(index int, t T) error {
	if index < 0 || index >= l.Len() {
		return errs.NewErrIndexOutOfRange(l.Len(), index)
	}
	node := l.findNode(index)
	node.val = t
	return nil
}

// Cap implements List.
func (l *LinkedList[T]) Cap() int {
	return l.Len()
}

// Delete 删除指定位置的元素
func (l *LinkedList[T]) Delete(index int) (T, error) {
	if index < 0 || index >= l.Len() {
		var zero T
		return zero, errs.NewErrIndexOutOfRange(l.Len(), index)
	}

	node := l.findNode(index)
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev, node.next = nil, nil
	l.length--
	return node.val, nil
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.Len() {
		var zero T
		return zero, errs.NewErrIndexOutOfRange(l.Len(), index)
	}
	n := l.findNode(index)
	return n.val, nil
}

// Len
func (l *LinkedList[T]) Len() int {
	return l.length
}

// Range
func (l *LinkedList[T]) Range(fn func(index int, t T) error) error {
	for cur, i := l.head.next, 0; i < l.length; i++ {
		err := fn(i, cur.val)
		if err != nil {
			return err
		}
		cur = cur.next
	}
	return nil
}

// AsSlice implements List.
func (l *LinkedList[T]) AsSlice() []T {
	result := make([]T, l.Len())
	for cur, i := l.head.next, 0; i < l.length; i++ {
		result[i] = cur.val
		cur = cur.next
	}
	return result
}
