// 最小的k个数
// 1. topk, 堆  2. partion 分治
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	p := partition(arr, 0, len(arr)-1)
	if k == p {
		return arr[:p]
	}
	if k < p {
		return getLeastNumbers(arr[:p], k)
	} else {
		var result = arr[:p+1]
		result = append(result, getLeastNumbers(arr[p+1:], k-p-1)...)
		return result
	}
}

func partition(a []int, lo, hi int) int {
	if hi <= lo {
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

		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}

func partition_(a []int, lo, hi int) int {
	var i, j = lo, hi + 1
	for i <= j {
		for i <= hi {
			i++
			if i == hi || a[i] > a[lo] {
				break
			}
		}
		for j >= lo {
			j--
			if j == lo || a[j] < a[lo] {
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
