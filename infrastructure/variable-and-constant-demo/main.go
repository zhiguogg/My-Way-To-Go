package main

import (
	"flag"
	"fmt"
)

// 声明变量
func declareVariable()  {
	var s string
	var (
		age int8
		name string = "令狐冲"
	)
	fmt.Println(s,age,name)
	name , address := "蓝盈盈","华山"
	fmt.Println(name,address)
}

func declareConstant()  {
	const (
		c0 = iota     // 0
		c1
		c2 = 4
		c3 = iota     //3
	)
}

func getTheFlag() *string {
	return flag.String("name","everyone","the greeting object")
}

var block = "package"


func main()  {
	//declareVariable()

	//var name = getTheFlag()
	//flag.Parse()
	//fmt.Printf("Hello, %v!\n", *name)

	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)

}