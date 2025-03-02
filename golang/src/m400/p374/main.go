package p374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	return 0
}

const (
	LESS = -1
	MORE = 1
	SAME = 0
)

func guessNumber(n int) int {
	l, r := 1, 1<<31-1

	for l < r {
		mid := l + (r-l)/2
		switch guess(mid) {
		case LESS:
			r = mid
		case MORE:
			l = mid + 1
		case SAME:
			return mid
		}
	}
	return l
}
