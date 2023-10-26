package ADT

import (
	"testing"
)

func TestHeap(t *testing.T) {
	//largest first
	heapMax := NewHeap[int](func(a, b int) bool {
		return a > b
	})
	//t.Logf("%#v", heapMax)

	heapMax.Push(1)
	heapMax.Push(2)
	heapMax.Push(3)
	heapMax.Push(444)
	heapMax.Push(5555555)

	if l := heapMax.Len(); l != 5 {
		t.Fatalf(`want len %v got %v`, 5, l)
	}
	if r, err := heapMax.Pop(); r != 5555555 || err != nil {
		t.Fatalf(`want %v got %v, err %v`, 5555555, r, err)
	}
	if r, err := heapMax.Pop(); r != 444 || err != nil {
		t.Fatalf(`want %v got %v, err %v`, 2, r, err)
	}
	if r, err := heapMax.Pop(); r != 3 || err != nil {
		t.Fatalf(`want %v got %v, err %v`, 1, r, err)
	}

	if l := heapMax.Len(); l != 2 {
		t.Fatalf(`want len %v got %v`, 5, l)
	}
	_, _ = heapMax.Pop()
	_, _ = heapMax.Pop()
	if r, err := heapMax.Pop(); err != nil {
		t.Fatalf(`want error %v,  got value %v and error nil`, err, r)
	}
	// smallest first
	heapMin := NewHeap[int](func(a, b int) bool {
		return a < b
	})
	heapMin.Push(1)
	heapMin.Push(2)
	heapMin.Push(3)
	heapMin.Push(444)
	heapMin.Push(5555555)

	if l := heapMin.Len(); l != 5 {
		t.Fatalf(`want len %v got %v`, 5, l)
	}
	if r, err := heapMin.Pop(); r != 1 || err != nil {
		t.Fatalf(`want %v got %v, err %v`, 3, r, err)
	}
	if r, err := heapMin.Pop(); r != 2 || err != nil {
		t.Fatalf(`want %v got %v, err %v`, 3, r, err)
	}
	if r, err := heapMin.Pop(); r != 3 || err != nil {
		t.Fatalf(`want %v got %v, err %v`, 3, r, err)
	}
	if r, err := heapMin.Pop(); r != 444 || err != nil {
		t.Fatalf(`want %v got %v, err %v`, 3, r, err)
	}
	if r, err := heapMin.Pop(); r != 5555555 || err != nil {
		t.Fatalf(`want %v got %v, err %v`, 5555555, r, err)
	}
}
