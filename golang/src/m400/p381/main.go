package p381

import "math/rand"

type RandomizedCollection struct {
	valIndex map[int]map[int]bool
	valArr   []int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		valArr:   make([]int, 0),
		valIndex: make(map[int]map[int]bool),
	}
}

func (c *RandomizedCollection) Insert(val int) bool {
	idx := len(c.valArr)
	c.valArr = append(c.valArr, val)

	if idxMap, ok := c.valIndex[val]; !ok {
		c.valIndex[val] = map[int]bool{idx: true}
		return true
	} else {
		idxMap[idx] = true
		return len(idxMap) == 1
	}
}

func (c *RandomizedCollection) Remove(val int) bool {
	idxMap := c.valIndex[val]
	if len(idxMap) == 0 {
		return false
	}

	var idx int
	for idx, _ = range idxMap {
		break
	}
	delete(idxMap, idx)

	var last = len(c.valArr) - 1
	if idx < last {
		var lastVal = c.valArr[last]
		var lastIdxMap = c.valIndex[lastVal]
		delete(lastIdxMap, last)
		lastIdxMap[idx] = true
		c.valArr[idx] = c.valArr[last]
	}
	c.valArr = c.valArr[:last]
	return true
}

func (c *RandomizedCollection) GetRandom() int {
	idx := rand.Intn(len(c.valArr))
	return c.valArr[idx]
}
