package strcture

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap_Size(t *testing.T) {
	h := NewHeap()
	assert.Equal(t, h.Size(), 0)
}

func TestHeap_IsEmpty(t *testing.T) {
	h := NewHeap()
	assert.Equal(t, h.IsEmpty(), true)

	n := NewHeapNode(1, 2)
	h.Add(n)

	assert.Equal(t, h.IsEmpty(), false)
}

func TestHeap_Parent(t *testing.T) {
	h := NewHeap()
	assert.Equal(t, h.parent(0), -1)
	assert.Equal(t, h.parent(1), 0)
	assert.Equal(t, h.parent(2), 0)
	assert.Equal(t, h.parent(3), 1)
	assert.Equal(t, h.parent(4), 1)
}

func TestHeap_Swap(t *testing.T) {
	i, j := 0, 1
	h := NewHeap()
	assert.Error(t, h.swap(i, j))

	n := NewHeapNode(1, 2)
	h.Add(n)
	n = NewHeapNode(1, 2)
	h.Add(n)
	assert.Nil(t, h.swap(i, j))
}

func TestHeap(t *testing.T) {
	h := NewHeap()

	for i := 100000; i > 0; i-- {
		n := NewHeapNode(i, i)
		h.Add(n)
	}

	for i := 0; i < 100000; i++ {
		assert.Equal(t, i+1, h.Top().key)
		h.Remove()
	}
}
