// 验证二叉树前序序列化
import "strings"

func isValidSerialization(preorder string) bool {
	s := strings.Split(preorder, ",")
	var diff int = 1
	for _, node := range s {
		diff -= 1
		if diff < 0 {
			return false
		}
		if node != "#" {
			diff += 2
		}
	}
	return diff == 0
}

// TODO, 更多做法
