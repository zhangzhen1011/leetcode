import "sort"

type MyCalendarTwo struct {
	times []timePoint
}

type timePoint struct {
	start int
	end   int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{make([]timePoint, 0)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	sort.SliceStable(this.times, func(i, j int) bool {
		return this.times[i].start < this.times[j].start
	})

	tmp := make([]timePoint, 0)
	for _, tp := range this.times {
		if start >= tp.end || end <= tp.start {
			continue
		}
		tmp = append(tmp, tp)
	}

	for i := 0; i < len(tmp)-1; i++ {
		if tmp[i].end <= tmp[i+1].start {
			continue
		}
		return false
	}
	this.times = append(this.times, timePoint{start, end})

	return true
}

// 方案：a, b, c, 假设a，b相交，a，c相交，如果b，c相交，那么存在三者相交

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
