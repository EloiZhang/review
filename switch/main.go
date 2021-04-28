package main

import "fmt"

func main() {
	m := map[string]func(){
		"1": func1,
		"2": func2,
	}
	mReq := make(map[string]string)
	mReq["1"] = "1"
	mReq["2"] = "2"
	for req := range mReq{
		if f,ok := m[req];ok{
			f()
		}
	}
}
func func1() {
	fmt.Println("1")
}
func func2() {
	fmt.Print("2")
}