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

type Myint int
// 任意类型添加方法
func (m *Myint )addMethod(name string) string  {
	return name
}

// 班级结构体
type class struct {
	string  //班级号
	int    // 班级人数
}

//Address 地址结构体
type Address struct {
	Province string
	City string
}

// Email
type Email struct {
	city string //发送者地址
	toCity string //接收者地址
}



type User struct {
	name string
	age int
	address *Address  //address Address
}

type user struct {
	name string
	Address
	*Email
}


type Animal struct {
	name string
}

func (a Animal)move()  {
	fmt.Println(a.name,"can move")
}

type person struct {
	thing string
	*Animal    //这里必须是匿名，否则无法实现“继承”
}

func (p *person)think()  {
	fmt.Println(p.name,"is thinking: ",p.thing)
}

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

type Person struct {
	name string
	age int8
	dreams []string
}

// 无法改变
func (p *Person)SetDreams(dreams []string )  {
	p.dreams = dreams
}

//正确做法
func (p *Person)SetDreams1(dreams []string)  {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)    // copy函数
}

func main()  {
	
	//define()
	//initialize()
	//t := NewTeacher(21,"2016213321","黄蓉")
	//t.SetAge(17)
	//fmt.Println(t)  //&{17 2016213321 黄蓉}
	//fmt.Println(t.teach("打狗棒法"))
	//fmt.Printf("%#v\n", t)
	//var m Myint
	//fmt.Println(m.addMethod("沐剑屏"))

	//c1 := class{
	//	"三年二班",
	//	52,
	//}
	//fmt.Println(c1)

	//u1 := User{
	//	name: "郭靖",
	//	age: 22,
	//	address: &Address{
	//		City: "蒙古",
	//		Province: "蒙古",
	//	},
	//}
	//fmt.Println(u1)
	//fmt.Println(u1.address.Province)

	//var u2 user
	//u2.name = "杨过"
	//u2.Province = "蒙古"
	//u2.Address.City = "蒙古"

	//u3 := user{
	//	name: "郭襄",
	//	Address: Address{
	//		City: "峨眉",
	//		Province: "四川",
	//	},
	//	Email: &Email{
	//		city: "西安",
	//	},
	//}
	//fmt.Println(u3)
	//u3.Address.City = "峨眉1"
	//u3.Email.city = "西安1"
	//fmt.Println(u3)

	//p := &person{
	//	"math",
	//	&Animal{
	//		name: "丘处机",
	//	},
	//}
	//p1 := person{
	//	thing: "english",
	//	Animal: &Animal{
	//		name: "柯正恶",
	//	},
	//}
	//
	//p.move()
	//p1.think()

//	c := &Class{
//		Title: "101",
//		Students: make([]*Student,0,200),
//	}
//	for i := 0; i<10; i++ {
//		stu := &Student{
//			Name:   fmt.Sprintf("stu%02d", i),
//			Gender: "男",
//			ID:     i,
//		}
//		c.Students = append(c.Students, stu)
//	}
//	//序列化
//	data,err := json.Marshal(c)
//	if err != nil {
//		fmt.Printf("JSON err %v\n:",err)
//		return
//	}
//	fmt.Printf("json:%s\n",data)
//
//	// 反序列化  ``中所有的字符原样输出 可以换行
//	str := `{"Title":"101","Students":
//[{"ID":0,"Gender":"男","Name":"stu00"},
//{"ID":1,"Gender":"男","Name":"stu01"},
//{"ID":2,"Gender":"男","Name":"stu02"},
//{"ID":3,"Gender":"男","Name":"stu03"},
//{"ID":4,"Gender":"男","Name":"stu04"},
//{"ID":5,"Gender":"男","Name":"stu05"},
//{"ID":6,"Gender":"男","Name":"stu06"},
//{"ID":7,"Gender":"男","Name":"stu07"},
//{"ID":8,"Gender":"男","Name":"stu08"},
//{"ID":9,"Gender":"男","Name":"stu09"}]}`
//
//	c1 := &Class{}
//	err = json.Unmarshal([]byte(str),c1)
//	if err != nil {
//		fmt.Println("json unmarshal failed!")
//		return
//	}
//	fmt.Printf("%#v\n", c1)


	p := Person{
		name: "张无忌",
		age: 21,
	}
	dreams := []string{"九阳神功","乾坤大挪移","太极剑法","太极拳法"}
	p.SetDreams1(dreams)

	dreams[1] = "七伤拳"
	fmt.Println(p.dreams)
}