package main

import (
	"fmt"
	"time"
)

func main() {
	startedAt := time.Now()
	defer func() {
		fmt.Println(time.Since(startedAt))
	}()

	time.Sleep(time.Second)

	//fmt.Println("test1")
	//fmt.Println(test())
	//defer fmt.Println("test2")
	//defer fmt.Println("test3")
	//defer fmt.Println("test4")
}

func test() string {

	go func() {
		fmt.Println("test5")
	}()
	time.Sleep(1 * time.Second)
	return "slice"
}
