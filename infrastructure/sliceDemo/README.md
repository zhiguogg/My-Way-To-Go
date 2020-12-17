# 切片（slice）

切片（Slice）是一个拥有相同类型元素的**可变长度**的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个**引用类型**，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。

切片和数组类似，可以把它理解为动态数组。切片是基于数组实现的，它的底层就是一个数组。对数组任意分隔，就可以得到一个切片。

## 切片声明
```go
func declareSlice()  {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	fmt.Println(d)
	 //fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}
```

## 切片的长度和容量
切片拥有自己的长度和容量，我们可以通过使用内置的`len()`函数求长度，使用内置的`cap()`函数求切片的容量。

**切片声明之后一定要初始化才能使用**

## 切片表达式
#### 通过字面量的方式声明和初始化

```go
slice1:=[]string{"a","b","c","d","e"}
fmt.Println(len(slice1),cap(slice1))
```

切片和数组的字面量初始化方式，差别就是中括号 [] 里的长度。此外，**通过字面量初始化的切片，长度和容量相同。
### 简单切片表达式

切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。 切片表达式中的`low`和`high`表示一个索引范围（**左包含，右不包含**），也就是下面代码中从数组a中选出`1<=索引值<3`的元素组成切片s，得到的切片`长度=high-low`，容量等于得到的切片的**底层数组的容量**。

```GO
func main() {
	a := []int{0,1,2,3,4,5}
    	s := a[1:3]  //  1 2
    	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))   //s:[1 2] len(s):2 cap(s):5
}

```



为了方便起见，可以省略切片表达式中的任何索引。省略了`low`则默认为0；省略了`high`则默认为切片操作数的长度:

```go
a[2:]  // 等同于 a[2:len(a)]
a[:3]  // 等同于 a[0:3]
a[:]   // 等同于 a[0:len(a)]
```

**注意：**

对于数组或字符串，如果`0 <= low <= high <= len(a)`，则索引合法，否则就会索引越界（out of range）。

对切片再执行切片表达式时（切片再切片），`high`的上限边界是切片的**容量**`cap(a)`，而不是长度。**常量索引**必须是非负的，并且可以用int类型的值表示;对于数组或常量字符串，常量索引也必须在有效范围内。如果`low`和`high`两个指标都是常数，它们必须满足`low <= high`。如果索引在运行时超出范围，就会发生运行时`panic`。
```go
a := []int{0,1,2,3,4,5}
	s := a[1:3]  //  1 2
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))   //s:[1 2] len(s):2 cap(s):5
	s2 := s[0:5]
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))  //s2:[1 2 3 4 5] len(s2):5 cap(s2):5
```

### 完整切片表达式

对于数组，指向数组的指针，或切片a(**注意不能是字符串**)支持完整切片表达式：

```go
a[low : high : max]
```

上面的代码会构造与简单切片表达式`a[low: high]`相同类型、相同长度和元素的切片。另外，它会将得到的结果切片的容量设置为`max-low`。在完整切片表达式中**只有第一个索引值（low）可以省略；它默认为0。**

```go
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}

//结果
t:[2 3] len(t):2 cap(t):4
```

完整切片表达式需要满足的条件是`0 <= low <= high <= max <= cap(a)`，其他条件和简单切片表达式相同。

```GO
s := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
t := s[0:5]
m := s[4:8]
fmt.Printf("s:%v len(s):%v cap(s):%v\n", t, len(t), cap(t)) // 5 10
fmt.Printf("s:%v len(s):%v cap(s):%v\n", m, len(m), cap(m)) // 4 6
n := s[0:5:7]
p := s[4:8:9]
fmt.Printf("s:%v len(s):%v cap(s):%v\n", n, len(n), cap(n)) // 5 7
fmt.Printf("s:%v len(s):%v cap(s):%v\n", p, len(p), cap(p)) // 4 5
r := t[2:7]
fmt.Printf("s:%v len(s):%v cap(s):%v\n", r, len(r), cap(r)) // 5 8
r1 := t[2:7:9]
fmt.Printf("s:%v len(s):%v cap(s):%v\n", r1, len(r1), cap(r1)) // 5 7
```

