func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	qSort(intervals, 0, len(intervals)-1)
	var result [][]int
	var cur int = -1
	for _, val := range intervals {
		if cur == -1 {
			cur = val[1]
			result = append(result, val)
		} else {
			if val[0] <= cur {
				cur = max(val[1], cur)
				result[len(result)-1][1] = cur
			} else {
				result = append(result, val)
				cur = val[1]
			}
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func partition(a [][]int, lo, hi int) int {
	var i, j = lo, hi + 1
	for {
		for {
			i++
			if i == hi || a[i][0] > a[lo][0] {
				break
			}
		}
		for {
			j--
			if j == lo || a[j][0] < a[lo][0] {
				break
			}
		}

		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}

func qSort(a [][]int, lo, hi int) {
	if hi <= lo {
		return
	}
	p := partition(a, lo, hi)
	qSort(a, lo, p-1)
	qSort(a, p+1, hi)
}
