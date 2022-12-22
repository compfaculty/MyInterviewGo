package myinterviewgo

import (
	"reflect"
	"testing"
)

func permutationsLoops(arr []int) [][]int {
	permutations := [][]int{arr}
	l := len(arr)
	for i := 0; i < l-1; i++ {
		rows := permutations
		for j := i + 1; j < l; j++ {
			for _, row := range rows {
				permut := make([]int, len(row))
				copy(permut, row)
				permut[i], permut[j] = permut[j], permut[i]
				permutations = append(permutations, permut)
			}
		}
	}
	return permutations
}

func permutationsHeaps(arr []int) [][]int {
	permutations := [][]int{arr}
	l := len(arr)
	c := make([]int, l)
	for i := 1; i < l; {
		if c[i] < i {
			if i%2 == 0 {
				arr[0], arr[i] = arr[i], arr[0]
			} else {
				arr[c[i]], arr[i] = arr[c[i]], arr[i]
			}
			permutations = append(permutations, arr)
			tmp := make([]int, l)
			copy(tmp, arr)
			arr = tmp
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return permutations
}

func TestPermutationsLoops(t *testing.T) {
	want := [][]int{{1, 2}, {2, 1}}
	out := permutationsLoops([]int{1, 2})
	t.Logf("OUT: %v\n", out)
	if !reflect.DeepEqual(want, out) {
		t.Fatalf(`want %v == %v`, want, out)
	}
	want2 := [][]int{{1, 2, 3}, {2, 1, 3}, {3, 2, 1}, {1, 3, 2}, {2, 3, 1}, {3, 1, 2}}
	out2 := permutationsLoops([]int{1, 2, 3})
	t.Logf("OUT: %v\n", out2)
	if !reflect.DeepEqual(want2, out2) {
		t.Fatalf(`want %v == %v`, want2, out2)
	}
	//for _, r1 := range want {
	//	for _, r2 := range out {
	//		if !reflect.DeepEqual(r1, r2) {
	//			t.Fatalf(`want %v == %v`, r1, r2)
	//		}
	//	}
	//}
}

func TestPermutationsHeaps(t *testing.T) {
	want := [][]int{{1, 2}, {2, 1}}
	out := permutationsLoops([]int{1, 2})
	t.Logf("OUT: %v\n", out)
	if !reflect.DeepEqual(want, out) {
		t.Fatalf(`want %v == %v`, want, out)
	}
	want2 := [][]int{{1, 2, 3}, {2, 1, 3}, {3, 2, 1}, {1, 3, 2}, {2, 3, 1}, {3, 1, 2}}
	out2 := permutationsLoops([]int{1, 2, 3})
	t.Logf("OUT: %v\n", out2)
	if !reflect.DeepEqual(want2, out2) {
		t.Fatalf(`want %v == %v`, want2, out2)
	}
	//for _, r1 := range want {
	//	for _, r2 := range out {
	//		if !reflect.DeepEqual(r1, r2) {
	//			t.Fatalf(`want %v == %v`, r1, r2)
	//		}
	//	}
	//}
}

//var table_permutations = []struct {
//	input []int
//}{
//	{input: []int{1, 2, 3}},
//	{input: []int{1, 2}},
//}
//
//func BenchmarkPrimeNumbersTable(b *testing.B) {
//	for _, v := range table {
//		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				primeNumbers(v.input)
//			}
//		})
//	}
//}
