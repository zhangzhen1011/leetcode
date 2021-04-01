// 有理数：能用分数表示的叫有理数，否则是无理数
// 如何将有理数转为分数
// 题目中3部分表示：正数+小数+小数循环部分
import (
	"math"
	"strconv"
	"strings"
)

func isRationalEqual(s string, t string) bool {
	a1, b1 := convertToFraction(s)
	a2, b2 := convertToFraction(t)
	return a1*b2 == a2*b1
}

// return: 分母，分子
func convertToFraction(t string) (int, int) {
	ss := strings.Split(t, ".")
	a, _ := strconv.Atoi(ss[0])
	var bstr, cstr string
	if len(ss) > 1 {
		sss := strings.Split(ss[1], "(")
		bstr = sss[0]
		if len(sss) == 2 {
			cstr = strings.TrimSuffix(sss[1], ")")
		}
	}
	bint, _ := strconv.Atoi(bstr)
	cint, _ := strconv.Atoi(cstr)
	a1 := int(math.Pow10(len(bstr)))
	a2 := int(math.Pow10(len(bstr) + len(cstr)))
	b1 := a*int(math.Pow10(len(bstr))) + bint
	b2 := b1*int(math.Pow10(len(cstr))) + cint
	if a1 == a2 {
		return a1, b1
	}
	return a2 - a1, b2 - b1
}