总结：没有第三个参数，默认从第一个切点到原数组**最后**

​				若有第三个参数，从第一个切点到第三个参数，遵循**左闭右开原则**

### 使用make()函数构造切片

我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的`make()`函数，格式如下：

```bash
make([]T, size, cap)
```

其中：

- T:切片的元素类型
- size:切片中元素的数量
- cap:切片的容量

举个例子：

```go
func main() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}
```

上面代码中`a`的内部存储空间已经分配了10个，但实际上只用了2个。 容量并不会影响当前元素的个数，所以`len(a)`返回2，`cap(a)`则返回该切片的容量。

## 切片的本质
切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。
我们其实可以把切片看做是对数组的一层简单的封装，因为在每个切片的底层数据结构中，**一定会包含一个数组**。数组可以被叫做切片的底层数组，而切片也可以被看作是对数组的某个连续片段的引用。

> 也正因为如此，Go 语言的**切片类型**属于引用类型，同属引用类型的还有字典类型、通道类型、函数类型。接口类型等；而 Go 语言的**数组类型**则属于值类型，同属值类型的有**基础数据类型以及结构体类型**。注意，Go 语言里不存在像 Java 等编程语言中令人困惑的“传值或传引用”问题。在 Go 语言中，我们判断所谓的“传值”或者“传引用”只要看**被传递的值的类型**就好了。如果传递的值是引用类型的，那么就是“传引用”。如果传递的值是值类型的，那么就是“传值”。从传递成本的角度讲，**引用类型的值往往要比值类型的值低很多**。我们在数组和切片之上都可以应用索引表达式，得到的都会是某个元素。我们在它们之上也都可以应用切片表达式，也都会得到一个新的切片。

用make函数初始化切片时，如果不指明其容量，那么它就会和长度一致。如果在初始化时指明了容量，那么切片的实际容量也就是它了。



可以把切片看做是对数组的一层简单的封装，因为在每个切片的底层数据结构中，一定会包含一个数组。数组可以被叫做切片的底层数组，而切片也可以被看作是对数组的某个连续片段的引用。在这种情况下，切片的容量实际上代表了它的底层数组的长度，这里是8。（注意，切片的**底层数组**等同于我们前面讲到的数组，其**长度不可变**。）

**有一个窗口，你可以通过这个窗口看到一个数组，但是不一定能看到该数组中的所有元素，有时候只能看到连续的一部分元素。**

这个数组就是切片s2的底层数组，而这个窗口就是切片s2本身。s2的长度实际上指明的就是这个窗口的宽度，决定了你透过s2，可以看到其底层数组中的哪几个连续的元素。

当我们用make函数或切片值字面量（比如[]int{1, 2, 3}）初始化一个切片时，该窗口最左边的那个小格子总是会对应其底层数组中的第 1 个元素。

但是当我们通过切片表达式基于某个数组或切片生成新切片的时候，情况就变得复杂起来了。

```go
s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
s4 := s3[3:6]
fmt.Printf("The length of s4: %d\n", len(s4))
fmt.Printf("The capacity of s4: %d\n", cap(s4))
fmt.Printf("The value of s4: %d\n", s4)
```
切片s3中有 8 个元素，分别是从1到8的整数。s3的长度和容量都是8。然后，我用切片表达式s3[3:6]初始化了切片s4。问题是，这个s4的长度和容量分别是多少？
再来看容量。我在前面说过，**切片的容量代表了它的底层数组的长度，但这仅限于使用make函数或者切片值字面量初始化切片的情况。**

更通用的规则是：**一个切片的容量可以被看作是透过这个窗口最多可以看到的底层数组中元素的个数。**

由于s4是通过在s3上施加切片操作得来的，所以s3的底层数组就是s4的底层数组。又因为，在底层数组不变的情况下，切片代表的窗口可以向右扩展，直至其底层数组的末尾。所以，s4的容量就是其底层数组的长度8, 减去上述切片表达式中的那个起始索引3，即5。

