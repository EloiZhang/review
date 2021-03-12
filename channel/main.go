package main

import (
	"fmt"
	"sync"
)

/*
	测试有缓存和无缓存通道
	如果不设置缓存数，并发写入会有问题
*/


func test1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- 1
}

func test2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- 2
}

func main() {
	var wg sync.WaitGroup

	//succ
	ch := make(chan int, 200)
	//failed
	//ch := make(chan int, 200)

	wg.Add(2)
	go test1(ch, &wg)
	go test2(ch, &wg)
	wg.Wait()
	close(ch)
	for c := range ch{
		fmt.Println(c)
	}
}



