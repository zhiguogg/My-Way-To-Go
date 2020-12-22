# 控制结构
## if 条件语句

```go
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
```

Go语言规定与`if`匹配的左括号`{`必须与`if和表达式`放在同一行，`{`放在其他位置会触发编译错误。 同理，与`else`匹配的`{`也必须与`else`写在同一行，`else`也必须与上一个`if`或`else if`右边的大括号在同一行。

由于 `if` 和 `switch` 可接受初始化语句， 因此用它们来**设置局部变量**十分常见。
```go
func ifDemo()  {
	if i:=6; i >10 {

		fmt.Println("i>10")

	} else if  i>5 && i<=10 {

		fmt.Println("5<i<=10")

	} else {

		fmt.Println("i<=5")

	}
}

```
**通过 if 简单语句声明的变量，只能在整个 if……else if……else 条件语句中使用**
## switch 选择语句

Go 的 switch 比 C 的更通用。其表达式无需为常量或整数，case 语句会自上而下逐一进行求值直到匹配为止。若 switch 后面没有表达式，它将匹配 true，因此，我们可以将 if-else-if-else 链写成一个 switch，这也更符合 Go 的风格。

```GO
switch i:=6;{

case i>10:

    fmt.Println("i>10")

case i>5 && i<=10:

    fmt.Println("5<i<=10")

default:

    fmt.Println("i<=5")

}

```

Go语言规定每个`switch`只能有一个`default`分支。

一个分支可以有多个值，多个case值中间使用英文逗号分隔。

```GO
func testSwitch3() {
	switch n := 7; n {  //和if之前加一个执行语句多像 不能有var
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}
```

分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。

```GO
func switchDemo4() {
	age := 30
	switch {   // 不能有表达式
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}
```

在 Go 语言中，switch 的 case 从上到下逐一进行判断，一旦满足条件，立即执行对应的分支并返回，其余分支不再做判断。也就是说 Go 语言的 switch 在默认情况下，case 最后自带 break。这和其他编程语言不一样，比如 C 语言在 case 分支里必须要有明确的 break 才能退出一个 case。Go 语言的这种设计就是为了防止忘记写 break 时，下一个 case 被执行。

`fallthrough`语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。

```go
func switchDemo5() {
	s := "a"
	switch {  
	case s == "a":
		fmt.Println("a")
		fallthrough // 自动执行下一个分支
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

//  output
a
b
```



`switch` 也可用于判断接口变量的动态类型。如 **类型选择** 通过圆括号中的关键字 `type` 使用类型断言语法。若 `switch` 在表达式中声明了一个变量，那么该变量的每个子句中都将有该变量对应的类型。

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T 打印任何类型的 t
case bool:
    fmt.Printf("boolean %t\n", t)             // t 是 bool 类型
case int:
    fmt.Printf("integer %d\n", t)             // t 是 int 类型
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t 是 *bool 类型
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t 是 *int 类型
}
```

## for 循环语句

Go 语言中的所有循环类型均可以使用`for`关键字来完成。

```go
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```

for循环的初始语句可以被忽略，但是初始语句后的分号必须要写

```go
func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}
```

for循环的初始语句和结束语句都可以省略

```go
func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++ // 结束语句在这
	}
}
```

这种写法类似于其他编程语言中的`while`，在`while`后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。

无限循环

```go
for {
    循环体语句
}
```

for循环可以通过`break`、`goto`、`return`、`panic`语句强制退出循环。
## for range

Go语言中可以使用`for range`遍历数组、切片、字符串、map 及通道（channel）。 通过`for range`遍历的返回值有以下规律：

1. 数组、切片、字符串返回索引和值。
2. map返回键和值。
3. 通道（channel）只返回通道内的值。

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

若你只需要该遍历中的第一个项（键或下标），去掉第二个就行了：

```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```

若你只需要该遍历中的第二个项（值），请使用**空白标识符**，即下划线来丢弃第一个值：

```go
sum := 0
for _, value := range array {
    sum += value
}
```

## 使用携带range子句的for语句时需要注意哪些细节
for语句：

```go
numbers1 := [...]int{1, 2, 3, 4, 5, 6}    //这里循环的是数组
	for i := 10; i < len(numbers1); i++ {
		if i == 3 {
			numbers1[i] += 10
		}
	}
	fmt.Println(numbers1)   //[1 2 3 4 5 6]
```

```go
numbers1 := []int{1, 2, 3, 4, 5, 6}   //这里循环的是切片
	for i := 10; i < len(numbers1); i++ {
		if i == 3 {
			numbers1[i] += 10
		}
	}
	fmt.Println(numbers1)  //[1 2 3 4 5 6]
```

可见不会改变原有的值！



range表达式：

```go
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] += 10
		}
	}
	fmt.Println(numbers1)

// 打印结果
[1 2 3 14 5 6]
```

改变一下：

```go
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}   // 注意 这是数组
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)

// 打印结果
[7 3 5 7 9 11]

