# 字典或映射 即Map
Go语言中的map是**引用类型，必须初始化才能使用。**

## map定义
Go语言中 `map`的定义语法如下：

```go
map[KeyType]ValueType
```

其中，

- KeyType:表示键的类型。
- ValueType:表示键对应的值的类型。

map类型的变量默认初始值为nil，**需要使用make()函数来分配内存**。语法为：

```go
make(map[KeyType]ValueType, [cap])
```

其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

除了可以通过 make 函数创建 map 外，还可以通过字面量的方式。

```go
nameAgeMap:=map[string]int{"我不是归人":20}
```

在创建 map 的同时添加键值对，如果不想添加键值对，使用空大括号 {} 即可，要注意的是，**大括号一定不能省略**。

## map基本使用
map中的数据都是成对出现的，map的基本使用示例代码如下：

```go
func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)
}
```

输出：

```bash
map[小明:100 张三:90]
100
type of a:map[string]int
```

map也支持在声明的时候填充元素，例如：

```go
func main() {
	userInfo := map[string]string{
		"username": "我不是归人",
		"password": "123456",
	}
	fmt.Println(userInfo) //
}
```

## 判断某个键是否存在
Go语言中有个判断map中键是否存在的特殊写法，格式如下:

```go
value, ok := map[key]
```

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
```

## map的遍历
Go语言中使用`for range`遍历map。

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}
```

但我们只想遍历key的时候，可以按下面的写法：

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k := range scoreMap {
		fmt.Println(k)
	}
}
```

## 使用delete()函数删除键值对
使用`delete()`内建函数从map中删除一组键值对，`delete()`函数的格式如下：

```go
delete(map, key)
```

## 按照指定顺序遍历map
```go
func traverseMapByOrder()  {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```
## Map 的大小

和数组切片不一样，map 是没有容量的，它只有长度，也就是 map 的大小（键值对的个数）。要获取 map 的大小，使用内置的 len 函数即可，如下代码所示：

```go
fmt.Println(len(nameAgeMap))
```

## 为什么字典的键类型会受到约束
Go 语言的字典类型其实是一个哈希表（hash table）的特定实现，在这个实现中，键和元素的最大不同在于，**键的类型是受限的，而元素却可以是任意类型的**。

**映射过程的第一步就是：把键值转换为哈希值。**

在 Go 语言的字典中，每一个键值都是由它的哈希值代表的。也就是说，**字典不会独立存储任何键的值，但会独立存储它们的哈希值。**

**字典的键类型不能是哪些类型？**

**它的典型回答是：Go 语言字典的键类型不可以是函数类型、字典类型和切片类型。**

Go 语言规范规定，在**键类型的值之间必须可以施加操作符==和!=**。换句话说，键类型的值必须要支持判等操作。由于函数类型、字典类型和切片类型的值并不支持判等操作，所以字典的键类型不能是这些类型。

另外，如果**键的类型是接口类型的，那么键值的实际类型也不能是上述三种类型**，否则在程序运行过程中会引发 panic（即运行时恐慌）。

```go
var badMap2 = map[interface{}]int{
  "1":   1,
  []int{2}: 2, // 这里会引发panic。
  3:    3,
}
```

这里的变量badMap2的类型是键类型为interface{}、值类型为int的字典类型。这样声明并不会引起什么错误。或者说，我通过这样的声明躲过了 Go 语言编译器的检查。

但是，当我们运行这段代码的时候，Go 语言的运行时（runtime）系统就会发现这里的问题，它会抛出一个 panic，并把根源指向字面量中定义第二个键 - 元素对的那一行。我们越晚发现问题，修正问题的成本就会越高，所以最好**不要把字典的键类型设定为任何接口类型**。如果非要这么做，请一定确保代码在可控的范围之内。

还要注意，如果键的类型是**数组类型**，那么还要确保该类型的**元素类型不是函数类型、字典类型或切片类型。**另外，如果键的类型是**结构体类型，那么还要保证其中字段的类型的合法性**。无论不合法的类型被埋藏得有多深。

为什么键类型的值必须支持判等操作？

如果有相等的，那就再用键值本身去对比一次。为什么还要对比？原因是，不同值的哈希值是可能相同的。这有个术语，叫做“哈希碰撞”。

所以，即使哈希值一样，键值也不一定一样。如果键类型的值之间无法判断相等，那么此时这个映射的过程就没办法继续下去了。最后，只有键的哈希值和键值都相等，才能说明查找到了匹配的键 - 元素对。

## 在值为nil的字典上执行读操作会成功吗，那写操作呢

由于字典是引用类型，所以当我们仅声明而不初始化一个字典类型的变量的时候，它的值会是nil。

**除了添加键 - 元素对，我们在一个值为nil的字典上做任何操作都不会引起错误**。当我们试图在一个值为nil的字典中添加键 - 元素对的时候，Go 语言的运行时系统就会立即抛出一个 panic。






