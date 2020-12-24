package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PrintTest() {
	fmt.Print("大理", "段氏", "段誉")
	fmt.Println("大理", "段氏", "段誉")
	name := "北冥神功"
	fmt.Printf("段誉的武功： %s\n", name)
}

func FprintTest() {
	//向标准输出写入内容
	fmt.Fprintln(os.Stdout, "桃花岛黄老邪")

	fileObj, err := os.OpenFile("./fmt.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "黄蓉"
	//向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "黄老邪女儿是：%s", name)

}

func SprintTest() {
	story1 := fmt.Sprintln("郭靖最终娶了黄蓉，没有华筝什么事")
	name1 := "降龙十八掌"
	name2 := "打狗棒法"
	story2 := fmt.Sprintf("郭靖会%s,黄蓉会%s\n", name1, name2)
	fmt.Println(story1, story2)
}

type person struct {
	name string
	age  int32
}

func CommonSign() {
	p := person{
		"张无忌",
		22,
	}
	fmt.Printf("倚天屠龙记主角是：%v\n", p)
	fmt.Printf("倚天屠龙记主角是：%+v\n", p)
	fmt.Printf("倚天屠龙记主角是：%#v\n", p)
	fmt.Printf("倚天屠龙记主角是：%T\n", p)
	fmt.Printf("100%%\n")
	b := true
	fmt.Printf("真%t\n", b)
	fmt.Printf("%p\n", &p)
	fmt.Printf("%#p\n", &p)

}

func IntegerSign() {
	n := 1165
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%U\n", n)
	fmt.Printf("%q\n", n)
}

func stringSign()  {
	s := "雪山飞狐"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
}

func widthSign()  {
	width := 47.21
	fmt.Printf("width: %f\n",width)
	fmt.Printf("width: %11f\n",width)
	fmt.Printf("width: %.2f\n",width)
	fmt.Printf("width: %11.1f\n",width)
	fmt.Printf("width: %11.f\n",width)
}

func otherSign()  {
	s := "王语嫣"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)
}

func ScanTest()  {
	var (
		name string
		age int
		gender bool
	)
	fmt.Scan(&name,&age,&gender)
	fmt.Printf("name is: %s,age is: %d,gender is %t\n",name,age,gender)
}

func ScanfTest()  {
	var (
		name string
		age int
		gender bool
	)
	fmt.Scanf("姓名：%s 年龄：%d 性别：%t",&name,&age,&gender)
	fmt.Printf("name is: %s,age is: %d,gender is %t\n",name,age,gender)
}


func bufioDemo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容")
	text,_ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n",text)
}

func main() {
	//PrintTest()
	//FprintTest()
	//SprintTest()

	//err := fmt.Errorf("ss")
	//fmt.Println(err)

	//CommonSign()
	//IntegerSign()
	//stringSign()
	//widthSign()
	//otherSign()

	//ScanfTest()

	bufioDemo()
}
