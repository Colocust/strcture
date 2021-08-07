package strcture

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_AddHead(t *testing.T) {
	list := NewList()
	for i := 0; i < 10; i++ {
		list.AddHead(i)
		assert.Equal(t, i, list.head.value.(int))
	}

	assert.Equal(t, uint(10), list.Len())

	n := list.head
	for i := 9; i >= 0; i-- {
		assert.Equal(t, i, n.value.(int))
		n = n.next
	}

	n = list.tail
	for i := 0; i < 10; i++ {
		assert.Equal(t, i, n.value.(int))
		n = n.prev
	}
}

func TestList_AddTail(t *testing.T) {
	list := NewList()
	for i := 0; i < 10; i++ {
		list.AddTail(i)
		assert.Equal(t, i, list.tail.value.(int))
	}

	assert.Equal(t, uint(10), list.Len())

	n := list.head
	for i := 0; i < 10; i++ {
		assert.Equal(t, i, n.value.(int))
		n = n.next
	}

	n = list.tail
	for i := 9; i >= 0; i-- {
		assert.Equal(t, i, n.value.(int))
		n = n.prev
	}
}

func TestList_RotateHeadToTail(t *testing.T) {
	list := NewList()
	for i := 1; i < 10; i++ {
		list.AddHead(i)
	}
	list.AddHead(0)
	assert.Equal(t, 0, list.Head().value.(int))
	list.RotateHeadToTail()

	assert.Equal(t, 0, list.Tail().value.(int))
	assert.Equal(t, 9, list.Head().value.(int))

	node := list.head
	for i := 9; i >= 0; i-- {
		assert.Equal(t, i, node.value.(int))
		node = node.next
	}
}

func TestList_RotateTailToHead(t *testing.T) {
	list := NewList()
	for i := 1; i < 10; i++ {
		list.AddTail(i)
	}
	list.AddTail(0)
	assert.Equal(t, 0, list.Tail().value.(int))
	assert.Equal(t, 1, list.Head().value.(int))
	list.RotateTailToHead()

	assert.Equal(t, 0, list.Head().value.(int))
	assert.Equal(t, 9, list.Tail().value.(int))

	node := list.head
	for i := 0; i < 10; i++ {
		assert.Equal(t, i, node.value.(int))
		node = node.next
	}
}

func TestList_Index(t *testing.T) {
	list := NewList()
	for i := 0; i < 10; i++ {
		list.AddTail(i)
	}

	assert.Equal(t, 0, list.Index(0).value.(int))
	assert.Equal(t, 9, list.Index(9).value.(int))
	assert.Equal(t, 9, list.Index(-1).value.(int))
	assert.Equal(t, 0, list.Index(-10).value.(int))
	assert.Equal(t, 5, list.Index(-5).value.(int))
	assert.Equal(t, 3, list.Index(3).value.(int))
	assert.Nil(t, list.Index(-11))
	assert.Nil(t, list.Index(10))
}

func TestList_Insert(t *testing.T) {
	list := NewList()
	for i := 1; i < 5; i++ {
		list.AddTail(i)
	}

	n1 := list.Index(0)
	list.Insert(n1, 0, false)

	assert.Equal(t, 0, list.Head().value.(int))
	assert.Equal(t, uint(5), list.Len())

	n4 := list.Index(4)
	list.Insert(n4, 7, true)

	assert.Equal(t, 7, list.Tail().value.(int))
	assert.Equal(t, uint(6), list.Len())

	list.Insert(n4, 5, true)
	assert.Equal(t, 5, list.Index(5).value.(int))
	assert.Equal(t, uint(7), list.Len())

	tail := list.Tail()
	list.Insert(tail, 6, false)
	assert.Equal(t, 6, list.Index(6).value.(int))
	assert.Equal(t, uint(8), list.Len())

	cur := list.Head()
	for i := 0; i < 8; i++ {
		assert.Equal(t, i, cur.value.(int))
		cur = cur.next
	}

	cur = list.Tail()
	for i := 7; i >= 0; i-- {
		assert.Equal(t, i, cur.value.(int))
		cur = cur.prev
	}
}

func TestList_Del(t *testing.T) {
	list := NewList()
	for i := 0; i < 10; i++ {
		list.AddTail(i)
	}

	list.Del(list.Head())
	assert.Equal(t, 1, list.Head().value.(int))
	assert.Equal(t, uint(9), list.Len())

	list.Del(list.Tail())
	assert.Equal(t, 8, list.Tail().value.(int))
	assert.Equal(t, uint(8), list.Len())

	cur := list.Head()
	for i := 1; i < 9; i++ {
		assert.Equal(t, i, cur.value.(int))
		cur = cur.next
	}

	cur = list.Tail()
	for i := 8; i >= 1; i-- {
		assert.Equal(t, i, cur.value.(int))
		cur = cur.prev
	}

	list.Del(list.Index(4))
	assert.Equal(t, uint(7), list.Len())

	cur = list.Head()
	for i := 1; i < 8; i++ {
		if i == 5 {
			assert.Equal(t, 6, cur.value.(int))
		}

		cur = cur.next
	}

	cur = list.Tail()
	for i := 7; i >= 1; i-- {
		if i == 4 {
			assert.Equal(t, 4, cur.value.(int))
		}
		cur = cur.prev
	}

	assert.Nil(t, list.Search(5))
}

func TestList_Search(t *testing.T) {
	list := NewList()
	for i := 0; i < 10; i++ {
		list.AddTail(i)
	}

	assert.Equal(t, 0, list.Search(0).value.(int))
	assert.Equal(t, 9, list.Search(9).value.(int))
	assert.Nil(t, list.Search(-1))
}
