package main

import "fmt"

// sunday 算法
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	var i, nl, hl = 0, len(needle), len(haystack)
	if nl > hl {
		return -1
	}

	for i < hl {
		tmp := i
		j := 0
		for ; j < nl && tmp < hl; j++ {
			if haystack[tmp] != needle[j] {
				break
			}
			tmp++
		}
		if j == nl {
			return i
		} else { // 失配时，更新i
			if i+nl < hl {
				for j = nl - 1; j >= 0; j-- {
					if haystack[i+nl] == needle[j] { //  从右到左，第一个匹配时
						i += nl - j
						break
					}
				}
				if j == -1 {
					i += nl + 1
				}
			} else {
				return -1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("mississippia", "sippia"))
}
