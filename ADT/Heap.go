package ADT

import "fmt"

// Heap is a min-heap of any type.
type Heap[T Integer] struct {
	values []T
	less   func(a, b T) bool
}

// NewHeap creates a new heap with the given less function.
func NewHeap[T Integer](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		values: []T{},
		less:   less,
	}
}

// Len returns the number of elements in the heap.
func (h *Heap[T]) Len() int {
	return len(h.values)
}

// Push adds a value to the heap.
func (h *Heap[T]) Push(x T) {
	h.values = append(h.values, x)
	h.heapfyUp(T(h.Len() - 1))
}

// Pop removes and returns the minimum value from the heap.
func (h *Heap[T]) Pop() (T, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("pop from empty heap: %v", r)
		}
	}()
	n := h.Len()
	if n == 0 {
		panic("pop from empty heap")
	}
	x := h.values[0]
	h.values[0] = h.values[n-1]
	h.values = h.values[:n-1]
	h.heapfyDown(0, n-1)
	return x, err
}

func (h *Heap[T]) heapfyUp(j T) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(h.values[j], h.values[i]) {
			break
		}
		h.values[i], h.values[j] = h.values[j], h.values[i]
		j = i
	}
}

func (h *Heap[T]) heapfyDown(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(h.values[j2], h.values[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(h.values[j], h.values[i]) {
			break
		}
		h.values[i], h.values[j] = h.values[j], h.values[i]
		i = j
	}
}