注意，**切片代表的窗口是无法向左扩展的**。也就是说，我们永远无法透过s4看到s3中最左边的那 3 个元素。

最后，顺便提一下把切片的窗口向右扩展到最大的方法。对于s4来说，切片表达式s4[0:cap(s4)]就可以做到。我想你应该能看懂。该表达式的结果值（即一个新的切片）会是[]int{4, 5, 6, 7, 8}，其长度和容量都是5。

## 切片的修改
对切片相应的索引元素赋值就是修改
```go
func changeSlice()  {
	a := []int{1,2,3}
	a[1] = 12
	fmt.Println(a)

	b := []interface{}{1,"空接口可以存储任何数据",false}
	fmt.Println(b)
}
```
切片改变了，底层数组是否改变？
```go
array:=[5]string{"a","b","c","d","e"}
slice:=array[2:5]
slice[1] ="f"
fmt.Println(array)
//打印结果
[a b c f e]
```
数组对应的值已经被修改为 f，所以这也证明了**基于数组的切片，使用的底层数组还是原来的数组，一旦修改切片的元素值，那么底层数组对应的值也会被修改。**

## 判断切片是否为空
要检查切片是否为空，**请始终使用`len(s) == 0`来判断**，而不应该使用`s == nil`来判断。
### 切片不能直接比较

切片之间是不能比较的，我们不能使用`==`操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和`nil`比较。 一个`nil`值的切片**并没有底层数组**，一个`nil`值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是`nil`

### 切片的赋值拷贝

下面的代码中演示了拷贝前后两个变量**共享底层数组**，对一个切片的修改会影响另一个切片的内容，这点需要特别注意

```go
func sliceAssignment() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
}
```
## 切片遍历

切片的遍历方式和数组是一致的，支持索引遍历和`for range`遍历。

```go
func main() {
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
```

## append()方法为切片添加元素

```go
func main(){
	var s []int   //注意这里是声明 但是没有初始化 此时打印显示  []
	s = append(s, 1)        // [1]
	s = append(s, 2, 3, 4)  // [1 2 3 4]
	s2 := []int{5, 6, 7}  
	s = append(s, s2...)    // [1 2 3 4 5 6 7]
}
```

**通过var声明的零值切片可以在`append()`函数直接使用，无需初始化。**

```go
var s []int
s = append(s, 1, 2, 3)
```

**每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在`append()`函数调用时，所以我们通常都需要用原变量接收append函数的返回值。如果没超过底层数组，则修改切片值底层数组值也会改变；如果超过则无影响**

```go
func main() {
	//append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}


// append 函数扩容导致底层数组是否变动
	array2 := [3]int{0, 1, 2}
	array3 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var ss1 = array2[0:1]
	fmt.Println(ss1) //[0]
	ss1 = append(ss1, 6, 6, 6)
	fmt.Println(ss1)    //[0 6 6 6]
	fmt.Println(array2) //[0 1 2] 如果没超过底层数组，则修改切片值底层数组值也会改变；如果超过则无影响
	var ss2 = array3[0:4]
	fmt.Println(ss2) //[0 1 2 3]
	ss2 = append(ss2, 6, 6, 6)
	fmt.Println(ss2)  //[0 1 2 3 6 6 6]
	fmt.Println(array3) //[0 1 2 3 6 6 6 7 8 9]  如果没超过底层数组，则修改切片值底层数组值也会改变；如果超过则无影响
```

> 小技巧：在创建新切片的时候，最好要让新切片的长度和容量一样，这样在追加操作的时候就会生成新的底层数组，从而和原有数组分离，就不会因为共用底层数组导致修改内容的时候影响多个切片。

一旦一个切片无法容纳更多的元素，Go 语言就会想办法扩容。但它并不会改变原来的切片，而是会生成一个容量更大的切片，然后将把原有的元素和新元素一并拷贝到新切片中。

一般的情况下，你可以简单地认为新切片的容量（以下简称新容量）将会是原切片容量（以下简称原容量）的 2 倍。

但是，当原切片的长度（以下简称原长度）大于或等于1024时，Go 语言将会以原容量的1.25倍作为新容量的基准（以下新容量基准）

