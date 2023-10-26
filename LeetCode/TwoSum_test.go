package leetcode

func TwoSum(numbers []int, target int) []int {
	ret := make([]int, 2)
	store := make(map[int]int)
	for i, item := range numbers {
		if _, ok := store[item]; ok {
			continue
		}
		store[item] = i
	}
	for start, val := range numbers {
		if end, ok := store[target-val]; ok {
			if start == end {
				continue
			}
			ret[0], ret[1] = start, end
			break
		}
	}
	return ret
}
