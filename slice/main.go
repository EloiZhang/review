package main

import "fmt"


func main() {
	s := []int{1,2,3,4}
	fmt.Println(s)

	Add(s)
	fmt.Println(s)
}

func Append(s []int) {
	s = append(s, 5)
}

func Add(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = s[i]+5
	}
}
