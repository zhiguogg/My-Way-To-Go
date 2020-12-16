package main

import "fmt"

type dog struct {}

type cat struct {}

type Sayer interface {
	Say()
}

func (d dog)Say()  {
	fmt.Println("汪汪汪")
}

func (c cat)Say()  {
	fmt.Println("喵喵喵")
}

type Mover interface {
	move()
}

func (d dog) move() {
	fmt.Println("狗会动")
}


func dataDemo()  {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()
}

//func (d *dog) move() {
//	fmt.Println("狗会动")
//}

//func pointDemo()  {
//	var x Mover
//	var wangcai = dog{} // 旺财是dog类型
//	x = wangcai         // x不可以接收dog类型
//	var fugui = &dog{}  // 富贵是*dog类型
//	x = fugui           // x可以接收*dog类型
//}

func interfaceVariable()  {
	var s Sayer
	c := cat{}
	d := dog{}
	s = c
	s.Say()
	s = d
	s.Say()
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func emptyInterface()  {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}

func main()  {
	//interfaceVariable()

	// 可以看到 洗衣机只实现了一个接口wash，但是它的嵌套结构体实现了甩干
	h := haier{}
	h.dry()

}