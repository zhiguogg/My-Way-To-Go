# fmt

fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分。

## 向外输出

### Print

`Print`系列函数会将内容输出到系统的标准输出，区别在于`Print`函数直接输出内容，`Printf`函数支持格式化输出字符串，`Println`函数会在输出内容的结尾添加一个换行符。

```go
func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}
```

举个简单的例子：

```go
func printTest()  {
	fmt.Print("大理","段氏","段誉")
	fmt.Println("大理","段氏","段誉")
	name := "北冥神功"
	fmt.Printf("段誉的武功： %s\n",name)
}

//output
大理段氏段誉大理 段氏 段誉
段誉的武功： 北冥神功
```



### Fprint

`Print`系列函数会将内容输出到**系统的标准输出**，区别在于`Print`函数直接输出内容，`Printf`函数支持格式化输出字符串，`Println`函数会在输出内容的结尾添加一个换行符。

```go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrint(a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintln(a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```

举例：

```go
func FprintTest()  {
	//向标准输出写入内容
	fmt.Fprintln(os.Stdout,"桃花岛黄老邪")

	fileObj,err := os.OpenFile("./fmt.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "黄蓉"
	//向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj,"黄老邪女儿是：%s",name)

}

//output 控制台会打印一句话，同时在当前目录生成一个fmt.txt文件，里面有黄老邪女儿是：黄蓉
桃花岛黄老邪
```

注意，只要满足`io.Writer`接口的类型都支持写入。



### Sprint

`Sprint`系列函数会把传入的数据生成并返回一个字符串。

```go
func Sprintln(a ...interface{}) string {
	p := newPrinter()
	p.doPrintln(a)
	s := string(p.buf)
	p.free()
	return s
}

func Sprintf(format string, a ...interface{}) string {
	p := newPrinter()
	p.doPrintf(format, a)
	s := string(p.buf)
	p.free()
	return s
}

func Sprint(a ...interface{}) string {
	p := newPrinter()
	p.doPrint(a)
	s := string(p.buf)
	p.free()
	return s
}
```

例子：

```go
func SprintTest()  {
	story1 := fmt.Sprintln("郭靖最终娶了黄蓉，没有华筝什么事")
	name1 := "降龙十八掌"
	name2 := "打狗棒法"
	story2 := fmt.Sprintf("郭靖会%s,黄蓉会%s\n",name1,name2)
	fmt.Println(story1,story2)
}
```



### Errorf

`Errorf`函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。

```go
func Errorf(format string, a ...interface{}) error {
	p := newPrinter()
	p.wrapErrs = true
	p.doPrintf(format, a)
	s := string(p.buf)
	var err error
	if p.wrappedErr == nil {
		err = errors.New(s)
	} else {
		err = &wrapError{s, p.wrappedErr}
	}
	p.free()
	return err
}
```

通常使用这种方式来自定义错误类型，例如：

```go
err := fmt.Errorf("这是一个错误")
```

Go1.13版本为`fmt.Errorf`函数新加了一个`%w`占位符用来生成一个可以包裹Error的Wrapping Error。

```go
e := errors.New("原始错误e")
w := fmt.Errorf("Wrap了一个错误%w", e)
```



## 格式化占位符

`*printf`系列函数都支持format格式化参数，在这里我们按照占位符将被替换的变量类型划分，方便查询和记忆。

### 通用占位符

| 占位符 | 说明                               |
| ------ | ---------------------------------- |
| %v     | 值的默认格式表示                   |
| %+v    | 类似%v，但输出结构体时会添加字段名 |
| %#v    | 值的Go语法表示                     |
| %T     | 打印值的类型                       |
| %%     | 百分号                             |

示例代码如下：

```go
type person struct {
	name string
	age int32
}

func CommonSign()  {
	p := person{
		"张无忌",
		22,
	}
	fmt.Printf("倚天屠龙记主角是：%v\n",p)
	fmt.Printf("倚天屠龙记主角是：%+v\n",p)
	fmt.Printf("倚天屠龙记主角是：%#v\n",p)
	fmt.Printf("倚天屠龙记主角是：%T\n",p)
	fmt.Printf("100%%\n")
}
```

output

```shell
倚天屠龙记主角是：{张无忌 22}
倚天屠龙记主角是：{name:张无忌 age:22}
倚天屠龙记主角是：main.person{name:"张无忌", age:22}
倚天屠龙记主角是：main.person
100%
```

### 布尔型

| 占位符 | 说明        |
| ------ | ----------- |
| %t     | true或false |

```go
b := true
	fmt.Printf("真%t\n",b)
```



### 整型

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| %b     | 表示为二进制                                                 |
| %c     | 该值对应的unicode码值                                        |
| %d     | 表示为十进制                                                 |
| %o     | 表示为八进制                                                 |
| %x     | 表示为十六进制，使用a-f                                      |
| %X     | 表示为十六进制，使用A-F                                      |
| %U     | 表示为Unicode格式：U+1234，等价于”U+%04X”                    |
| %q     | 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示 |

示例代码如下：

```go
func IntegerSign() {
	n := 1165   // 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%U\n", n)
	fmt.Printf("%q\n", n)
}
```

output

```shell
10010001101
ҍ   //A
1165
2215
48d
48D
U+048D
'ҍ' //'A'
```



### 浮点数与复数

