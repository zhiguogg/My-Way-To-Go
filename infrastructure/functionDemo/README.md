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





