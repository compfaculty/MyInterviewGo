package ADT

const Size = 10

type HashTable struct {
	array [Size]*LinkedList
}

func (t *HashTable) Insert(key string) {
	index := hash(key)
	t.array[index].Insert(key)
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % Size
}

func (t *HashTable) Search(key string) bool {
	index := hash(key)
	ret := t.array[index].Contains(key)
	return ret
}

func (t *HashTable) Delete(key string) *Node {
	index := hash(key)
	ret := t.array[index].Delete(key)
	return ret
}

func NewHashTable() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &LinkedList{}
	}
	return result
}
