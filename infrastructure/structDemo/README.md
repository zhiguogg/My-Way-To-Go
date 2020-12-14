# struct
Go语言中**没有“类”的概念**，也不支持“类”的继承等面向对象的概念。Go语言中通过**结构体的内嵌再配合接口**比面向对象具有更高的扩展性和灵活性。

**结构体是值类型！**


## 结构体的定义

使用`type`和`struct`关键字来定义结构体，具体代码格式如下：

```go
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
}
```

其中：

- 类型名：标识自定义结构体的名称，在**同一个包内不能重复**。
- 字段名：表示结构体字段名。结构体中的**字段名必须唯一**。
- 字段类型：表示结构体字段的具体类型。


## 结构体实例化

### 通过通过声明实例化的方式
```go
    var s1 student // s1 := student{}
	s1.studentName = "lyf"
	s1.studentId = 2014213134
	fmt.Printf("%v,%p\n",s1,&s1)
```
输出结果
```
{17 lyf},0xc0000a6020
```

### 通过new方式
```go
var t1 = new(teacher)
	fmt.Printf("t1 的地址是：%p\n",t1)   //通过new初始化得到的是地址
	fmt.Println("t1的值是：",*t1)
	t1.teacherId = "2014213351"
	t1.age = 31
	t1.teacherName = "wyy"
	fmt.Printf("teacher name is: %v\n",t1.teacherName)
```
通过new我们得到是指针类型结构体
输出结果
```
t1 的地址是：0xc000056180
t1的值是： {0  }
teacher name is: wyy
```

### 取结构体地址实例化
```go
    var s2 = &student{}  // s2 := &student{}
	fmt.Printf("s2 的地址是：%p\n",s2)
	fmt.Println("s2的值是：",*s2)
```
本质和new一样
输出结果
```
s2 的地址是：0xc00000c0c0
s2的值是： { }
```

## 结构体初始化
**没有初始化的结构体，其成员变量都是对应其类型的零值。**
### 使用键值对初始化
```go
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
```
打印结果
```
t1=main.teacher{age:19, teacherId:"20187624", teacherName:"mwq"}
s1=&main.student{studentId:"2017213355", studentName:"木婉清"}
t2=&main.teacher{age:0, teacherId:"", teacherName:"王语嫣"}

```
### 使用值的列表初始化
```go
//使用值的列表初始化
	//初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
	s2 := student{
		"2019214344",
		"钟灵儿",
	}
	fmt.Printf("s2=%#v\n",s2)
```
打印结果
```
s2=main.student{studentId:"2019214344", studentName:"钟灵儿"}
```
使用这种格式初始化时，需要注意：

1. 必须初始化结构体的**所有字段**。
2. 初始值的填充顺序必须与字段在结构体中的声明**顺序一致**。
3. 该方式**不能**和键值初始化方式**混用**。

## 匿名结构体和空结构体
```go
var p1 struct{
		name string; age int //注意如果字段在同一行，需要分号分隔
	}
	p1.name = "李秋水"
	p1.age = 34
	fmt.Println(p1)

	var v struct{}
	fmt.Println(unsafe.Sizeof(v))  // 0
```

## 构造函数
**Go语言的结构体没有构造函数，我们可以自己实现**。  **因为`struct`是值类型**，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是**结构体指针类型**。
```go
func NewTeacher(age int,teacherId,teacherName string) *teacher  {
	return &teacher{
		teacherId: teacherId,
		teacherName: teacherName,
		age: age,
	}
}

func main()  {
    t := NewTeacher(21,"2016213321","黄蓉")
 	fmt.Printf("%#v\n", t) //&main.teacher{age:21, teacherId:"2016213321", teacherName:"黄蓉"}
}
```

## 方法和接收者
Go语言中的`方法（Method）`是一种**作用于特定类型变量的函数**。这种特定类型变量叫做`接收者（Receiver）`。接收者的概念就类似于其他语言中的`this`或者 `self`。

方法的定义格式如下：

