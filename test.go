package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a[3:])
	fmt.Println(append(a[:1], a[2:]...), a)
}
