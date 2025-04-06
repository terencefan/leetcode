package p415

type Number struct {
	nums []int
}

func NewNumber(s string) Number {
	nums := make([]int, len(s))
	for i := range len(s) {
		nums[i] = int(s[len(s)-i-1] - '0')
	}
	return Number{nums: nums}
}

func Add(n1, n2 Number) Number {
	nums := make([]int, 0)

	d := 0
	for i, j := 0, 0; i < len(n1.nums) || j < len(n2.nums); i, j = i+1, j+1 {
		num := d
		if i < len(n1.nums) {
			num += n1.nums[i]
		}
		if j < len(n2.nums) {
			num += n2.nums[j]
		}
		nums = append(nums, num%10)
		d = num / 10
	}
	if d > 0 {
		nums = append(nums, d)
	}
	return Number{nums: nums}
}

func (n Number) toString() string {
	bytes := make([]byte, len(n.nums))
	for i := range len(n.nums) {
		bytes[i] = byte(n.nums[len(n.nums)-i-1] + '0')
	}
	return string(bytes)
}

func addStrings(num1 string, num2 string) string {
	n1, n2 := NewNumber(num1), NewNumber(num2)
	return Add(n1, n2).toString()
}
