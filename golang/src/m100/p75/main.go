package p75

const (
	RED = iota
	WHITE
	BLUE
)

func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for i := 0; i <= r; i++ {
		if i > r {
			break
		}
		for nums[i] == BLUE && i < r {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
		if nums[i] == RED {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
}
