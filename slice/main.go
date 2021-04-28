package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := []int{1,2,3,4}
	//fmt.Println(s)
	//
	//Add(s)
	//fmt.Println(s)

	str := "+1+2+3+4"
	vec := strings.Split(str, "+")
	fmt.Println(len(vec))
	for _, vid := range vec {
		fmt.Printf("vid:%v\n", vid)
	}

}

func Append(s []int) {
	s = append(s, 5)
}

func Add(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = s[i]+5
	}
}