```go
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

其中，

- 接收者变量：接收者中的参数变量名在命名时，官方建议使用**接收者类型名称首字母的小写**，而不是`self`、`this`之类的命名。例如，`Person`类型的接收者变量应该命名为 `p`，`Connector`类型的接收者变量应该命名为`c`等。
- 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
- 方法名、参数列表、返回参数：具体格式与函数定义相同。

```go
// 方法 老师有着教学的方法
func (t teacher) teach (course string) string {
	fmt.Println(t.teacherName,"老师正在上的课是：",course)
	return course
}
func main() {
    t := NewTeacher(21,"2016213321","黄蓉")
 	fmt.Println(t.teach("打狗棒法"))
}

```
打印结果
```
黄蓉 老师正在上的课是： 打狗棒法
打狗棒法
```
方法与函数的区别是，**函数不属于任何类型，方法属于特定的类型。**

### 指针类型的接收者
指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，**在方法结束后，修改都是有效的**。这种方式就十分接近于其他语言中面向对象中的`this`或者`self`。
```go
// 设置年龄 使用指针接收者
func (t *teacher) SetAge(age int) {
	t.age =age
}
func main() {
    t := NewTeacher(21,"2016213321","黄蓉")
 	t.SetAge(17)
 	fmt.Println(t)  //&{17 2016213321 黄蓉}
}
```
### 值类型的接收者
当方法作用于值类型接收者时，Go语言会在代码运行时将**接收者**的值**复制一份**。在值类型接收者的方法中可以获取接收者的成员值，但**修改操作只是针对副本，无法修改接收者变量本身。**

### 什么时候应该使用指针类型接收者

1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

### 任意类型添加方法
在Go语言中，接收者的类型可以是**任何类型**，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的`int`类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

这里为什么不能直接对int基本类型进行添加方法而是要用type关键字定义新的自定义类型？
因为 **非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。**即：因为int不是我们自己的包 故我们需要在自己的包中定义一个类型
```go
type Myint int
// 任意类型添加方法
func (m *Myint )addMethod(name string) string  {
	return name
}
func main() {
    var m Myint
 	fmt.Println(m.addMethod("沐剑屏"))  // 沐剑屏
}
```

### 结构体的匿名字段
结构体允许其成员字段在声明时**没有字段名而只有类型**，这种没有名字的字段就称为匿名字段。
```go
// 班级结构体
type class struct {
	string  //班级号
	int    // 班级人数
}
func main() {
 	c1 := class{
 		"三年二班",
 		52,
 	}
 	fmt.Println(c1) //{三年二班 52}
}
```
**注意：**这里匿名字段的说法并不代表没有字段名，而是**默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个**。

## 嵌套结构体
一个结构体中可以嵌套包含另一个结构体或结构体指针，就像下面的示例代码那样。
```go
//Address 地址结构体
type Address struct {
	Province string
	City string
}

type User struct {
	name string
	age int
	address *Address  //address Address
}
func main() {
 	u1 := User{
 		name: "郭靖",
 		age: 22,
 		address: &Address{
 			City: "蒙古",
 			Province: "蒙古",
 		},
 	}
 	fmt.Println(u1)  //{郭靖 22 0xc00000c080}
 	fmt.Println(u1.address.Province)  //蒙古
}
```
### 嵌套匿名字段
```go
//Address 地址结构体
type Address struct {
	Province string
	City string
}


type user struct {
	name string
	Address
}
func main() {
        var u2 user
    	u2.name = "杨过"
    	u2.Address.City = "蒙古"  // 匿名字段默认使用类型名作为字段名
        u2.Province = "蒙古"         // 匿名字段可以省略
}
```
当访问结构体成员时会**先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。**
### 嵌套结构体的字段名冲突
嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要**通过指定具体的内嵌结构体字段名。**
```go
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
type user struct {
	name string
	Address
	*Email 
}
func main() {
 u3 := user{
 		name: "郭襄",
 		Address: Address{
 			City: "峨眉",
 			Province: "四川",
 		},
 		Email: &Email{
 			city: "西安",
 		},
 	}
 	fmt.Println(u3)
 	u3.Address.City = "峨眉1"
 	u3.Email.city = "西安1"
 	fmt.Println(u3)
}
```

## 结构体的“继承”
Go语言中没有继承，但是Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
```go
type Animal struct {
	name string
}