//分析过程
	// i = 0 e = 1  else : n1 = 2+1 = 3
	// i = 1 e = 2  else : n2 = 3+2 = 5
	// i = 2 e = 3  else : n3 = 4+3 = 7
	// i = 3 e = 4  else : n4 = 5+4 = 9
	// i = 4 e = 5  else : n5 = 6+5 = 11
	// i = 5 4 = 6  else : n0 = 1+6 = 7
```

由此可以推论出：

- range表达式**只会在for语句开始执行时被求值一次**，无论后边会有多少次迭代；
- range表达式的求值结果会被复制，也就是说，被迭代的对象是range表达式结果值的**副本而不是原值**。



现在将数组修改为切片：

```go
numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers3) - 1
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)

//打印结果
[22 3 6 10 15 21]
```

**切片与数组是不同的，前者是引用类型的，而后者是值类型的。**

因为引用类型的副本其实还是地址，对地址进行操作会改变底层数组值，故和数组的不一致。即一边迭代一边改变了number3的值。

## switch语句中的switch表达式和case表达式之间有着怎样的联系

```go
value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
switch 1 + 3 {
case value1[0], value1[1]:
  fmt.Println("0 or 1")
case value1[2], value1[3]:
  fmt.Println("2 or 3")
case value1[4], value1[5], value1[6]:
  fmt.Println("4 or 5 or 6")
}
```

**无法编译！**

switch语句对switch表达式的结果类型，以及各个case表达式中子表达式的结果类型都是有要求的。毕竟，在 Go 语言中，**只有类型相同的值之间才有可能被允许进行判等操作**。

如果switch表达式的结果值是无类型的常量，比如1 + 3的求值结果就是无类型的常量4，那么这个常量会被自动地转换为此种常量的默认类型的值，比如整数4的默认类型是int，又比如浮点数3.14的默认类型是float64。

因此，由于上述代码中的switch表达式的结果类型是int，而那些case表达式中子表达式的结果类型却是int8，它们的类型并不相同，所以这条switch语句是无法通过编译的。



改变一下

```GO
value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
switch value2[4] {
case 0, 1:
  fmt.Println("0 or 1")
case 2, 3:
  fmt.Println("2 or 3")
case 4, 5, 6:
  fmt.Println("4 or 5 or 6")
}
```

**可以编译！**

**如果case表达式中子表达式的结果值是无类型的常量，那么它的类型会被自动地转换为switch表达式的结果类型**，又由于上述那几个整数都可以被转换为int8类型的值，所以对这些表达式的结果值进行判等操作是没有问题的。

![switch](https://img-blog.csdnimg.cn/20201222203825273.png)

## switch语句对它的case表达式有哪些约束

**switch语句在case子句的选择上是具有唯一性的。**

```GO
value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
switch value3[4] {
case 0, 1, 2:
  fmt.Println("0 or 1 or 2")
case 2, 3, 4:
  fmt.Println("2 or 3 or 4")
case 4, 5, 6:
  fmt.Println("4 or 5 or 6")
}
```

**无法编译！**

**不过，好在这个约束本身还有个约束，那就是只针对结果值为常量的子表达式。**

```GO
value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
switch value5[4] {
case value5[0], value5[1], value5[2]:
  fmt.Println("0 or 1 or 2")
case value5[2], value5[3], value5[4]:
  fmt.Println("2 or 3 or 4")
case value5[4], value5[5], value5[6]:
  fmt.Println("4 or 5 or 6")
}
```

**可以编译！**

变量名换成了value5，但这不是重点。重点是，**我把case表达式中的常量都换成了诸如value5[0]这样的索引表达式。**



不过，这种绕过方式对用于**类型判断的switch语句**（以下简称为类型switch语句）就无效了。因为类型switch语句中的case表达式的子表达式，都**必须直接由类型字面量**表示，而无法通过间接的方式表示

```GO
value6 := interface{}(byte(127))
switch t := value6.(type) {
case uint8, uint16:
  fmt.Println("uint8 or uint16")
case byte: //byte类型是uint8类型的别名类型
  fmt.Printf("byte")
default:
  fmt.Printf("unsupported type: %T", t)
}
```

**无法编译！**

byte类型是uint8类型的别名类型.因此，它们两个本质上是**同一个类型**，只是类型名称不同罢了。在这种情况下，这个类型switch语句是**无法通过编译**的，因为子表达式byte和uint8重复了。

普通case子句的编写顺序很重要，**最上边的case子句中的子表达式总是会被最先求值**，在判等的时候顺序也是这样。因此，如果某些子表达式的结果值有重复并且它们与switch表达式的结果值相等，那么位置靠上的case子句总会被选中。

## 在类型switch语句中，我们怎样对被判断类型的那个值做相应的类型转换？

其实这个事情可以让 Go 语言自己来做，例如：

```go
switch t := x.(type) {
// cases
}
```

当流程进入到某个case子句的时候，变量t的值就已经被自动地转换为相应类型的值了。

## 在if语句中，初始化子句声明的变量的作用域是什么？

如果这个变量是新的变量，那么它的作用域就是当前if语句所代表的代码块。注意，后续的else if子句和else子句也包含在当前的if语句代表的代码块之内。






