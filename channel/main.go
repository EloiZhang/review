package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	测试有缓存和无缓存通道
	如果不设置缓存数，并发写入会有问题
*/

const zero = 0.0

func test1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- 1
}
func test2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- 2
}

func main() {

	//x := []int{1,2,3, 2,3,4,}
	/*
		测试有缓存和无缓存通道
		如果不设置缓存数，并发写入会有问题
	*/
	//var wg sync.WaitGroup
	////succ
	//ch := make(chan int, 200)
	////failed
	////ch := make(chan int, 200)
	//wg.Add(2)
	//go test1(ch, &wg)
	//go test2(ch, &wg)
	//wg.Wait()
	//close(ch)
	//for c := range ch{
	//	fmt.Println(c)
	//}

	//GetChannelMessage()

	//channel := make(chan string)
	//go worker(channel)
	//<- channel
	//fmt.Println("hah")
	//
	////只写
	//read := make(chan<- string)
	//read <- ""
	//<- read

	//c1 := make(chan string)
	////c2 := make(chan string)
	//
	//go func() {
	//	//c1 <- "c1"
	//}()
	//go func() {
	//	//c2 <- "c2"
	//}()
	//
	//
	//select {
	//case res := <- c1:
	//	fmt.Println(res)
	//default:
	//	fmt.Println("default")
	////case <-time.After(1 * time.Second):
	////	fmt.Println("out")
	//}

	//msg := "hi"
	////messages <- msg
	//select {
	//case messages <- msg:
	//	fmt.Println("sent message", msg)
	//default:
	//	fmt.Println("no message sent")
	//}
	//
	//select {
	//case msg := <-messages:
	//	fmt.Println("received message", msg)
	//case sig := <-signals:
	//	fmt.Println("received signal", sig)
	//default:
	//	fmt.Println("no activity")
	//}

	//messages := make(chan string)
	//signals := make(chan bool)
	//go func() {
	//	messages <- "jj"
	//}()
	//select {
	//case msg := <-messages:
	//	fmt.Println("received message", msg)
	//case <-time.After(2 * time.Second):
	//	fmt.Println("no message received")
	//}
	//
	//msg := "hi"
	//select {
	//case messages <- msg:
	//	fmt.Println("sent message", msg)
	//default:
	//	fmt.Println("no message sent")
	//}
	//
	//select {
	//case msg := <-messages:
	//	fmt.Println("received message", msg)
	//case sig := <-signals:
	//	fmt.Println("received signal", sig)
	//default:
	//	fmt.Println("no activity")
	//}

	//c1:=make(chan string, 2)
	////func(){
	////	time.Sleep(time.Second)
	//	c1 <-"1"
	////}()
	//fmt.Println("c1 is",<-c1)

	//queue := make(chan string, 2)
	//queue <- "one"
	//close(queue)
	//
	////queue <- "two"
	//for elem := range queue {
	//	fmt.Println(elem)
	//}
	ch := make(chan string)
	go func() {
		ch <- "ee"
		//time.Sleep(1*time.Second)
	}()
	fmt.Println(<-ch)

	close(ch)

}

func worker(channel chan string) {
	fmt.Println("test")
	time.Sleep(2 * time.Second)
	fmt.Println("done")
	channel <- "true"
}

//场景，在并发函数中写channel，在主函数读出来
func GetChannelMessage() {
	//创建chan变量
	channel := make(chan string)

	go func() {
		//赋值给channel
		channel <- "hello world"
		//channel <- "test"

	}()

	//time.Sleep(5*time.Second)
	//直接取出来
	//fmt.Println(<-channel)

	//第二次取就会失败
	//fmt.Println(<-channel)

	for res := range channel {
		fmt.Println(res)
	}

}
