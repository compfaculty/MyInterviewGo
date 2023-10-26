package ADT

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	ht := NewHashTable[string]()
	t.Logf("%#v", ht)
	ht.Insert("test")  //&LinkedList{&Node{key: "a", next: nil}}
	ht.Insert("test1") //&LinkedList{&Node{key: "a", next: nil}}
	ht.Insert("test2") //&LinkedList{&Node{key: "a", next: nil}}
	if !ht.Search("test") {
		t.Fatalf(`want %v == %v`, "test", false)

	}
	if !ht.Search("test1") {
		t.Fatalf(`want %v == %v`, "test1", false)
	}
	if !ht.Search("test2") {
		t.Fatalf(`want %v == %v`, "test2", false)
	}
	ht.Remove("test")
	if ht.Search("test") {
		t.Fatalf(`want test key missing but it exists`)
	}
	ht.Remove("test1")
	if ht.Search("test1") {
		t.Fatalf(`want test key missing but it exists`)
	}

	//want := [][]int{{1, 2}, {2, 1}}
	//out := permutationsLoops([]int{1, 2})
	//t.Logf("OUT: %v\n", out)
	//if !reflect.DeepEqual(want, out) {
	//	t.Fatalf(`want %v == %v`, want, out)
	//}
	//want2 := [][]int{{1, 2, 3}, {2, 1, 3}, {3, 2, 1}, {1, 3, 2}, {2, 3, 1}, {3, 1, 2}}
	//out2 := permutationsLoops([]int{1, 2, 3})
	//t.Logf("OUT: %v\n", out2)
	//if !reflect.DeepEqual(want2, out2) {
	//	t.Fatalf(`want %v == %v`, want2, out2)
	//}
	////for _, r1 := range want {
	////	for _, r2 := range out {
	////		if !reflect.DeepEqual(r1, r2) {
	////			t.Fatalf(`want %v == %v`, r1, r2)
	////		}
	////	}
	////}
}
