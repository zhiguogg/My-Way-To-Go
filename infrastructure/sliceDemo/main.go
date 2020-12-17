package main

import "fmt"

func declareSlice()  {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	fmt.Println(d)
	 //fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}

func easySlice()  {
	a := []int{0,1,2,3,4,5}
	s := a[1:3]  //  1 2
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))   //s:[1 2] len(s):2 cap(s):5
	s2 := s[0:5]
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))  //s2:[1 2 3 4 5] len(s2):5 cap(s2):5

	s3 := s[0:3]
	fmt.Printf("s3:%v len(s3):%v cap(s3):%v\n", s3, len(s3), cap(s3))
	s3 = s3[0:cap(s3)]
	fmt.Printf("s3:%v len(s3):%v cap(s3):%v\n", s3, len(s3), cap(s3))
}

func changeSlice()  {
	a := []int{1,2,3}
	a[1] = 12
	fmt.Println(a)

	b := []interface{}{1,"空接口可以存储任何数据",false}
	fmt.Println(b)
}

func sliceAssignment()  {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
}

func copySlice()  {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]

	// copy()复制切片
	a1 := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a1)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a1) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a1) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}

func deleteSlice()  {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}

func main()  {
	//easySlice()
	//changeSlice()
}
