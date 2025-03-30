package main

type Version struct {
	nums []int
}

func Build(version string) Version {
	var v = Version{nums: []int{}}

	num := 0
	for _, c := range version {
		if c == '.' {
			v.nums = append(v.nums, num)
			num = 0
		} else {
			num *= 10
			num += int(c - '0')
		}
	}
	if num != 0 {
		v.nums = append(v.nums, num)
	}
	return v
}

func compareVersion(version1 string, version2 string) int {
	v1, v2 := Build(version1), Build(version2)

	k := min(len(v1.nums), len(v2.nums))
	for i := range k {
		if v1.nums[i] < v2.nums[i] {
			return -1
		} else if v1.nums[i] > v2.nums[i] {
			return 1
		}
	}

	if len(v1.nums) > len(v2.nums) {
		for i := range len(v1.nums) - k {
			if v1.nums[i+k] > 0 {
				return 1
			}
		}
		return 0
	} else if len(v1.nums) < len(v2.nums) {
		for i := range len(v2.nums) - k {
			if v2.nums[i+k] > 0 {
				return -1
			}
		}
	}
	return 0
}
