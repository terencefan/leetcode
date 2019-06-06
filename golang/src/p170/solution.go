package main

type TwoSum struct {
	m     map[int]int
	cache map[int]bool
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		make(map[int]int),
		make(map[int]bool),
	}
}

/** Add the number to an internal data structure.. */
func (s *TwoSum) Add(number int) {
	s.m[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (s *TwoSum) Find(value int) bool {
	if s.cache[value] {
		return true
	}

	for k, v := range s.m {
		delta := value - k
		if (delta == k && v > 1) || (delta != k && s.m[delta] > 0) {
			s.cache[value] = true
			return true
		}
	}
	return false
}
