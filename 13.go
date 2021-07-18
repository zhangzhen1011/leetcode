import "strings"

func romanToInt(s string) int {
	var sum = 0
	for k, v := range m2 {
		n := len(s)
		s = strings.ReplaceAll(s, k, "")
		if len(s) != n {
			sum += v
		}
	}

	for _, v := range s {
		sum += m[byte(v)]
	}
	return sum
}

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var m2 = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}
