package main

import (
	"fmt"
	"reflect"
)

func main() {
	test([]int{1,2,3})
	test([]string{"1","3"})
}

func test(T interface{}) {

	fmt.Println(reflect.TypeOf(T))
}