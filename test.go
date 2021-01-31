package main

func main() {
	m := make([]int, 0)
	for i := 0; i < 14; i++ {
		m = append(m, i)
	}
	var t []int
	copy(t, m)
	println(t)
}
