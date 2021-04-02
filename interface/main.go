package main

import "fmt"

type T interface {

}
func main() {
	var(
		t T
		p *T
		i1 interface{} = t
		i2 interface{} = p
	)

	fmt.Println(i1 == t,i1 == nil)
	fmt.Println(i1 == p,i2 == nil)
}
