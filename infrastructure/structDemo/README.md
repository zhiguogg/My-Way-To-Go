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