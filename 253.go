import (
	"container/heap"
	"sort"
)

// 会议室II
// 至少需要多少会议室，使区间不冲突
// 优先队列保存正在进行中的会议
func minMeetingRooms(intervals [][]int) int {
	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	var ret int = 1
	q := pq{}
	for i := 0; i < len(intervals); i++ {
		if len(q) == 0 {
			heap.Push(&q, intervals[i][1])
		} else if intervals[i][0] < q[0] {
			heap.Push(&q, intervals[i][1])
			ret = max(ret, len(q))
		} else {
			heap.Pop(&q)
			heap.Push(&q, intervals[i][1])
		}

		//fmt.Println(q)
	}

	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type pq []int

func (p *pq) Less(i, j int) bool {
	return (*p)[i] < (*p)[j]
}
func (p *pq) Len() int {
	return len(*p)
}
func (p *pq) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}
func (p *pq) Push(x interface{}) {
	*p = append(*p, x.(int))
}
func (p *pq) Pop() interface{} {
	old := *p
	*p = old[:len(old)-1]
	return old[len(old)-1]
}
