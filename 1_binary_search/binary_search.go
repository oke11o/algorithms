package binary_search

func search(l []int, item int) (int, int, bool) {
	low := 0
	high := len(l) - 1

	var guess int
	var c int
	for low <= high {
		c++
		mid := (low + high) / 2
		guess = l[mid]
		if guess == item {
			return mid, c, true
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, c, false
}
func search2(l []int, item int) (int, int, bool) {
	low := 0
	high := len(l) - 1

	var guess int
	var c int
	for low <= high {
		c++
		mid := low + high
		guess = l[mid]
		if guess == item {
			return mid, c, true
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, c, false
}
