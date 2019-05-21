package main

type RandomizedSet struct {
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {

}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {

}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {

}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {

}

func main() {
	obj := Constructor()
	_ := obj.Insert(2)
	_ := obj.Remove(2)
	_ := obj.GetRandom()
}
