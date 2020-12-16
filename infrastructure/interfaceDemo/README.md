# 接口
## 接口类型
在Go语言中接口（interface）是一种类型，一种抽象的类型。

`interface`是一组`method`的集合，是`duck-type programming`的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。**不关心属性（数据），只关心行为（方法）**。

为了保护你的Go语言职业生涯，**请牢记接口（interface）是一种类型，引用类型！**。

## 为什么要使用接口
```go
type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

func main() {
	c := Cat{}
	fmt.Println("猫:", c.Say())
	d := Dog{}
	fmt.Println("狗:", d.Say())
}
```
上面的代码中定义了猫和狗，然后它们都会叫，你会发现main函数中明显有重复的代码，如果我们后续再加上猪、青蛙等动物的话，我们的代码还会一直重复下去。那我们能不能把它们当成“能叫的动物”来处理呢？

Go语言中为了解决类似上面的问题，就设计了接口这个概念。接口区别于我们之前所有的具体类型，接口是一种抽象的类型。当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么。

## 接口的定义
Go语言提倡面向接口编程

每个接口由数个方法组成，接口的定义格式如下：

```go
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

其中：

- 接口名：使用`type`将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加`er`，如有写操作的接口叫`Writer`，有字符串功能的接口叫`Stringer`等。接口名最好要能突出该接口的类型含义。
- 方法名：当方**法名首字母是大写且这个接口类型名首字母也是大写**时，这个方法可以被接口所在的包（package）之外的代码访问。
- 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

举个例子：

```go
type writer interface{
    Write([]byte) error
}
```

当你看到这个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的Write方法来做一些事情。

## 实现接口的条件

一个对象只要**全部实现了接口中的方法**，那么就实现了这个接口。换句话说，接口就是一个**需要实现的方法列表**。

我们来定义一个`Sayer`接口：

```go
// Sayer 接口
type Sayer interface {
	say()
}
```

定义`dog`和`cat`两个结构体：

```go
type dog struct {}

type cat struct {}
```

因为`Sayer`接口里只有一个`say`方法，所以我们只需要给`dog`和`cat `分别实现`say`方法就可以实现`Sayer`接口了。

```go
// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}
```

接口的实现就是这么简单，只要实现了接口中的所有方法，就实现了这个接口。

## 接口类型变量

接口类型变量能够**存储所有实现了该接口的实例**。 例如上面的示例中，`Sayer`类型的变量能够存储`dog`和`cat`类型的变量。

```go
func main() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}
```

## 值接收者和指针接收者实现接口的区别

使用值接收者实现接口和使用指针接收者实现接口有什么区别呢？接下来我们通过一个例子看一下其中的区别。

我们有一个`Mover`接口和一个`dog`结构体。

```go
type Mover interface {
	move()
}

type dog struct {}
```

### 值接收者实现接口

```go
func (d dog) move() {
	fmt.Println("狗会动")
}
```

此时实现接口的是`dog`类型：

```go
func main() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()
}
```

从上面的代码中我们可以发现，使用**值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量**。因为Go语言中有**对指针类型变量求值**的语法糖，dog指针`fugui`内部会自动求值`*fugui`,故指针类型可以赋值给值类型。
### 指针接收者实现接口

同样的代码我们再来测试一下使用指针接收者有什么区别：

```go
func (d *dog) move() {
	fmt.Println("狗会动")
}
func main() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x不可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
}
```

此时实现`Mover`接口的是`*dog`类型，所以不能给`x`传入`dog`类型的wangcai，此时x只能存储`*dog`类型的值。


## 类型与接口的关系

### 一个类型实现多个接口

一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。 例如，狗可以叫，也可以动。我们就分别定义Sayer接口和Mover接口，如下： `Mover`接口。

```go
// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}
```

dog既可以实现Sayer接口，也可以实现Mover接口。

```go
type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()
}
```

### 多个类型实现同一接口

Go语言中不同的类型还可以实现同一接口 首先我们定义一个`Mover`接口，它要求必须由一个`move`方法。

```go
// Mover 接口
type Mover interface {
	move()
}
```

例如狗可以动，汽车也可以动，可以使用如下代码实现这个关系：

```go
type dog struct {
	name string
}

type car struct {
	brand string
}

// dog类型实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}
```

这个时候我们在代码中就可以把狗和汽车当成一个会动的物体来处理了，不再需要关注它们具体是什么，只需要调用它们的`move`方法就可以了。

```go
func main() {
	var x Mover
	var a = dog{name: "旺财"}
	var b = car{brand: "保时捷"}
	x = a
	x.move()
	x = b
	x.move()
}
```

上面的代码执行结果如下：

```go
旺财会跑
保时捷速度70迈
```

**并且一个接口的方法，不一定需要由一个类型完全实现**，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

```go
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

func main() {
  // 可以看到 洗衣机只实现了一个接口wash，但是它的嵌套结构体实现了甩干
	h := haier{}
	h.dry() //甩一甩
}
```

## 接口嵌套

接口与接口间可以通过嵌套创造出新的接口。

```go
// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}
```

嵌套得到的接口的使用与普通接口一样，这里我们让cat实现animal接口：

```go
type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

