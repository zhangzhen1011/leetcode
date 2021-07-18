import "strconv"

// 大数乘法
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m := len(num1)
	n := len(num2)
	var ans = make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ans[i+j+1] += x * y
		}
	}
	//var c int
	for i := m + n - 1; i > 0; i-- {
		ans[i-1] += ans[i] / 10
		ans[i] = ans[i] % 10
		//c = tmp
	}
	if ans[0] == 0 {
		ans = ans[1:]
	}
	var ret string
	for i := 0; i < len(ans); i++ {
		ret += strconv.Itoa(ans[i])
	}
	return ret
}
