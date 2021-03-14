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


//工厂名
type ProductFacory struct {
}
//工厂方法
func (pf ProductFacory) Create(op int) Product {
	if op == 1 {
		return &Product1{}
	}

	if op == 2 {
		return &Product2{}
	}
	return nil
}

//工厂模式：知道工厂名、工厂方法、产品在工厂内的参数，就可以创建指定的产品
func main() {
	p1 := &Product1{}
	p1.SetName("酸奶")
	fmt.Println(p1.name)

	factory := ProductFacory{}
	product := factory.Create(1)
	product.SetName("醋")
	println(product.GetName())
}




