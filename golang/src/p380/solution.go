package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	set map[int]int
	arr []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]int),
		arr: make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; !ok {
		this.set[val] = len(this.arr)
		this.arr = append(this.arr, val)
		return true
	} else {
		return false
	}
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.set[val]; !ok {
		return false
	} else {
		p := this.set[val]
		delete(this.set, val)

		end := len(this.arr) - 1
		if p != end {
			this.set[this.arr[end]] = p
			this.arr[p] = this.arr[end]
		}
		this.arr = this.arr[:end]
		return true
	}
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	// seed := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(seed)

	if len(this.arr) > 0 {
		return this.arr[rand.Intn(len(this.arr))]
	} else {
		return 0
	}
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(0))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(-1))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
}
