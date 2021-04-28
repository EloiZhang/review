package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var m1 map[string]string
	var m2 map[string]string = map[string]string{}         // or m2 := map[string]string{}
	var m3 map[string]string = make(map[string]string, 10) // or m3 := make(map[string]string)

	//m1["1"] = "1"   // panic: assignment to entry in nil map
	m2["2"] = "2"
	m3["3"] = "3"

	for key, value := range m1 {
		fmt.Println("Key:", key, "Value:", value)
	}
	for key, value := range m2 {
		fmt.Println("Key:", key, "Value:", value)
	}
	for key, value := range m3 {
		fmt.Println("Key:", key, "Value:", value)
	}

	s1 := m1["1"]
	s2 := m2["2"]
	s3 := m3["3"]

	fmt.Printf("val=%s,%s,%s\n", s1, s2, s3)
	fmt.Printf("len=%d,%d,%d\n", len(m1), len(m2), len(m3))
	fmt.Printf("size=%d,%d,%d\n", unsafe.Sizeof(m1), unsafe.Sizeof(m2), unsafe.Sizeof(m3))

}