func (a Animal)move()  {
	fmt.Println(a.name,"can move")
}

type person struct {
	thing string
	*Animal //这里必须是匿名，否则无法实现“继承”
}

func (p *person)think()  {
	fmt.Println(p.name,"is thinking: ",p.thing)
}
func main() {
 p := &person{
 		"math",
 		&Animal{
 			name: "丘处机",
 		},
 	}
 	p1 := person{
 		thing: "english",
 		Animal: &Animal{
 			name: "柯正恶",
 		},
 	}
 
 	p.move()   // 丘处机 can move
 	p1.think()  //柯正恶 is thinking:  english
}
```

### 结构体字段的可见性

结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

## 结构体与JSON序列化
JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。易于人阅读和编写。同时也易于机器解析和生成。JSON键值对是用来保存JS对象的一种方式，键/值对组合中的键名写在前面并用双引号`""`包裹，使用冒号`:`分隔，然后紧接着值；多个键值之间使用英文`,`分隔。
```go
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
func main() {
 c := &Class{
 		Title: "101",
 		Students: make([]*Student,0,200),
 	}
 	for i := 0; i<10; i++ {
 		stu := &Student{
 			Name:   fmt.Sprintf("stu%02d", i),
 			Gender: "男",
 			ID:     i,
 		}
 		c.Students = append(c.Students, stu)
 	}
 	//序列化
 	data,err := json.Marshal(c)
 	if err != nil {
 		fmt.Printf("JSON err %v\n:",err)
 		return
 	}
 	fmt.Printf("json:%s\n",data)
 
 	// 反序列化  ``中所有的字符原样输出 可以换行
 	str := `{"Title":"101","Students":
 [{"ID":0,"Gender":"男","Name":"stu00"},
 {"ID":1,"Gender":"男","Name":"stu01"},
 {"ID":2,"Gender":"男","Name":"stu02"},
 {"ID":3,"Gender":"男","Name":"stu03"},
 {"ID":4,"Gender":"男","Name":"stu04"},
 {"ID":5,"Gender":"男","Name":"stu05"},
 {"ID":6,"Gender":"男","Name":"stu06"},
 {"ID":7,"Gender":"男","Name":"stu07"},
 {"ID":8,"Gender":"男","Name":"stu08"},
 {"ID":9,"Gender":"男","Name":"stu09"}]}`
 
 	c1 := &Class{}
 	err = json.Unmarshal([]byte(str),c1)
 	if err != nil {
 		fmt.Println("json unmarshal failed!")
 		return
 	}
 	fmt.Printf("%#v\n", c1)
}
```

### 结构体标签（Tag）
`Tag`是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 `Tag`在结构体字段的后方定义，由一对**反引号**包裹起来，具体的格式如下：
```go
`key1:"value1" key2:"value2"`
```
结构体tag由一个或多个键值对组成。**键与值使用冒号分隔，值用双引号括起来**。同一个结构体字段可以设置多个键值对tag，**不同的键值对之间使用空格分隔**。

**注意事项：** 为结构体编写`Tag`时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如**不要在key和value之间添加空格**。
```go
//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "段誉",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}
```

## 结构体和方法补充知识点
因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。我们来看下面的例子：
```go
type Person struct {
	name string
	age int8
	dreams []string
}

func (p *Person)SetDreams(dreams []string )  {
	p.dreams = dreams
}
func main() {
 p := Person{
 		name: "张无忌",
 		age: 21,
 	}
 	dreams := []string{"九阳神功","乾坤大挪移","太极剑法","太极拳法"}
 	p.SetDreams(dreams)
 
 	dreams[1] = "七伤拳"  //此时 你改变了dreams切片就等于改变了p.dreams
 	fmt.Println(p.dreams)  //[九阳神功 七伤拳 太极剑法 太极拳法]
    
}
```

正确的做法是在方法中**使用传入的slice的拷贝进行结构体赋值。**
```go
func (p *Person)SetDreams(dreams []string)  {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)    // copy函数
}
```
同样的问题也存在于**返回值slice和map**的情况，在实际编码过程中一定要注意这个问题。