import "strconv"

// 逆波兰表达式，后序表达式，结合栈处理
func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, tmp := range tokens {
		switch tmp {
		case "+":
			o1 := stack[len(stack)-1]
			o2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, o1+o2)
		case "-":
			o1 := stack[len(stack)-1]
			o2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, o2-o1)
		case "/":
			o1 := stack[len(stack)-1]
			o2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, o2/o1)
		case "*":
			o1 := stack[len(stack)-1]
			o2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, o2*o1)
		default:
			a, _ := strconv.Atoi(tmp)
			stack = append(stack, a)
		}
	}
	return stack[0]
}
