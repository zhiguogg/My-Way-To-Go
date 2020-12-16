# 基本数据类型
## 整型
在 Go 语言中，整型分为：

有符号整型：如 int、int8、int16、int32 和 int64。

无符号整型：如 uint、uint8、uint16、uint32 和 uint64。

uint  32位操作系统上就是`uint32`，64位操作系统上就是`uint64`

int    32位操作系统上就是`int32`，64位操作系统上就是`int64`

uintptr  **无符号整型，用于存放一个指针**

在整型中，如果能确定 int 的 bit 就选择比较明确的 int 类型，因为这会让你的程序具备很好的移植性。

 **数字字面量语法**:

2进制：以`0b`或`0B`为前缀

```GO
var num01 int = 0b1100   // 是 零B 
```

8进制：以`0o`或者 `0O`为前缀

```GO
var num02 int = 0o14
```

16进制：以`0x` 为前缀

```GO
var num03 int = 0xC
```

v := 123_456 等于 123456。


## 浮点型
Go语言支持两种浮点型数：`float32`和`float64`。这两种浮点型数据格式遵循`IEEE 754`标准： `float32` 的浮点数的最大范围约为 `3.4e38`，可以使用常量定义：`math.MaxFloat32`。 `float64` 的浮点数的最大范围约为 `1.8e308`，可以使用一个常量定义：`math.MaxFloat64`。
项目中最常用的是 float64，因为它的精度高，浮点计算的结果相比 float32 误差会更小。

## 复数

complex64和complex128

```go
var c1 complex64
c1 = 1 + 2i
var c2 complex128
c2 = 2 + 3i
fmt.Println(c1)
fmt.Println(c2)
```

复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。

## 布尔型

Go语言中以`bool`类型进行声明布尔型数据，布尔型数据只有`true（真）`和`false（假）`两个值。

**注意：**

1. 布尔类型变量的默认值为`false`。
2. Go 语言中不允许将整型**强制转换**为布尔型.
3. 布尔型无法参与数值运算，也无法与其他类型进行转换。

## byte与 rune

byte，**占用1个节字**，就 8 个比特位，所以它和 `uint8` 类型本质上没有区别，它表示的是 **ACSII 表中的一个字符**。
```go
var a byte = 65
var b uint8 = 66
fmt.Println(a)
fmt.Printf("a的值： %c \n", a)
fmt.Println(b)
fmt.Printf("a的值： %c \n", b)

//结果
65
a的值： A 
66
a的值： B 
```
在 ASCII 表中，由于字母 A 的ASCII 的编号为 65 ，字母 B 的ASCII 编号为 66，所以上面的代码也可以写成这样
```go
var a byte = 'A'
var b uint8 = 'B'
fmt.Print(a)
fmt.Printf("a 的值: %c \nb 的值: %c", a, b)
//结果
65
a 的值: A
b 的值: B
```
**rune**，**占用4个字节**，共32位比特位，所以它和 `uint32` 本质上也没有区别。它表示的是一个 Unicode字符（Unicode是一个可以表示世界范围内的绝大部分字符的编码规范）。
由于 byte 类型能表示的值是有限，只有 2^8=256 个。所以如果你想表示中文的话，你只能使用 rune 类型。
```go
var name rune = '中'
fmt.Println(b)
fmt.Printf("b 的值: %c",  b)
//结果
20013
b 的值: 中
```

byte 和 uint8 没有区别，rune 和 uint32 没有区别，那为什么还要多出一个 byte 和 rune 类型呢？

理由很简单，因为uint8 和 uint32 ，直观上让人以为这是一个数值，但是实际上，它也可以表示一个字符，所以为了消除这种直观错觉，就诞生了 byte 和 rune 这两个**别名类型**。
## 字符串

### 本质

Go语言中的字符串以**原生数据类型**出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用`UTF-8`编码。

string 的本质，其实是一个 **byte数组**

`hello,中国` 占用几个字节？

Go 语言的 string 是用 uft-8 进行编码的，**英文字母占用一个字节，而中文字母占用 3个字节**，所以 `hello,中国` 的长度为 5+1+（3＊2)= 12个字节。

**除了双引号之外 ，你还可以使用反引号。**

大多情况下，二者并没有区别，但如果你的字符串中有转义字符`\` ，这里就要注意了，它们是有区别的。

比如我想表示 `\r\n` 这个 字符串，使用双引号是这样写的，这种叫解释型表示法
```go
var mystr01 string = "\\r\\n"   //// 因为\r \n是转义字符
var mystr02 string = `\r\n`
```
同时反引号可以**不写换行符**（因为没法写）来表示一个多行的字符串。
```go
    var a string = `ssss
yyyyyyy`
   var b string = "wwww\n yyy"
   fmt.Println(a)
   fmt.Println(b)

// output
ssss
yyyyyyy
wwww
 yyy
```

### 修改字符串
要修改字符串，需要先将其转换成`[]rune`或`[]byte`，完成后再转换为`string`。无论哪种转换，都会**重新分配内存，并复制字节数组**。

```go
func changeString() {
	//纯粹英文
	s := "zhiguogg"
	// 强制类型转换
	byteS := []byte(s)
	byteS[0] = 'c'
	fmt.Println(string(byteS))

	//纯粹中文
	s1 := "我不是归人"
	runeS := []rune(s1)
	runeS[0] = '你'
	fmt.Println(string(runeS))
}
```

既有中文又有英文怎么处理？

首先看转换成byte数组是什么样的

```go
	s2 := "hello,中国"
	byteS2 := []byte(s2)
	fmt.Println(byteS2)
// [104 101 108 108 111 44 228 184 173 229 155 189]

```

可以看见 后面的“中国”每一个字符都是三个字节表示

如果要改变英文字符，只需要更改一个字节就行，而修改中文字符则需要三个变动

```go
byteS2[0] = 'l'
	byteS2[9] = 228
	byteS2[10] = 184
	byteS2[11] = 173
	//我们把第一个英文字符改成 l，把后面代表'国'的三个字节更改成和'中'一样 预测为 lello,中中
	fmt.Println(string(byteS2))
// lello,中中
```

显然是比较麻烦的，看看转换成rune数组看看

```go
s3 := "hello,中国"
	runeS3 := []rune(s3)
	fmt.Println(runeS3)
// [104 101 108 108 111 44 20013 22269]
```

可以看到 前面英文字符的没有改变，rune兼容了type。即Unicode字符兼容了ACSII码。

此时可以看到更改较为方便

```go
runeS3[0] = 'l'
	runeS3[7] = '中'
	fmt.Println(string(runeS3))
// lello,中中
```

 

一个值在从string类型向[]byte类型转换时代表着以 UTF-8 编码的字符串会被拆分成零散、独立的字节。除了与 ASCII 编码兼容的那部分字符集，以 UTF-8 编码的某个单一字节是无法代表一个字符的。

```go
string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}) // 你好
```

比如，UTF-8 编码的三个字节\xe4、\xbd和\xa0合在一起才能代表字符'你'，而\xe5、\xa5和\xbd合在一起才能代表字符'好'。

其次，一个值在从string类型向[]rune类型转换时代表着字符串会被拆分成一个个 Unicode 字符。

```go
string([]rune{'\u4F60', '\u597D'}) // 你好
```

