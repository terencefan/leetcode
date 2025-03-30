package p2334

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Top() int {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

func validSubarraySize(nums []int, threshold int) int {
	a1, a2 := make([]int, len(nums)), make([]int, len(nums))

	for i := range len(nums) {
		a1[i] = len(nums)
		a2[i] = -1
	}

	var s = make(Stack, 0)
	for i, num := range nums {
		for s.Len() > 0 && nums[s.Top()] > num {
			a1[s.Pop()] = i
		}
		s.Push(i)
	}

	s = make(Stack, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]

		for s.Len() > 0 && nums[s.Top()] > num {
			a2[s.Pop()] = i
		}
		s.Push(i)
	}

	for i := range len(nums) {
		size := a1[i] - a2[i] - 1
		if nums[i] > threshold/size {
			return size
		}
	}
	return -1
}