func main() {
	var x animal
	x = cat{name: "花花"}
	x.move()
	x.say()
}
```

## 空接口
### 空接口的定义

空接口是指没有定义任何方法的接口。因此**任何类型都实现了空接口**。

空接口类型的变量可以存储任意类型的变量。

```go
func main() {
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
```

### 空接口的应用

#### 空接口作为函数的参数

使用空接口实现可以接收任意类型的函数参数。

```go
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
```

#### 空接口作为map的值

使用空接口实现可以保存任意值的字典。

```go
// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
```

## 类型断言
空接口可以存储任意类型的值**，那我们如何获取其存储的具体数据呢？

### 接口值

一个接口的值（简称接口值）是由`一个具体类型`和`具体类型的值`两部分组成的。这两部分分别称为接口的`动态类型`和`动态值`。
想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：

```go
x.(T)
```

其中：

- x：表示类型为`interface{}`的变量
- T：表示断言`x`可能是的类型。

该语法返回两个参数，第一个参数是`x`转化为`T`类型后的变量，第二个值是一个布尔值，若为`true`则表示断言成功，为`false`则表示断言失败。

举个例子：

```go
func main() {
	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
```

上面的示例中如果要断言多次就需要写多个`if`判断，这个时候我们可以使用`switch`语句来实现：

```go
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
```

因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛。

关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。

## 怎样判定一个数据类型的某一个方法实现的就是某个接口类型中的某个方法呢？

这有两个充分必要条件，一个是“**两个方法的签名需要完全一致**”，另一个是“**两个方法的名称要一模一样**”。显然，这比判断一个函数是否实现了某个函数类型要更加严格一些。

```go
package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	// 示例2。
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}

```

我声明的类型Dog附带了 3 个方法。其中有 2 个值方法，分别是Name和Category，另外还有一个指针方法SetName。这就意味着，**Dog类型本身的方法集合中只包含了 2 个方法，也就是所有的值方法。而它的指针类型*Dog方法集合却包含了 3 个方法，**

也就是说，它拥有Dog类型附带的所有值方法和指针方法。又由于这 3 个方法恰恰分别是Pet接口中某个方法的实现，所以*Dog类型就成为了Pet接口的实现类型。

正因为如此，我可以声明并初始化一个Dog类型的变量dog，然后把它的指针值赋给类型为Pet的变量pet。

这里有几个名词需要你先记住。对于一个接口类型的变量来说，例如上面的变量pet，**我们赋给它的值可以被叫做它的实际值（也称动态值），而该值的类型可以被叫做这个变量的实际类型（也称动态类型）。**

动态类型这个叫法是相对于**静态类型**而言的。对于变量pet来讲，它的**静态类型**就是Pet，并且永远是Pet，但是它的动态类型却会随着我们赋给它的动态值而变化。

在我们给一个接口类型的变量赋予实际的值之前，它的动态类型是不存在的。



## 当我们为一个接口变量赋值时会发生什么？

为了突出问题，我把Pet接口的声明简化了一下。

```go
type Pet interface {
  Name() string
  Category() string
}
```

我从中去掉了Pet接口的那个名为SetName的方法。这样一来，Dog类型也就变成Pet接口的实现类型了。

现在，我先声明并初始化了一个Dog类型的变量dog，这时它的name字段的值是"little pig"。然后，我把该变量赋给了一个Pet类型的变量pet。最后我通过调用dog的方法SetName把它的name字段的值改成了"monster"。

```go
dog := Dog{"little pig"}
var pet Pet = dog
dog.SetName("monster")
```

所以，我要问的具体问题是：在以上代码执行后，pet变量的字段name的值会是什么？

**这个题目的典型回答是：pet变量的字段name的值依然是"little pig"。**

首先，由于dog的SetName方法是指针方法，所以该方法持有的接收者就是指向dog的指针值的副本，因而其中对接收者的name字段的设置就是对变量dog的改动。那么当dog.SetName("monster")执行之后，dog的name字段的值就一定是"monster"。如果你理解到了这一层，那么请小心前方的陷阱

为什么dog的name字段值变了，**而pet的却没有呢？这里有一条通用的规则需要你知晓：如果我们使用一个变量给另外一个变量赋值，那么真正赋给后者的，并不是前者持有的那个值，而是该值的一个副本。**

例如，我声明并初始化了一个Dog类型的变量dog1，这时它的name是"little pig"。然后，我在把dog1赋给变量dog2之后，修改了dog1的name字段的值。这时，dog2的name字段的值是什么？

```go
dog1 := Dog{"little pig"}
dog2 := dog1
dog1.name = "monster"
```

这个问题与前面那道题几乎一样，只不过这里没有涉及接口类型。这时的dog2的name仍然会是"little pig"。这就是我刚刚告诉你的那条通用规则的又一个体现。

我在前面说过，接口类型本身是无法被值化的。**在我们赋予它实际的值之前，它的值一定会是nil，这也是它的零值。**

反过来讲，一旦它被赋予了某个实现类型的值，它的值就不再是nil了。不过要注意，即使我们像前面那样把dog的值赋给了pet，pet的值与dog的值也是不同的。**这不仅仅是副本与原值的那种不同。**

当我们给一个接口变量赋值的时候，**该变量的动态类型会与它的动态值一起被存储在一个专用的数据结构中。**

严格来讲，这样一个变量的值其实是这个专用数据结构的一个实例，而不是我们赋给该变量的那个实际的值。所以我才说，pet的值与dog的值肯定是不同的，无论是从它们存储的内容，还是存储的结构上来看都是如此。不过，我们可以认为，这时pet的值中包含了dog值的副本。

我们就把这个专用的数据结构叫做iface吧，在 Go 语言的runtime包中它其实就叫这个名字。

iface的实例会包含**两个指针**，一个是**指向类型信息的指针**，另一个是**指向动态值的指针**。这里的类型信息是由另一个专用数据结构的实例承载的，其中包含了动态值的类型，以及使它实现了接口的方法和调用它们的途径，等等。总之，**接口变量被赋予动态值的时候，存储的是包含了这个动态值的副本的一个结构更加复杂的值。**



## 接口变量的值在什么情况下才真正为nil？

这个问题初看起来就不是个问题。对于一个引用类型的变量，它的值是否为nil完全取决于我们赋给它了什么，是这样吗？我们先来看一段代码：

```go
var dog1 *Dog
fmt.Println("The first dog is nil. [wrap1]")
dog2 := dog1
fmt.Println("The second dog is nil. [wrap1]")
var pet Pet = dog2
if pet == nil {
  fmt.Println("The pet is nil. [wrap1]")
} else {
  fmt.Println("The pet is not nil. [wrap1]")
}
```

我先声明了一个*Dog类型的变量dog1，并且没有对它进行初始化。这时该变量的值是什么？显然是nil。然后我把该变量赋给了dog2，后者的值此时也必定是nil，对吗？

现在问题来了：当我把dog2赋给Pet类型的变量pet之后，变量pet的值会是什么？答案是nil吗？



当我们把dog2的值赋给变量pet的时候，dog2的值会先被复制，不过由于在这里它的值是nil，所以就没必要复制了。

然后，Go 语言会用我上面提到的那个专用数据结构iface的实例包装这个dog2的值的副本，这里是nil。

虽然被包装的动态值是nil，但是pet的值却不会是nil，因为这个动态值只是pet值的一部分而已。

顺便说一句，这时的pet的动态类型就存在了，是*Dog。我们可以通过fmt.Printf函数和占位符%T来验证这一点，另外reflect包的TypeOf函数也可以起到类似的作用。

换个角度来看。我们把nil赋给了pet，但是pet的值却不是nil。

这很奇怪对吗？其实不然。在 Go 语言中，我们把由字面量nil表示的值叫做无类型的nil。这是真正的nil，因为它的类型也是nil的。虽然dog2的值是真正的nil，但是当我们把这个变量赋给pet的时候，Go 语言会把它的类型和值放在一起考虑。



也就是说，这时 Go 语言会识别出赋予pet的值是一个*Dog类型的nil。然后，**Go 语言就会用一个iface的实例包装它，包装后的产物肯定就不是nil了。**

只要我们把一个有类型的nil赋给接口变量，那么这个变量的值就一定不会是那个真正的nil。因此，当我们使用判等符号==判断pet是否与字面量nil相等的时候，答案一定会是false。

那么，怎样才能让一个接口变量的值真正为nil呢？**要么只声明它但不做初始化，要么直接把字面量nil赋给它**。

## 怎样实现接口之间的组合？

接口类型间的嵌入要更简单一些，因为它不会涉及方法间的“屏蔽”。只要组合的接口之间有**同名的方法**就会产生冲突，从而无法通过编译，即使同名方法的签名彼此不同也会是如此。因此，接口的组合根本不可能导致“屏蔽”现象的出现。

与结构体类型间的嵌入很相似，我们只要把一个接口类型的名称直接写到另一个接口类型的成员列表中就可以了。比如：

```go
type Animal interface {
  ScientificName() string
  Category() string
}

type Pet interface {
  Animal
  Name() string
}
```

接口类型Pet包含了两个成员，一个是代表了另一个接口类型的Animal，一个是方法Name的定义。它们都被包含在Pet的类型声明的花括号中，并且都各自独占一行。此时，Animal接口包含的所有方法也就成为了Pet接口的方法。

**Go 语言团队鼓励我们声明体量较小的接口，并建议我们通过这种接口间的组合来扩展程序、增加程序的灵活性。**

这是因为相比于包含很多方法的大接口而言，**小接口可以更加专注地表达某一种能力或某一类特征，同时也更容易被组合在一起。**



## 如果我们把一个值为nil的某个实现类型的变量赋给了接口变量，那么在这个接口变量上仍然可以调用该接口的方法吗？如果可以，有哪些注意事项？如果不可以，原因是什么？

可以调用。但是请注意，这个被调用的方法在此时所持有的接收者的值是nil。因此，**如果该方法引用了其接收者的某个字段，那么就会引发 panic！
