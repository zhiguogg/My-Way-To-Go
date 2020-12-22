package main

import "fmt"

func pointerDemo()  {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b)
}

func pointerGet()  {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {    //传的是指针，即地址
	*x = 100
}

func modifyDemo()  {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}

func main()  {
	pointerDemo()
	pointerGet()
}