另外，如果我们一次追加的元素过多，以至于使新长度比原容量的 2 倍还要大，那么新容量就会以新长度为基准。注意，与前面那种情况一样，最终的新容量在很多时候都要比新容量基准更大一些。

扩容之后在源码分析下。

## 使用copy()函数复制切片
首先我们来看一个问题：

```go
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]
}
```

**由于切片是引用类型，所以a和b其实都指向了同一块内存地址**。修改b的同时a的值也会发生变化。

Go语言内建的`copy()`函数可以迅速地将一个切片的数据复制到另外一个切片空间中，`copy()`函数的使用格式如下：

```bash
copy(destSlice, srcSlice []T)
```

其中：

- srcSlice: 数据来源切片
- destSlice: 目标切片

举个例子：

```go
func main() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}
```

## 从切片中删除元素
Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。 代码如下：

```go
func main() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}
```

总结一下就是：**要从切片a中删除索引为`index`的元素，操作方法是`a = append(a[:index], a[index+1:]...)`**

## 二维切片
Go 的数组和切片都是一维的。要创建等价的二维数组或切片，就必须定义一个数组的数组， 或切片的切片，就像这样：

```go
type Transform [3][3]float64  // 一个 3x3 的数组，其实是包含多个数组的一个数组。
type LinesOfText [][]byte     // 包含多个字节切片的一个切片。
```

由于切片长度是可变的，因此其内部可能**拥有多个不同长度的切片**。在我们的 `LinesOfText` 例子中，这是种常见的情况：每行都有其自己的长度。

```go
text := LinesOfText{
    []byte("Now is the time"),
    []byte("for all good gophers"),
    []byte("to bring some fun to the party."),
}
```

有时必须分配一个二维数组，例如在处理像素的扫描行时，这种情况就会发生。 我们有两种方式来达到这个目的。一种就是独立地分配每一个切片；而另一种就是只分配一个数组， 将各个切片都指向它。采用哪种方式取决于你的应用。若切片会增长或收缩， 就应该通过独立分配来避免覆盖下一行；若不会，用单次分配来构造对象会更加高效。 以下是这两种方法的大概代码，仅供参考。首先是一次一行的：

```go
// 分配底层切片.
picture := make([][]uint8, YSize) // y每一行的大小
//循环遍历每一行
for i := range picture {
    picture[i] = make([]uint8, XSize)
}
```

现在是一次分配，对行进行切片：

```go
// 分配底层切片
picture := make([][]uint8, YSize) //  每 y 个单元一行。
// 分配一个大一些的切片以容纳所有的元素
pixels := make([]uint8, XSize*YSize) // 指定类型[]uint8, 即便图片是 [][]uint8.
//循环遍历图片所有行，从剩余像素切片的前面对每一行进行切片。
for i := range picture {
    picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```
 

## 切片的底层数组什么时候会被替换
**一个切片的底层数组永远不会被替换。为什么？虽然在扩容的时候 Go 语言一定会生成新的底层数组，但是它也同时生成了新的切片。**

请记住，在无需扩容时，append函数返回的是指向原底层数组的原切片，而在需要扩容时，append函数返回的是指向新底层数组的新切片。

只要新长度不会超过切片的原容量，那么使用append函数对其追加元素的时候就不会引起扩容。这只会使紧邻切片窗口右边的（底层数组中的）元素被新的元素替换掉


## 如果有多个切片指向了同一个底层数组，那么你认为应该注意些什么？
我们需要特别注意的是，当操作其中一个切片的时候是否会影响到其他指向同一个底层数组的切片。

如果是，那么问一下自己，这是你想要的结果吗？无论如何，**通过这种方式来组织或共享数据是不正确的**。你需要做的是，**要么彻底切断这些切片的底层联系，要么立即为所有的相关操作加锁**。

## 怎样沿用“扩容”的思想对切片进行“缩容”？请写出代码。
关于切片的“缩容”，可参看官方的相关 wiki。不过，如果你需要频繁的“缩容”，那么就可能需要考虑其他的数据结构了，比如：container/list代码包中的List。


