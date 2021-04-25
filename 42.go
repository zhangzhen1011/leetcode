// 接雨水
//1 . 每个位置左右的最大高度，求差
func trap(height []int) int {
	var leftMost = make([]int, len(height))
	var rightMost = make([]int, len(height))

	for i := 1; i < len(height); i++ {
		leftMost[i] = max(leftMost[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMost[i] = max(rightMost[i+1], height[i+1])
	}

	var ret int
	for i := 0; i < len(height); i++ {
		h := min(leftMost[i], rightMost[i]) - height[i]
		ret += max(0, h)
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// 2. 单调栈（减）, 一次扫描
func trap(height []int) int {
	var ret int
	var stack = []ele{}
	for i := 0; i < len(height); i++ {
		e := ele{i, height[i]}
		for len(stack) > 0 && stack[len(stack)-1].height < e.height {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] //pop
			if len(stack) > 0 {
				tmp := stack[len(stack)-1]
				ret += (min(tmp.height, e.height) - top.height) * (i - tmp.idx - 1) // i-tmp.idx-1 NOTICE
			}
		}
		stack = append(stack, e)
	}
	return ret
}

type ele struct {
	idx    int // bar的位置
	height int
}
