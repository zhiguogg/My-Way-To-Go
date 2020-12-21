package main

import (
	"errors"
	"fmt"
)

func intSum(x,y int) int {
	return x+y
}

func intSumAll(x ...int) int  {
	sum := 0
	for _,v :=range x{
		sum += v
	}
	return sum
}

func calc(x,y int ) (int,int)  {
	if x <y {
		return calc(y,x)
	}
	sum := x+y
	sub := x-y
	return sum,sub
}

func calc1(x,y int ) (sum , sub int)  {
	if x <y {
		return calc(y,x)
	}
	sum = x+y
	sub = x-y
	return
}

type calculation func(x,y int) (int,int)

func parameterFunc(x,y int , f func(int,int) int)  int  {
	return x+y + f(x,y)
}

func operation(x string) (func(int,int) int,error) {
	switch x {
	case "+":
		return intSum,nil
	default:
		err := errors.New("只能传加号")
		return nil,err
	}
}

/*
匿名函数
*/
func anonymousFunc() {
	addAngin := func(x, y int) int {
		return x + y
	}
	fmt.Println("匿名函数：", addAngin(11, 22))
	a := func(x, y int) int {
		return x - y
	}(29, 12)
	fmt.Println(a)
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func deferDemo()  {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func funA()  {
	fmt.Println("func A")
}

func funB()  {
	panic("panic in B")
}

func recoverB()  {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funC()  {
	fmt.Println("Func C")
}
func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}
func modifySlice(a []string) []string {
	a[0] = "gg"
	return a
}

func modifySliceArray(a [3][]string) [3][]string {
	a[0] = []string{"hello","golang"}
	return a
}

func modifySliceArray1(a [3][]string) [3][]string {
	a[0][0] = "Hello go"
	return a
}

func parameterDemo()  {
	//array1 := [3]string{"a", "b", "c"}
	//fmt.Printf("The array: %v\n", array1)
	//array2 := modifyArray(array1)
	//fmt.Printf("The modified array: %v\n", array2)
	//fmt.Printf("The original array: %v\n", array1)
	//
	//a := array1[:]
	//fmt.Printf("The slice a: %v\n",a)
	//a1 := modifySlice(a)
	//fmt.Printf("The slice a1: %v\n",a1)
	//fmt.Printf("The slice a: %v\n",a)
	//fmt.Printf("The array: %v\n", array1)

	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	fmt.Printf("The array: %v\n", complexArray1)
	complexArray2 := modifySliceArray1(complexArray1)
	fmt.Printf("The array: %v\n", complexArray2)
	fmt.Printf("The array: %v\n", complexArray1)
}

func main()  {
	//fmt.Println(intSum(1,2),intSumAll(1,2,3))
	//fmt.Println(calc(1,2))
	//sum,sub := calc1(2,1)
	//fmt.Println(sum,sub)

	//var a = calc
	//fmt.Println(a(1,2))

	//var a calculation
	//a = calc
	//fmt.Println(a)
	//fmt.Println(a(1,3))

	//fmt.Println(parameterFunc(1,2,intSum))

	//c,_ := operation("+")
	//fmt.Println(c(1,2))


	//deferDemo()

	//funA()
	//recoverB()
	//funC()

	parameterDemo()
}
