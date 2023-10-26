package ADT

import "time"

const Size = 10

type HashTable[T AlNum] struct {
	array [Size]*LinkedList[T]
}

func (t *HashTable[T]) Insert(key T) {
	index := hash[T](key)
	t.array[index].Append(key)
}

func hash[T AlNum](key T) int64 {
	//sum := 0
	//for _, v := range key {
	//	sum += int(v)
	//}
	//return sum % Size
	return time.Now().UnixNano() % Size
}

func (t *HashTable[T]) Search(key T) bool {
	index := hash(key)
	ret := t.array[index].Contains(key)
	return ret
}

func (t *HashTable[T]) Remove(key T) {
	index := hash(key)
	t.array[index].Remove(key)
}

func NewHashTable[T AlNum]() *HashTable[T] {
	result := &HashTable[T]{}
	for i := range result.array {
		result.array[i] = &LinkedList[T]{}
	}
	return result
}
