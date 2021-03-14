package main

import "fmt"

//实现抽象产品
type Product interface {
	GetName() string
	SetName(name string)
}


//产品1
type Product1 struct{
	name string
}
func (p *Product1) GetName() string {
	return p.name
}
func (p *Product1) SetName(name string) {
	p.name = name
}

//产品2
type Product2 struct{
	name string
}
func (p Product2) GetName() string {
	return p.name
}
func (p Product2) SetName(name string) {
	p.name = name
}


//抽象工厂
type AbstractFactory interface {
	Create() Product
}

//实现产品1的工厂
type Product1Factory struct{}

func (pf1 Product1Factory) Create() Product {
	return &Product1{}
}

//实现产品2的工厂
type Product2Factory struct{}

func (pf2 Product2Factory) Create() Product {
	return &Product2{}
}

func main() {
	factory := &Product1Factory{}
	p1 := factory.Create()
	p1.SetName("酸奶")
	fmt.Println(p1.GetName())
}




