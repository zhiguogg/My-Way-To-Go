// struct define demo
package main

import (
	"fmt"
	"unsafe"
)

// student struct
type student struct {
	studentId string
	studentName string
}

// teacher struct
type teacher struct {
	age int
	teacherId ,teacherName string
}

// define ways
func define()  {
	// 通过声明实例化
	var s1 student   // s1 := student{}
	s1.studentName = "lyf"
	s1.studentId = "2014213134"
	fmt.Printf("%v,%p\n",s1,&s1)

	// 通过new方式实例化
	var t1 = new(teacher)
	fmt.Printf("t1 的地址是：%p\n",t1)   //通过new初始化得到的是地址
	fmt.Println("t1的值是：",*t1)
	t1.teacherId = "2014213351"
	t1.age = 31
	t1.teacherName = "wyy"
	//虽然是指针类型结构体，但是还是使用 t1.teacherName 实际上是(*t1).teacherName，不过go存在语法糖
	fmt.Printf("teacher name is: %v\n",t1.teacherName)

	// 取结构体地址实例化
	var s2 = &student{}  // s2 := &student{}
	fmt.Printf("s2 的地址是：%p\n",s2)
	fmt.Println("s2的值是：",*s2)
}

// initialize ways
func initialize()  {
	//键值对方式
	t1 := teacher{
		teacherId: "20187624",
		teacherName: "mwq",
		age: 19,   //最后一个逗号不能少
	}
	var s1 = &student{
		studentId: "2017213355",
		studentName: "木婉清",
	}
	// 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	t2 := &teacher{
		teacherName: "王语嫣",
	}
	fmt.Printf("t1=%#v\n",t1)
	fmt.Printf("s1=%#v\n",s1)
	fmt.Printf("t2=%#v\n",t2)

	//使用值的列表初始化
	//初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
	s2 := student{
		"2019214344",
		"钟灵儿",
	}
	fmt.Printf("s2=%#v\n",s2)
}

// 匿名结构体和空结构体
func anonymous()  {
	var p1 struct{
		name string; age int //注意如果字段在同一行，需要分号分隔
	}
	p1.name = "李秋水"
	p1.age = 34
	fmt.Println(p1)

	var v struct{}
	fmt.Println(unsafe.Sizeof(v))  // 0
}

// 构造函数
func NewTeacher(age int,teacherId,teacherName string) *teacher  {
	return &teacher{
		teacherId: teacherId,
		teacherName: teacherName,
		age: age,
	}
}

// 方法 老师有着教学的方法
func (t teacher) teach (course string) string {
	fmt.Println(t.teacherName,"老师正在上的课是：",course)
	return course
}

// 设置年龄 使用指针接收者
func (t *teacher) SetAge(age int) {
	t.age =age
}

func main()  {
	
	//define()
	//initialize()
	t := NewTeacher(21,"2016213321","黄蓉")
	t.SetAge(17)
	fmt.Println(t)  //&{17 2016213321 黄蓉}
	//fmt.Println(t.teach("打狗棒法"))
	//fmt.Printf("%#v\n", t)

}