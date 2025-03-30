package p981

type TimeValue struct {
	value     string
	timestamp int
}

type TimeMap struct {
	m map[string][]TimeValue
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]TimeValue)}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := t.m[key]; !ok {
		t.m[key] = make([]TimeValue, 0)
	}
	t.m[key] = append(t.m[key], TimeValue{value: value, timestamp: timestamp})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	if arr, ok := t.m[key]; !ok {
		return ""
	} else {

		l, r := 0, len(arr)

		for l < r {
			mid := l + (r-l)/2
			if arr[mid].timestamp > timestamp {
				r = mid
			} else {
				l = mid + 1
			}
		}

		if l == 0 {
			return ""
		} else {
			return arr[l-1].value
		}
	}
}
