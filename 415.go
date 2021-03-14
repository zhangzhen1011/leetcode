import "strconv"

func addStrings(num1 string, num2 string) string {
	var c int
	var result string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0') // byte to int
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		tmp := (x + y + c) % 10
		c = (x + y + c) / 10
		result = strconv.Itoa(tmp) + result
	}
	if c == 1 {
		result = "1" + result
	}
	return result
}
