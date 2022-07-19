package main

import "container/list"

type LRUCache struct {
	Cap  int
	Keys map[int]*list.Element
	List *list.List
}

type Pair struct {
	Key   int
	Value int
}

func main() {

}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	el, flag := this.Keys[key]
	if flag == false {
		return -1
	} else {
		this.List.MoveToFront(el)
		return el.Value.(Pair).Value
	}
}

func (this *LRUCache) Put(key int, value int) {
	el, flag := this.Keys[key]
	if flag == false {
		el := this.List.PushFront(Pair{Key: key, Value: value})
		this.Keys[key] = el
	} else {
		el.Value = Pair{Key: key, Value: value}
		this.Keys[key] = el
	}
	if this.List.Len() > this.Cap {
		el := this.List.Back()
		this.List.Remove(el)
		delete(this.Keys, el.Value.(Pair).Key)
	}
}
