package strcture

import "errors"

type (
	Heap struct {
		data []*HeapNode
	}

	HeapNode struct {
		key   int
		value interface{}
	}
)

func NewHeap() *Heap {
	return &Heap{}
}

func NewHeapNode(key int, value interface{}) *HeapNode {
	return &HeapNode{
		key:   key,
		value: value,
	}
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) IsEmpty() bool {
	return h.Size() == 0
}

func (h *Heap) leftChild(index int) int {
	return index*2 + 1
}

func (h *Heap) rightChild(index int) int {
	return index*2 + 2
}

func (h *Heap) parent(index int) int {
	if index == 0 {
		return -1
	}
	return (index - 1) / 2
}

func (h *Heap) swap(i, j int) (err error) {
	if i < 0 || j < 0 || i >= h.Size() || j >= h.Size() {
		err = errors.New("error index")
		return
	}

	h.data[i], h.data[j] = h.data[j], h.data[i]
	return
}

func (h *Heap) siftUp(index int) (err error) {
	for index >= 0 {
		parent := h.parent(index)
		if parent == -1 {
			return
		}
		if h.data[index].key > h.data[parent].key {
			return
		}
		if err = h.swap(index, parent); err != nil {
			return err
		}
		index = parent
	}
	return
}

func (h *Heap) siftDown(index int) (err error) {
	for h.leftChild(index) < h.Size() {
		i := h.leftChild(index)
		if h.rightChild(index) < h.Size() && h.data[h.leftChild(index)].key > h.data[h.rightChild(index)].key {
			i++
		}
		if h.data[i].key > h.data[index].key {
			return
		}
		if err = h.swap(index, i); err != nil {
			return err
		}
		index = i
	}
	return
}

func (h *Heap) Top() *HeapNode {
	if h.IsEmpty() {
		return nil
	}
	return h.data[0]
}

func (h *Heap) Remove() (n *HeapNode, err error) {
	n = h.Top()
	if n == nil {
		return
	}
	if err = h.swap(0, h.Size()-1); err != nil {
		return
	}

	h.data = h.data[:h.Size()-1]
	if err = h.siftDown(0); err != nil {
		return
	}
	return
}

func (h *Heap) Add(n *HeapNode) (err error) {
	h.data = append(h.data, n)
	if err = h.siftUp(h.Size() - 1); err != nil {
		return err
	}
	return
}

func (n *HeapNode) GetKey() int {
	return n.key
}

func (n *HeapNode) GetValue() interface{} {
	return n.value
}
