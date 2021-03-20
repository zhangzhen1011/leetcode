func partition(a []int, lo, hi int) int {
	if hi <= lo { // 这里可以放到外层
		return 0
	}
	var i, j = lo, hi + 1
	for {
		for {
			i++
			if i == hi || a[i] > a[lo] {
				break
			}
		}
		for {
			j--
			if j == lo || a[j] < a[lo] {
				break
			}
		}

		if i >= j { // 边界要加”=“
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}
