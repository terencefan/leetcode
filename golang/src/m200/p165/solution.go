package main

type Version struct {
	arr []int
}

func Construct(v string) *Version {
	var (
		arr []int
		num int
	)

	for i := range v {
		if v[i] == '.' {
			arr = append(arr, num)
			num = 0
		} else if v[i] >= '0' && v[i] <= '9' {
			num *= 10
			num += int(v[i] - '0')
		}
	}
	arr = append(arr, num)

	return &Version{arr}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (v *Version) get(index int) int {
	if index < len(v.arr) {
		return v.arr[index]
	} else {
		return 0
	}
}

func (v *Version) len() int {
	return len(v.arr)
}

func (v *Version) compare(other *Version) int {
	for i := 0; i < max(v.len(), other.len()); i++ {
		if v.get(i) < other.get(i) {
			return -1
		} else if v.get(i) > other.get(i) {
			return 1
		}
	}
	return 0
}

func compareVersion(version1 string, version2 string) int {
	v1, v2 := Construct(version1), Construct(version2)
	return v1.compare(v2)
}
