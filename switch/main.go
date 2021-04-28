package main

import "fmt"

func main() {
	m := map[string]func(str string) string {
		"1": func1,
		"2": func2,
	}
	str := "string"
	mReq := make(map[string]string)
	mReq["1"] = "1"
	mReq["2"] = "2"
	for req := range mReq{
		if f,ok := m[req];ok{
			res := f(str)
			fmt.Println("res", res)
		}
	}
}
func func1(str string) string {
	fmt.Println("func1", str)
	return str
}
func func2(str string) string {
	fmt.Println("func2", str)
	return str
}