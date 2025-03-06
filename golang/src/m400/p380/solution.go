package p380

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	indexVal []int
	valIndex map[int]int
	r        rand.Source
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		indexVal: make([]int, 0),
		valIndex: make(map[int]int),
		r:        rand.NewSource(time.Now().UnixNano()),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valIndex[val]; ok {
		return false
	} else {
		this.valIndex[val] = len(this.indexVal)
		this.indexVal = append(this.indexVal, val)
	}
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.valIndex[val]; ok {
		last := len(this.indexVal) - 1
		lastVal := this.indexVal[last]
		this.indexVal[idx], this.indexVal[last] = this.indexVal[last], this.indexVal[idx]
		this.indexVal = this.indexVal[:last]
		this.valIndex[lastVal] = idx
		delete(this.valIndex, val)
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	idx := this.r.Int63() % int64(len(this.indexVal))
	return this.indexVal[idx]
}
