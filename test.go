package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "abc"
	fmt.Println(reflect.TypeOf(s[:1]))
}
