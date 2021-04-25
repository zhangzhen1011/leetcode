package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func convertToFraction(t string) (int, int) {
	ss := strings.Split(t, ".")
	a, _ := strconv.Atoi(ss[0])
	var bstr, cstr string
	if len(ss) > 0 {
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

func main() {
	s := "    1  2"
	fmt.Println(len(strings.Split(s, " ")))
}
