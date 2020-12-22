# 函数
## 函数定义

Go语言中定义函数使用`func`关键字，具体格式如下：

```go
func 函数名(参数)(返回值){
    函数体
}
```

其中：

- 函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
- 参数：参数由参数变量和参数变量的类型组成，多个参数之间使用`,`分隔。
- 返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用`()`包裹，并用`,`分隔。
- 函数体：实现指定功能的代码块。

我们先来定义一个求两个数之和的函数：

```go
func intSum(x int, y int) int {
	return x + y
}
```

## 函数的调用

调用有返回值的函数时，可以不接收其返回值。

## 参数

### 类型简写

函数的参数中如果相邻变量的类型相同，则可以省略类型，例如：

```go
func intSum(x, y int) int {
	return x + y
}
```

### 可变参数

可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加`...`来标识。

注意：可变参数通常要作为函数的最后一个参数。

举个例子：

```go
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
```

**本质上，函数的可变参数是通过切片来实现的**。即可变参数的类型其实就是切片

## 返回值

### 多返回值

Go语言中函数支持多返回值，函数如果有多个返回值时必须用`()`将所有返回值包裹起来。

举个例子：

```go
func calc(x,y int ) (int,int)  {
	if x <y {
		return calc(y,x)
	}
	sum := x+y
	sub := x-y
	return sum,sub
}
```

### 返回值命名

函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过`return`关键字返回。

例如：

```go
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}
```

### 返回值补充

当我们的一个函数返回值类型为slice时，**nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。**

```go
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	...
}
```
## 包级函数

同一个包中的函数哪怕是私有的（函数名称首字母小写）也可以被调用。如果不同包的函数要被调用，那么函数的作用域必须是公有的，也就是**函数名称的首字母要大写**，比如 Println。

在后面的包、作用域和模块化的课程中我会详细讲解，这里可以先记住：

1. 函数名称首字母小写代表私有函数，只有在同一个包中才可以被调用；
2. 函数名称首字母大写代表公有函数，不同的包也可以调用；
3. 任何一个函数都会从属于一个包。

> 小提示：Go 语言没有用 public、private 这样的修饰符来修饰函数是公有还是私有，而是通过函数名称的大小写来代表，这样省略了烦琐的修饰符，更简洁。

## 变量作用域

### 全局变量

全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。
```go
package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}
func main() {
	testGlobalVar() //num=10
}
```
### 局部变量

局部变量又分为两种： 函数内定义的变量无法在该函数外使用。

如果局部变量和全局变量重名，**优先访问局部变量**。

语句块定义的变量，通常我们会在if条件判断、for循环、switch语句上使用这种定义变量的方式。**只在语句块中生效**

## 函数类型与变量

```go
func funcVariable(a, b int) (int, int) {
   c := a + b
   d := a - b
   return c, d
}

var a = funcVariable    //没有括号和参数，有括号和参数那个就是函数调用了
	sum, sub := a(12, 6)
	fmt.Println(sum, sub)
```

### 定义函数类型
我们可以使用`type`关键字来定义一个函数类型，具体格式如下：

```go
type calculation func(int, int) int
```

上面语句定义了一个`calculation`类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。

> **函数类型属于引用类型**，它的值可以为nil，而这种类型的零值恰恰就是nil。

简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。

```go
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
```

add和sub都能赋值给calculation类型的变量。

```go
var c calculation
c = add
```

### 函数类型变量

我们可以声明函数类型的变量并且为该变量赋值：

```go
func main() {
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}
```
## 高阶函数

高阶函数分为函数作为参数和函数作为返回值两部分。
### 函数作为参数

```go
func add(x, y int) int {
   return x + y
}

/*
	函数作为参数
*/
func parameterFunc(x, y int, ob func(int, int) int) int {
	fmt.Println(ob(x, y))
	return ob(x, y)
}

func main{
  parameterFunc(2, 3, add)
}



```

### 函数作为返回值

```go
/*
   函数作为返回值
*/
func operation(s string) (func(int, int) int, error) {
   switch s {
   case "+":
      return add, nil
   default:
      err := errors.New("只能传+号，因为我没有写减法函数")
      return nil,err
   }
}

func main{
  c, v := operation("+")
	fmt.Println(c(4, 5))
	fmt.Println(v)
}

```

## 匿名函数和闭包

### 匿名函数

```go
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
```

### 闭包

**闭包指的是一个函数和与其相关的引用环境组合而成的实体**。简单来说，`闭包=函数+引用环境`