| 占位符 | 说明                                                   |
| ------ | ------------------------------------------------------ |
| %b     | 无小数部分、二进制指数的科学计数法，如-123456p-78      |
| %e     | 科学计数法，如-1234.456e+78                            |
| %E     | 科学计数法，如-1234.456E+78                            |
| %f     | 有小数部分但无指数部分，如123.456                      |
| %F     | 等价于%f                                               |
| %g     | 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出） |
| %G     | 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出） |

示例代码如下：

```go
f := 12.34
fmt.Printf("%b\n", f)
fmt.Printf("%e\n", f)
fmt.Printf("%E\n", f)
fmt.Printf("%f\n", f)
fmt.Printf("%g\n", f)
fmt.Printf("%G\n", f)
```

output

```shell
6946802425218990p-49
1.234000e+01
1.234000E+01
12.340000
12.34
12.34
```



### 字符串和[]byte

| 占位符 | 说明                                                         |
| ------ | :----------------------------------------------------------- |
| %s     | 直接输出字符串或者[]byte                                     |
| %q     | 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示 |
| %x     | 每个字节用两字符十六进制数表示（使用a-f）                    |
| %X     | 每个字节用两字符十六进制数表示（使用A-F）                    |

示例代码如下：

```go
s := "雪山飞狐"
fmt.Printf("%s\n", s)
fmt.Printf("%q\n", s)
fmt.Printf("%x\n", s)
fmt.Printf("%X\n", s)
```

output

```shell
雪山飞狐
"雪山飞狐"
e99baae5b1b1e9a39ee78b90
E99BAAE5B1B1E9A39EE78B90
```



### 指针

| 占位符 | 说明                           |
| ------ | ------------------------------ |
| %p     | 表示为十六进制，并加上前导的0x |

```go
p := person{
		"张无忌",
		22,
	}
fmt.Printf("%p\n", &p)
fmt.Printf("%#p\n", &p)

//
0xc00011c040
c00011c040
```



### 宽度标识符

宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下：

| 占位符 | 说明               |
| ------ | ------------------ |
| %f     | 默认宽度，默认精度 |
| %9f    | 宽度9，默认精度    |
| %.2f   | 默认宽度，精度2    |
| %9.2f  | 宽度9，精度2       |
| %9.f   | 宽度9，精度0       |

示例代码如下：

```go
func widthSign()  {
	width := 47.21
	fmt.Printf("width: %f\n",width)
	fmt.Printf("width: %11f\n",width)
	fmt.Printf("width: %.2f\n",width)
	fmt.Printf("width: %11.1f\n",width)
	fmt.Printf("width: %11.f\n",width)
}
```

output

```shell
width: 47.210000
width:   47.210000
width: 47.21
width:        47.2
width:          47
```



### 其他falg

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| ’+’    | 总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）； |
| ’ ‘    | 对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格 |
| ’-’    | 在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）； |
| ’#’    | 八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值； |
| ‘0’    | 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面； |

举个例子：

```go
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
```

output

```shell
王语嫣
  王语嫣
王语嫣  
  王语嫣
王语嫣  
   王语
00王语嫣
```



## 获取输入
Go语言`fmt`包下有`fmt.Scan`、`fmt.Scanf`、`fmt.Scanln`三个函数，可以在程序运行过程中从标准输入获取用户的输入。

### Scan

```go
func Scan(a ...interface{}) (n int, err error) {
	return Fscan(os.Stdin, a...)
}
```

- Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
- 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。

```go
func ScanTest()  {
	var (
		name string
		age int
		gender bool
	)
	fmt.Scan(&name,&age,&gender)
	fmt.Printf("name is: %s,age is: %d,gender is %t\n",name,age,gender)
}
```

将上面的代码编译后在终端执行，在终端依次输入郭靖 23 true 使用空格分隔

output

```go
郭靖 23 true
//上面一行是输入的
name is: 郭靖,age is: 23,gender is true
```

`fmt.Scan`从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。



### Scanf

```go
func Scanf(format string, a ...interface{}) (n int, err error) {
	return Fscanf(os.Stdin, format, a...)
}
```

- Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

```go
func ScanfTest()  {
	var (
		name string
		age int
		gender bool
	)
	fmt.Scanf("姓名：%s 年龄：%d 性别：%t",&name,&age,&gender)
	fmt.Printf("name is: %s,age is: %d,gender is %t\n",name,age,gender)
}
```

将上面的代码编译后在终端执行，在终端按照指定的格式依次输入

```go
姓名：郭靖 年龄：11 性别：false  //这行为输入的内容
name is: 郭靖,age is: 11,gender is false
```



`fmt.Scanf`不同于`fmt.Scan`简单的以空格作为输入数据的分隔符，`fmt.Scanf`为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。



### Scanln

```go
func Sscanln(str string, a ...interface{}) (n int, err error) {
	return Fscanln((*stringReader)(&str), a...)
}
```

- Scanln类似Scan，它在**遇到换行时才停止扫描**。最后一个数据后面必须有换行或者到达结束位置。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

`fmt.Scanln`遇到回车就结束扫描了，这个比较常用。



### bufio.NewReader

有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用`bufio`包来实现。示例代码如下：

```go
func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
```



### Fscan系列

这几个函数功能分别类似于`fmt.Scan`、`fmt.Scanf`、`fmt.Scanln`三个函数，只不过它们不是从标准输入中读取数据而是从`io.Reader`中读取数据。



### Sscan系列

这几个函数功能分别类似于`fmt.Scan`、`fmt.Scanf`、`fmt.Scanln`三个函数，只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。


