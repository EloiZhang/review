package main

import "fmt"


/*
https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/
Go 语言在编译期间对代码进行类型检查，上述代码总共触发了三次类型检查：
将 *RPCError 类型的变量赋值给 error 类型的变量 rpcErr；
将 *RPCError 类型的变量 rpcErr 传递给签名中参数类型为 error 的 AsErr 函数；
将 *RPCError 类型的变量从函数签名的返回值类型为 error 的 NewRPCError 函数中返回；
*/

func main() {
	//var rpcErr error
	//rpcErr = NewRpcError()
	//err := Assr(rpcErr)
	//println(err)


	type test struct {}
	//v := test{}
	//Print(v)
}

func Print(v interface{}) {
	println(v)
}



func Assr(err error) error {
	return err
}
//声明interface
type error interface {
	Error() string
}

type RpcError struct {
	code int
	msg  string
}

func (r *RpcError) Error() string {
	return fmt.Sprintf("%s, %d", r.msg, r.code)
}

func NewRpcError() error {
	return &RpcError{
		code: 0,
		msg:  "ddd",
	}
}