```go
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := adder()     //此时为f1 已经不是f的生命周期了 所以在f1的生命周期x被初始化为0
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}
```

变量`f`是一个函数并且它引用了其外部作用域中的`x`变量，此时`f`就是一个闭包。 在`f`的生命周期内，变量`x`也一直有效

```go
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder2(10)  //给定外部作用域变量x
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40  注意这里是40 不是30 
	fmt.Println(f(30)) //70  因为闭包引用环境 在上次计算过程 x已经改变了 此时x是40

	f1 := adder2(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110
}
```

```go
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}
```

```go
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
```

## defer语句

Go语言中的`defer`语句会将其后面跟随的语句进行**延迟处理**。在`defer`归属的函数**即将返回**时，将延迟处理的语句按`defer`定义的**逆序**进行执行，也就是说，先被`defer`的语句最后被执行，最后被`defer`的语句，最先被执行。

```go
func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

// output
start
end
3
2
1
```

由于`defer`语句延迟调用的特性，所以`defer`语句能非常方便的**处理资源释放问题**。比如：资源清理、文件关闭、解锁及记录时间等。

### defer执行时机

在Go语言的函数中`return`语句在底层**并不是原子操作**，它分为给**返回值赋值和RET指令**两步。而`defer`语句执行的时机就在**返回值赋值操作后，RET指令执行前**。具体如下图所示：
![defer](https://img-blog.csdnimg.cn/20201221205046420.png)

### defer经典案例

```go
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1())   //5
	fmt.Println(f2())   //6
	fmt.Println(f3())   //5
	fmt.Println(f4())   //5
}
```

解析：
![defer](https://img-blog.csdnimg.cn/20201221205222849.png)

### defer面试题

```go
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
```

提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
![defer](https://img-blog.csdnimg.cn/20201221205454364.png)

## panic/recover

Go语言中目前（Go1.12）是没有异常机制，但是使用`panic/recover`模式来处理错误。 `panic`可以在任何地方引发，但`recover`只有在`defer`调用的函数中有效

```go
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}

// output
func A
panic: panic in B

goroutine 1 [running]:
main.funcB(...)
        .../code/func/main.go:12
main.main()
        .../code/func/main.go:20 +0x98
```

程序运行期间`funcB`中引发了`panic`导致程序崩溃，异常退出了。这个时候我们就可以通过`recover`将程序恢复回来，继续往后执行

```go
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
//打印结果
func A
recover in B
func C
```

**注意：**

1. `recover()`必须搭配`defer`使用。
2. `defer`一定要在可能引发`panic`的语句之前定义。

### 从 panic 被引发到程序终止运行的大致过程是什么？
在 Go 语言中，因 panic 导致程序结束运行的退出状态码一般都会是2。
大致的过程：某个函数中的某行代码有意或无意地引发了一个 panic。这时，初始的 panic 详情会被建立起来，并且该程序的**控制权会立即从此行代码转移至调用其所属函数的那行代码上，也就是调用栈中的上一级。**

这也意味着，**此行代码所属函数的执行随即终止**。紧接着，控制权并不会在此有片刻的停留，它又会**立即转移**至再上一级的调用代码处。控制权如此一级一级地沿着调用栈的反方向**传播至顶端**，也就是我们编写的最外层函数那里。

这里的最外层函数指的是go函数，对于主 goroutine 来说就是main函数。但是控制权也不会停留在那里，而是被 Go 语言运行时**系统收回**。

随后，程序崩溃并终止运行，承载程序这次运行的进程也会随之死亡并消失。与此同时，在这个控制权传播的过程中，panic 详情会被逐渐地积累和完善，并会在程序终止之前被打印出来。





panic 可能是我们在无意间（或者说一不小心）引发的，如前文所述的索引越界。这类 panic 是真正的、在我们意料之外的程序异常。

不过，除此之外，我们还是可以有意地引发 panic。Go 语言的内建函数panic是专门用于引发 panic 的。panic函数使程序开发者可以在程序运行期间报告异常。

注意，这与从函数返回错误值的意义是完全不同的。当我们的函数返回一个非nil的错误值时，函数的调用方有权选择不处理，并且不处理的后果往往是不致命的。

这里的“不致命”的意思是，不至于使程序无法提供任何功能（也可以说僵死）或者直接崩溃并终止运行（也就是真死）。

但是，当一个 panic 发生时，如果我们不施加任何保护措施，那么导致的直接后果就是**程序崩溃**，就像前面描述的那样，这显然是**致命的**。

我在这里再提示一点。**panic 详情会在控制权传播的过程中，被逐渐地积累和完善，并且，控制权会一级一级地沿着调用栈的反方向传播至顶端。**

因此，在针对某个 goroutine 的代码执行信息中，调用栈底端的信息会先出现，然后是上一级调用的信息，以此类推，最后才是此调用栈顶端的信息。

比如，main函数调用了caller1函数，而caller1函数又调用了caller2函数，那么caller2函数中代码的执行信息会先出现，然后是caller1函数中代码的执行信息，最后才是main函数的信息。

```shell
goroutine 1 [running]:
main.caller2()
 /Users/haolin/GeekTime/Golang_Puzzlers/src/puzzlers/article19/q1/demo48.go:22 +0x91
main.caller1()
 /Users/haolin/GeekTime/Golang_Puzzlers/src/puzzlers/article19/q1/demo48.go:15 +0x66
main.main()
 /Users/haolin/GeekTime/Golang_Puzzlers/src/puzzlers/article19/q1/demo48.go:9 +0x66
exit status 2
```



![从 panic 到程序崩溃](https://img-blog.csdnimg.cn/20201222211353652.png)



### 一个函数怎样才能把 panic 转化为error类型值，并将其作为函数的结果值返回给调用方？

```go
func doSomething() (err error) {
  defer func() {
    p := recover()
    err = fmt.Errorf("FATAL ERROR: %s", p)
  }()
  panic("Oops!!")
}
```

注意结果声明的写法。这是一个带有名称的结果声明。

### 怎样让 panic 包含一个值，以及应该让它包含什么样的值？

在调用panic函数时，把某个值作为参数传给该函数就可以了。由于panic函数的唯一一个参数是空接口（也就是interface{}）类型的，所以从语法上讲，它可以接受任何类型的值。

但是，我们最好传入error类型的错误值，或者其他的可以被有效序列化的值。这里的“有效序列化”指的是，可以更易读地去表示形式转换。

对于fmt包下的各种打印函数来说，error类型值的Error方法与其他类型值的String方法是等价的，它们的唯一结果都是string类型的。

**我们在通过占位符%s打印这些值的时候，它们的字符串表示形式分别都是这两种方法产出的。**

一旦程序异常了，我们就一定要把异常的相关信息记录下来，这通常都是记到程序日志里。

我们在为程序排查错误的时候，首先要做的就是查看和解读程序日志；而最常用也是最方便的日志记录方式，就是记下相关值的字符串表示形式。

所以，**如果你觉得某个值有可能会被记到日志里，那么就应该为它关联String方法。如果这个值是error类型的，那么让它的Error方法返回你为它定制的字符串表示形式就可以了**。

对于此，你可能会想到fmt.Sprintf，以及fmt.Fprintf这类可以格式化并输出参数的函数。

是的，它们本身就可以被用来输出值的某种表示形式。不过，它们在功能上，肯定远不如我们自己定义的Error方法或者String方法。因此**，为不同的数据类型分别编写这两种方法总是首选。**

可是，这与传给panic函数的参数值又有什么关系呢？其实道理是相同的。至少在程序崩溃的时候，panic 包含的那个值字符串表示形式会被打印出来。

另外，我们还可以施加某种保护措施，避免程序的崩溃。这个时候，panic 包含的值会被取出，而在取出之后，它一般都会被打印出来或者记录到日志里。



### 怎样施加应对 panic 的保护措施，从而避免程序崩溃？

**Go 语言的内建函数recover专用于恢复 panic，或者说平息运行时恐慌。recover函数无需任何参数，并且会返回一个空接口类型的值。**

如果用法正确，这个值实际上就是即将恢复的 panic 包含的值。并且，如果这个 panic 是因我们调用panic函数而引发的，那么该值同时也会是我们此次调用panic函数时，**传入的参数值副本。**请注意，这里强调用法的正确。我们先来看看什么是不正确的用法。

```go
package main

import (
 "fmt"
 "errors"
)

func main() {
 fmt.Println("Enter function main.")
 // 引发panic。
 panic(errors.New("something wrong"))
 p := recover()
 fmt.Printf("panic: %s\n", p)
 fmt.Println("Exit function main.")
}
```

在上面这个main函数中，我先通过调用panic函数引发了一个 panic，紧接着想通过调用recover函数恢复这个 panic。可结果呢？你一试便知，程序依然会崩溃，这个recover函数调用并不会起到任何作用，甚至都没有机会执行。

还记得吗？**我提到过 panic 一旦发生，控制权就会讯速地沿着调用栈的反方向传播。所以，在panic函数调用之后的代码，根本就没有执行的机会。**

那如果我把调用recover函数的代码提前呢？也就是说，先调用recover函数，再调用panic函数会怎么样呢？

这显然也是不行的，因为，如果在我们调用recover函数时未发生 panic，那么该函数就不会做任何事情，并且只会返回一个nil。

换句话说，这样做毫无意义。那么，到底什么才是正确的recover函数用法呢？这就不得不提到defer语句了。

顾名思义，**defer语句就是被用来延迟执行代码的。延迟到什么时候呢？这要延迟到该语句所在的函数即将执行结束的那一刻，无论结束执行的原因是什么。**

这与go语句有些类似，**一个defer语句总是由一个defer关键字和一个调用表达式组成。**

这里存在一些限制，有一些调用表达式是不能出现在这里的，包括：**针对 Go 语言内建函数的调用表达式，以及针对unsafe包中的函数的调用表达式。**

顺便说一下，**对于go语句中的调用表达式，限制也是一样的**。另外，在这里被调用的函数可以是有名称的，也可以是匿名的。我们可以把这里的函数叫做defer函数或者延迟函数。注意，**被延迟执行的是defer函数，而不是defer语句**。

```go
package main

import (
 "fmt"
 "errors"
)

func main() {
 fmt.Println("Enter function main.")
 defer func(){
  fmt.Println("Enter defer function.")
  if p := recover(); p != nil {
   fmt.Printf("panic: %s\n", p)
  }
  fmt.Println("Exit defer function.")
 }()
 // 引发panic。
 panic(errors.New("something wrong"))
 fmt.Println("Exit function main.")
}
```

在这个main函数中，我先编写了一条defer语句，并在defer函数中调用了recover函数。仅当调用的结果值不为nil时，也就是说只有 panic 确实已发生时，我才会打印一行以“panic:”为前缀的内容。

紧接着，我调用了panic函数，并传入了一个error类型值。这里一定要注意，**我们要尽量把defer语句写在函数体的开始处，因为在引发 panic 的语句之后的所有语句，都不会有任何执行机会。**

也只有这样，defer函数中的recover函数调用才会拦截，并恢复defer语句所属的函数，及其调用的代码中发生的所有 panic。

### 如果一个函数中有多条defer语句，那么那几个defer函数调用的执行顺序是怎样的？

如果只用一句话回答的话，那就是：**在同一个函数中，defer函数调用的执行顺序与它们分别所属的defer语句的出现顺序（更严谨地说，是执行顺序）完全相反。**

在defer语句每次执行的时候，**Go 语言会把它携带的defer函数及其参数值另行存储到一个链表中**。这个链表与该defer语句所属的函数是对应的，并且，它是先进后出（FILO）的，相当于一个栈。

### 我们可以在defer函数中恢复 panic，那么可以在其中引发 panic 吗？

当然可以。这样做可以把原先的 panic 包装一下再抛出去。

## 函数是一等的公民

函数可是一等的（first-class）公民，函数类型也是一等的数据类型。

这意味着**函数不但可以用于封装代码、分割功能、解耦逻辑，还可以化身为普通的值，在其他函数间传递、赋予变量、做类型判断和转换等等**，就像切片和字典的值那样。

函数值可以由此成为能够被随意传播的独立逻辑组件（或者说功能模块）。

> 函数的签名其实就是函数的**参数列表和结果列表**的统称，它定义了可用来鉴别不同函数的那些特征，同时也定义了我们与函数交互的方式。

注意，**各个参数和结果的名称不能算作函数签名的一部分**，甚至对于结果声明来说，没有名称都可以。



**只要两个函数的参数列表和结果列表中的元素顺序及其类型是一致的，我们就可以说它们是一样的函数，或者说是实现了同一个函数类型的函数。**

严格来说，**函数的名称也不能算作函数签名的一部分**，它只是我们在调用函数时，需要给定的标识符而已。

### 如何实现闭包

在一个函数中存在对外来标识符的引用。所谓的外来标识符，**既不代表当前函数的任何参数或结果，也不是函数内部声明的**，它是直接从外边拿过来的。即自由变量。

我们说的这个函数（以下简称闭包函数）就是因为引用了自由变量，而呈现出了一种“不确定”的状态，也叫“开放”状态。也就是说，它的内部逻辑并不是完整的，有一部分逻辑需要这个自由变量参与完成，而后者到底代表了什么在闭包函数被定义的时候却是未知的。即使对于像 Go 语言这种静态类型的编程语言而言，我们在定义闭包函数的时候最多也只能知道自由变量的类型。
![闭包](https://img-blog.csdnimg.cn/20201221210359797.png)
那么，实现闭包的意义又在哪里呢？**表面上看，我们只是延迟实现了一部分程序逻辑或功能而已，但实际上，我们是在动态地生成那部分程序逻辑。**

### 传入函数的那些参数值后来怎么样了

```go
package main

import "fmt"

func main() {
  array1 := [3]string{"a", "b", "c"}
  fmt.Printf("The array: %v\n", array1)
  array2 := modifyArray(array1)
  fmt.Printf("The modified array: %v\n", array2)
  fmt.Printf("The original array: %v\n", array1)
}

func modifyArray(a [3]string) [3]string {
  a[1] = "x"
  return a
}
```



关键问题是，原数组会因modify函数对参数值的修改而改变吗？

**原数组不会改变。为什么呢？原因是，所有传给函数的参数值都会被复制，函数在其内部使用的并不是参数值的原值，而是它的副本。**



由于数组是值类型，所以每一次复制都会拷贝它，以及它的所有元素值。我在modify函数中修改的只是原数组的副本而已，并不会对原数组造成任何影响。



注意，**对于引用类型，比如：切片、字典、通道，像上面那样复制它们的值，只会拷贝它们本身而已，并不会拷贝它们引用的底层数据。也就是说，这时只是浅表复制，而不是深层复制。**



**以切片值为例，如此复制的时候，只是拷贝了它指向底层数组中某一个元素的指针，以及它的长度值和容量值，而它的底层数组并不会被拷贝，但是会改变。**

```go
func parameterDemo()  {
	array1 := [3]string{"a", "b", "c"}
	a := array1[:]
	fmt.Printf("The slice a: %v\n",a)  //The slice a: [a b c]
	a1 := modifySlice(a)
	fmt.Printf("The slice a1: %v\n",a1) //The slice a1: [gg b c]
	fmt.Printf("The slice a: %v\n",a) // The slice a: [gg b c]
	fmt.Printf("The array1: %v\n", array1) //The array1: [gg b c]
}
```

另外还要注意，就算我们传入函数的是一个值类型的参数值，但如果这个参数值中的某个元素是引用类型的，那么我们仍然要小心。

```go
func main() {
  complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	fmt.Printf("The array: %v\n", complexArray1)  //The array: [[d e f] [g h i] [j k l]]
	complexArray2 := modifySliceArray(complexArray1)
	fmt.Printf("The array: %v\n", complexArray2)  // The array: [[hello golang] [g h i] [j k l]]
	fmt.Printf("The array: %v\n", complexArray1)   //The array: [[d e f] [g h i] [j k l]]
}

func modifySliceArray(a [3][]string) [3][]string {
	a[0] = []string{"hello","golang"}
	return a
}
```

变量complexArray1是[3][]string类型的，也就是说，虽然它是一个数组，但是其中的每个元素又都是一个切片。这样一个值被传入函数的话，函数中对该参数值的修改会影响到complexArray1本身吗？

**如果是进行一层修改，即数组的某个完整元素进行修改（指针变化），那么原有数组不变；如果进行二层修改，即数组中某个元素切片内的某个元素再进行修改(指针未改变)，那么原有数据也会跟着改变，传参可以理解是浅copy，参数本身的指针是不同，但是元素指针相同，对元素指针所指向目的的操作会影响传参过程中的原始数据；**



```go
func modifySliceArray1(a [3][]string) [3][]string {
	a[0][0] = "Hello go"
	return a
}
complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	fmt.Printf("The array: %v\n", complexArray1)  //The array: [[d e f] [g h i] [j k l]]
	complexArray2 := modifySliceArray1(complexArray1)
	fmt.Printf("The array: %v\n", complexArray2)  //The array: [[Hello go e f] [g h i] [j k l]]
	fmt.Printf("The array: %v\n", complexArray1)   //The array: [[Hello go e f] [g h i] [j k l]]
```

### 函数真正拿到的参数值其实只是它们的副本，那么函数返回给调用方的结果值也会被复制吗？

函数返回给调用方的结果值也会被复制。不过，在一般情况下，我们不用太在意。但如果函数在返回结果值之后依然保持执行并会对结果值进行修改，那么我们就需要注意了。





