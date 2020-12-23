# 并发
## 并发与并行

并发：同一**时间段**内执行多个任务

并行：同一**时刻**执行多个任务

并发：同一**时间段**内执行多个任务

并行：同一**时刻**执行多个任务

Go语言的并发通过`goroutine`实现。`goroutine`类似于线程，属于**用户态**的线程，我们可以根据需要创建成千上万个`goroutine`并发工作。**`goroutine`是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成。**

Go语言还提供`channel`在多个`goroutine`间进行通信。`goroutine`和`channel`是 Go 语言秉承的 CSP（Communicating Sequential Process）并发模式的重要实现基础
## 进程和线程

### 进程

在操作系统中，进程是一个非常重要的概念。当你启动一个软件（比如浏览器）的时候，操作系统会为这个软件创建一个进程，这个进程是该软件的工作空间，它包含了软件运行所需的所有资源，比如内存空间、文件句柄，还有下面要讲的线程等

### 线程

线程是进程的执行空间，一个进程可以有多个线程，线程被**操作系统调度执行**，比如下载一个文件，发送一个消息等。这种多个线程被操作系统同时调度执行的情况，就是多线程的并发。

一个程序启动，就会有对应的进程被创建，同时进程也会启动一个线程，这个线程叫作**主线程**。如果主线程结束，那么整个程序就退出了。有了主线程，就可以从主线里启动很多其他线程，也就有了多线程的并发。

## 协程（Goroutine）

`goroutine`的概念类似于线程，但 `goroutine`是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在**语言层面已经内置了调度和上下文切换的机制。**



相比线程来说，协程更加轻量，一个程序可以随意启动成千上万个 goroutine。

goroutine 被 Go runtime 所调度，这一点和线程不一样。也就是说，**Go 语言的并发是由 Go 自己所调度的**，自己决定同时执行多少个 goroutine，什么时候执行哪几个。这些对于我们开发者来说完全透明，只需要在编码的时候告诉 Go 语言要启动几个 goroutine，至于如何调度执行，我们不用关心。

## goroutine与线程

### 可增长的栈

OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）,一个`goroutine`的栈在其生命周期开始时只有很小的栈（典型情况下2KB），`goroutine`的栈不是固定的，他可以按需增大和缩小，`goroutine`的栈大小限制可以达到1GB，虽然极少会用到这么大。所以在Go语言中一次创建十万左右的`goroutine`也是可以的。
### goroutine调度

`GPM`是Go语言运行时（runtime）层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。

- `G`很好理解，就是个goroutine的，里面除了存放本goroutine信息外 还有与所在P的绑定等信息。
- `P`管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。
- `M（machine）`是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系， 一个groutine最终是要放到M上执行的；

## channel

单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

虽然可以使用共享内存进行数据交换，但是共享内存在不同的`goroutine`中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

Go语言的并发模型是`CSP（Communicating Sequential Processes）`，提倡**通过通信共享内存而不是通过共享内存而实现通信**。

如果说`goroutine`是Go程序并发的执行体，`channel`就是它们之间的连接。`channel`是可以让一个`goroutine`发送特定值到另一个`goroutine`的通信机制。

Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循**先入先出**（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

### channel类型

### channel类型

`channel`是一种类型，一种**引用类型**。声明通道类型的格式如下：

```go
var 变量 chan 元素类型
```

举几个例子：

```go
var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道
```

### 创建channel
通道是引用类型，通道类型的空值是`nil`。

```go
var ch chan int
fmt.Println(ch) // <nil>
```

声明的通道后需要使用`make`函数初始化之后才能使用。
创建channel的格式如下：

```go
make(chan 元素类型, [缓冲大小])
```

channel的缓冲大小是可选的。

举几个例子：

```go
ch4 := make(chan int)
ch5 := make(chan bool)
ch6 := make(chan []int)
```

### channel操作

通道有发送（send）、接收(receive）和关闭（close）三种操作。
发送和接收都使用`<-`符号。

现在我们先使用以下语句定义一个通道：

```go
ch := make(chan int,10)
```

#### 发送
将一个值发送到通道中。
```go
ch <- 10 // 把10发送到ch中
```
#### 接收

从一个通道中接收值。

```go
x := <- ch // 从ch中接收值并赋值给变量x
<-ch       // 从ch中接收值，忽略结果
```

#### 关闭

我们通过调用内置的`close`函数来关闭通道。

```go
close(ch)
```

关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都**发送完毕**的时候才需要关闭通道。通道是可以被**垃圾回收机制回收**的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但**关闭通道不是必须**的。

关闭后的通道有以下特点：

1. 对一个关闭的通道再发送值就会导致panic。
2. 对一个关闭的通道进行接收会一直获取值直到通道为空。
3. 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
4. 关闭一个已经关闭的通道会导致panic。

### 无缓冲的通道

无缓冲的通道又称为阻塞的通道。我们来看一下下面的代码：

```go
func main() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}
```

上面这段代码能够通过编译，但是执行的时候会出现以下错误：

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        .../src/github.com/Q1mi/studygo/day06/channel02/main.go:8 +0x54
```

为什么会出现`deadlock`错误呢？

因为我们使用`ch := make(chan int)`创建的是无缓冲的通道，无缓冲的通道只**有在有人接收值的时候才能发送值**。简单来说就是**无缓冲的通道必须有接收才能发送**。

上面的代码会阻塞在`ch <- 10`这一行代码形成死锁，那如何解决这个问题呢？

一种方法是启用一个`goroutine`去接收值，例如：

```go
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
```

无缓冲通道上的发送操作会阻塞，直到另一个`goroutine`在该通道上执行接收操作，这时值才能发送成功，两个`goroutine`将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个`goroutine`在该通道上发送一个值。

使用无缓冲通道进行通信将导致发送和接收的`goroutine`同步化。因此，无缓冲通道也被称为`同步通道`。

### 有缓冲的通道

解决上面问题的方法还有一种就是使用有缓冲区的通道。我们可以在使用make函数初始化通道的时候为其指定通道的容量，例如：

```go
func main() {
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功")
}
```

只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。

我们可以使用内置的`len`函数获取通道内元素的数量，使用`cap`函数获取通道的容量，虽然我们很少会这么做。

### for range从通道循环取值

当向通道中发送完数据时，我们可以通过`close`函数来关闭通道。

当通道被关闭时，再往该通道发送值会引发`panic`，从该通道取值的操作会先取完通道中的值，再然后取到的值一直都是对应类型的零值。那如何判断一个通道是否被关闭了呢？

我们来看下面这个例子：
```go
// channel 练习
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}
```
从上面的例子中我们看到有两种方式在接收值的时候判断该通道是否被关闭，不过我们通常使用的是`for range`的方式。使用`for range`遍历通道，当通道被关闭的时候就会退出`for range`

### 单向通道

有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。

Go语言中提供了**单向通道**来处理这种情况。例如，我们把上面的例子改造如下：

```go
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
```
其中，

- `chan<- int`是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
- `<-chan int`是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。

在**函数传参及任何赋值操作中可以将双向通道转换为单向通道**，但反过来是不可以的。

### 对通道的发送和接收操作都有哪些基本的特性？

它们的基本特性如下：

1. 对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的。
2. 发送操作和接收操作中对元素值的处理都是不可分割的。
3. 发送操作在完全完成之前会被阻塞。接收操作也是如此。

我们先来看第一个基本特性。 在同一时刻，Go 语言的运行时系统（以下简称运行时系统）只会执行对同一个通道的任意个发送操作中的某一个。

直到这个元素值被完全复制进该通道之后，其他针对该通道的发送操作才可能被执行。

类似的，在同一时刻，运行时系统也只会执行，对同一个通道的任意个接收操作中的某一个。直到这个元素值完全被移出该通道之后，其他针对该通道的接收操作才可能被执行。即使这些操作是并发执行的也是如此。

这里所谓的并发执行，你可以这样认为，多个代码块分别在不同的 goroutine 之中，并有机会在同一个时间段内被执行。

另外，对于通道中的同一个元素值来说，发送操作和接收操作之间也是互斥的。例如，虽然会出现，正在被复制进通道但还未复制完成的元素值，但是这时它绝不会被想接收它的一方看到和取走。

**这里要注意的一个细节是，元素值从外界进入通道时会被复制。更具体地说，进入通道的并不是在接收操作符右边的那个元素值，而是它的副本**。

另一方面，元素值从通道进入外界时会被移动。这个移动操作实际上包含了两步，**第一步是生成正在通道中的这个元素值的副本，并准备给到接收方，第二步是删除在通道中的这个元素值。**

顺着这个细节再来看第二个基本特性。 这里的“不可分割”的意思是，**它们处理元素值时都是一气呵成的，绝不会被打断。**

例如，发送操作要么还没复制元素值，要么已经复制完毕，绝不会出现只复制了一部分的情况。又例如，接收操作在准备好元素值的副本之后，一定会删除掉通道中的原值，绝不会出现通道中仍有残留的情况。这既是为了保证通道中元素值的完整性，也是为了保证通道操作的唯一性。对于通道中的同一个元素值来说，它只可能是某一个发送操作放入的，同时也只可能被某一个接收操作取出。

再来说第三个基本特性。 一般情况下，发送操作包括了“复制元素值”和“放置副本到通道内部”这两个步骤。

**在这两个步骤完全完成之前，发起这个发送操作的那句代码会一直阻塞在那里。也就是说，在它之后的代码不会有执行的机会，直到这句代码的阻塞解除。**

更细致地说，在通道完成发送操作之后，运行时系统会通知这句代码所在的 goroutine，以使它去争取继续运行代码的机会。

另外，接收操作通常包含了“复制通道内的元素值”“放置副本到接收方”“删掉原值”三个步骤。

在所有这些步骤完全完成之前，发起该操作的代码也会一直阻塞，直到该代码所在的 goroutine 收到了运行时系统的通知并重新获得运行机会为止。

说到这里，你可能已经感觉到，**如此阻塞代码其实就是为了实现操作的互斥和元素值的完整。**

### 发送操作和接收操作在什么时候可能被长时间的阻塞？

先说针对**缓冲通道**的情况。如果通道已满，那么对它的所有发送操作都会被阻塞，直到通道中有元素值被接收走。

这时，通道会**优先通知最早**因此而等待的、那个发送操作所在的 goroutine，后者会再次执行发送操作。

由于发送操作在这种情况下被阻塞后，它们所在的 goroutine 会顺序地进入通道内部的发送**等待队列**，所以通知的顺序总是公平的。

相对的，如果通道已空，那么对它的所有接收操作都会被阻塞，直到通道中有新的元素值出现。这时，通道会通知最早等待的那个接收操作所在的 goroutine，并使它再次执行接收操作。因此而等待的、所有接收操作所在的 goroutine，都会按照先后顺序被放入通道内部的接收等待队列。

对于**非缓冲通道**，情况要简单一些。无论是发送操作还是接收操作，**一开始执行就会被阻塞**，直到配对的操作也开始执行，才会继续传递。由此可见，非缓冲通道是在用同步的方式传递数据。也就是说，只有收发双方对接上了，数据才会被传递。

**并且，数据是直接从发送方复制到接收方的，中间并不会用非缓冲通道做中转。相比之下，缓冲通道则在用异步的方式传递数据。**

在大多数情况下，缓冲通道会作为收发双方的中间件。正如前文所述，元素值会先从发送方复制到缓冲通道，之后再由缓冲通道复制给接收方。

但是，当发送操作在执行的时候发现空的通道中，正好有等待的接收操作，那么它会直接把元素值复制给接收方。

以上说的都是在正确使用通道的前提下会发生的事情。下面我特别说明一下，由于错误使用通道而造成的阻塞。

对于**值为nil的通道**，不论它的具体类型是什么，对它的发送操作和接收操作都会**永久地处于阻塞状态**。它们所属的 goroutine 中的任何代码，都不再会被执行。

注意，由于通道类型是引用类型，所以它的零值就是nil。换句话说，当我们只声明该类型的变量但没有用make函数对它进行初始化时，该变量的值就会是nil。我们一定不要忘记初始化通道！

### 发送操作和接收操作在什么时候会引发 panic？

**对于一个已初始化，但并未关闭的通道来说，收发操作一定不会引发 panic。但是通道一旦关闭，再对它进行发送操作，就会引发 panic。**

另外，如果我们试图关闭一个已经关闭了的通道，也会引发 panic。注意，接收操作是可以感知到通道的关闭的，并能够安全退出。

更具体地说，当我们把接收表达式的结果同时赋给两个变量时，第二个变量的类型就是一定bool类型。它的值如果为false就说明通道已经关闭，并且再没有元素值可取了

**注意，如果通道关闭时，里面还有元素值未被取出，那么接收表达式的第一个结果，仍会是通道中的某一个元素值，而第二个结果值一定会是true。**

因此，通过接收表达式的第二个结果值，来判断通道是否关闭是可能有延时的。

**由于通道的收发操作有上述特性，所以除非有特殊的保障措施，我们千万不要让接收方关闭通道，而应当让发送方做这件事。**

### 通道的长度代表着什么？它在什么时候会通道的容量相同？

通道的长度代表它当前包含的元素值的个数。当通道已满时，其长度会与容量相同。

### 元素值在经过通道传递时会被复制，那么这个复制是浅表复制还是深层复制呢？

浅表复制。实际上，在 Go 语言中并不存在深层次的复制，除非我们自己来做

### 单向通道有什么应用价值？
**单向通道最主要的用途就是约束其他代码的行为。**

在实际场景中，这种约束一般会出现在接口类型声明中的某个方法定义上。请看这个叫Notifier的接口类型声明：

```go
type Notifier interface {
  SendInt(ch chan<- int)
}
```

在接口类型声明的花括号中，每一行都代表着一个方法的定义。接口中的方法定义与函数声明很类似，但是只包含了方法名称、参数列表和结果列表。

**一个类型如果想成为一个接口类型的实现类型，那么就必须实现这个接口中定义的所有方法。因此，如果我们在某个方法的定义中使用了单向通道类型，那么就相当于在对它的所有实现做出约束。**

在这里，Notifier接口中的SendInt方法只会接受一个发送通道作为参数，所以，在该接口的所有实现类型中的SendInt方法都会受到限制。这种约束方式还是很有用的，尤其是在我们编写模板代码或者可扩展的程序库的时候。

顺便说一下，我们在调用SendInt函数的时候，只需要把一个元素类型匹配的双向通道传给它就行了，没必要用发送通道，**因为 Go 语言在这种情况下会自动地把双向通道转换为函数所需的单向通道。**

在另一个方面，我们还可以在函数声明的结果列表中使用单向通道。如下所示：

```go
func getIntChan() <-chan int {
  num := 5
  ch := make(chan int, num)
  for i := 0; i < num; i++ {
    ch <- i
  }
  close(ch)
  return ch
}
```

函数getIntChan会返回一个<-chan int类型的通道，这就意味着得到该通道的程序，只能从通道中接收元素值。这实际上就是对函数调用方的一种约束了。

另外，我们在 Go 语言中还可以声明函数类型，如果我们在函数类型中使用了单向通道，那么就相等于在约束所有实现了这个函数类型的函数。

我们再顺便看一下调用getIntChan的代码：

```go
intChan2 := getIntChan()
for elem := range intChan2 {
  fmt.Printf("The element in intChan2: %v\n", elem)
}
```

我把调用getIntChan得到的结果值赋给了变量intChan2，然后用for语句循环地取出了该通道中的所有元素值，并打印出来。这里的for语句也可以被称为带有range子句的for语句。它的用法我在后面讲for语句的时候专门说明。现在你只需要知道关于它的三件事：

1. 上述for语句会不断地尝试从通道intChan2中取出元素值**。即使intChan2已经被关闭了，它也会在取出所有剩余的元素值之后再结束执行。**
2. 通常，当通道intChan2中没有元素值时，**这条for语句会被阻塞在有for关键字的那一行**，直到有新的元素值可取。不过，由于这里的getIntChan函数会事先将intChan2关闭，所以它在取出intChan2中的所有元素值之后会直接结束执行。
3. 倘若通道intChan2的值为nil，那么这条for语句就会被永远地阻塞在有for关键字的那一行。

### select语句与通道怎样联用，应该注意些什么？

**select语句只能与通道联用，它一般由若干个分支组成。每次执行这种语句的时候，一般只有一个分支中的代码会被运行。**

select语句的分支分为两种，一种叫做候选分支，另一种叫做默认分支。候选分支总是以关键字case开头，后跟一个case表达式和一个冒号，然后我们可以从下一行开始写入当分支被选中时需要执行的语句。

默认分支其实就是 default case，因为，当且仅当没有候选分支被选中时它才会被执行，所以它以关键字default开头并直接后跟一个冒号。同样的，我们可以在default:的下一行写入要执行的语句。

由于select语句是专为通道而设计的，所以每个case表达式中都**只能包含操作通道的表达式**，比如接收表达式。

当然，如果我们需要把接收表达式的结果赋给变量的话，还可以把这里写成赋值语句或者短变量声明。下面展示一个简单的例子。

```go
// 准备好几个通道。
intChannels := [3]chan int{
  make(chan int, 1),
  make(chan int, 1),
  make(chan int, 1),
}
// 随机选择一个通道，并向它发送元素值。
index := rand.Intn(3)
fmt.Printf("The index: %d\n", index)
intChannels[index] <- index
// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
select {
case <-intChannels[0]:
  fmt.Println("The first candidate case is selected.")
case <-intChannels[1]:
  fmt.Println("The second candidate case is selected.")
case elem := <-intChannels[2]:   //赋值给变量
  fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
default:
  fmt.Println("No candidate case is selected!")
}
```

在使用select语句的时候，我们首先需要注意下面几个事情。

1. 如果像上述示例那样加入了默认分支，那么无论涉及通道操作的表达式是否有阻塞，select语句都不会被阻塞。如果那几个表达式都阻塞了，或者说都没有满足求值的条件，那么默认分支就会被选中并执行。
2. 如果没有加入默认分支，那么一旦所有的case表达式都没有满足求值条件，那么select语句就会被阻塞。直到至少有一个case表达式满足条件为止。
3. 我们可能会因为通道关闭了，而直接从通道接收到一个其元素类型的零值。所以，在很多时候，我们需要通过接收表达式的第二个结果值来判断通道是否已经关闭。一旦发现某个通道关闭了，我们就应该及时地屏蔽掉对应的分支或者采取其他措施。这对于程序逻辑和程序性能都是有好处的。
4. select语句只能对其中的每一个case表达式各求值一次。所以，如果我们想连续或定时地操作其中的通道的话，就往往需要通过在for语句中嵌入select语句的方式实现。但这时要注意，**简单地在select语句的分支中使用break语句，只能结束当前的select语句的执行，而并不会对外层的for语句产生作用。**这种错误的用法可能会让这个for语句无休止地运行下去。

```go
intChan := make(chan int, 1)
// 一秒后关闭通道。
time.AfterFunc(time.Second, func() {
  close(intChan)
})
select {
case _, ok := <-intChan:
  if !ok {
    fmt.Println("The candidate case is closed.")
    break
  }
  fmt.Println("The candidate case is selected.")
}
```



### select语句的分支选择规则都有哪些？

1. 对于每一个case表达式，都至少会包含一个代表发送操作的发送表达式或者一个代表接收操作的接收表达式，同时也可能会包含其他的表达式。比如，如果case表达式是包含了接收表达式的短变量声明时，那么在赋值符号左边的就可以是一个或两个表达式，不过此处的表达式的结果必须是可以被赋值的。当这样的case表达式被求值时，它包含的多个表达式总会以从左到右的顺序被求值。
2. select语句包含的候选分支中的case表达式都会在该语句执行开始时先被求值，并且求值的顺序是依从代码编写的顺序从上到下的。结合上一条规则，在select语句开始执行时，排在最上边的候选分支中最左边的表达式会最先被求值，然后是它右边的表达式。仅当最上边的候选分支中的所有表达式都被求值完毕后，从上边数第二个候选分支中的表达式才会被求值，顺序同样是从左到右，然后是第三个候选分支、第四个候选分支，以此类推。
3. 对于每一个case表达式，如果其中的发送表达式或者接收表达式在被求值时，相应的操作正处于阻塞状态，那么对该case表达式的求值就是不成功的。在这种情况下，我们可以说，这个case表达式所在的候选分支是不满足选择条件的。
4. 仅当select语句中的所有case表达式都被求值完毕后，它才会开始选择候选分支。这时候，它只会挑选满足选择条件的候选分支执行。如果所有的候选分支都不满足选择条件，那么默认分支就会被执行。如果这时没有默认分支，那么select语句就会立即进入阻塞状态，直到至少有一个候选分支满足选择条件为止。一旦有一个候选分支满足选择条件，select语句（或者说它所在的 goroutine）就会被唤醒，这个候选分支就会被执行。
5. 如果select语句发现同时有多个候选分支满足选择条件，那么它就会用一种伪随机的算法在这些分支中选择一个并执行。注意，即使select语句是在被唤醒时发现的这种情况，也会这样做。
6. 一条select语句中只能够有一个默认分支。并且，默认分支只在无候选分支可选时才会被执行，这与它的编写位置无关。
7. select语句的每次执行，包括case表达式求值和分支选择，都是独立的。不过，至于它的执行是否是并发安全的，就要看其中的case表达式以及分支中，是否包含并发不安全的代码了。

### 如果在select语句中发现某个通道已关闭，那么应该怎样屏蔽掉它所在的分支？

很简单，把nil赋给代表了这个通道的变量就可以了。如此一来，对于这个通道（那个变量）的发送操作和接收操作就会永远被阻塞。

### 在select语句与for语句联用时，怎样直接退出外层的for语句？

这一般会用到goto语句和标签（label）



## select多路复用

在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。你也许会写出如下代码使用遍历的方式来实现：

```go
for{
    // 尝试从ch1接收值
    data, ok := <-ch1
    // 尝试从ch2接收值
    data, ok := <-ch2
    …
}
```

这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多。为了应对这种场景，Go内置了`select`关键字，可以同时响应多个通道的操作。

`select`的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。`select`会一直等待，直到某个`case`的通信操作完成时，就会执行`case`分支对应的语句。具体格式如下：

```go
select{
    case <-ch1:
        ...
    case data := <-ch2:
        ...
    case ch3<-data:
        ...
    default:
        默认操作
}
```

举个小例子来演示下`select`的使用：

```go
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
```

使用`select`语句能提高代码的可读性。

- 可处理一个或多个channel的发送/接收操作。
- 如果多个`case`同时满足，`select`会随机选择一个。
- 对于没有`case`的`select{}`会一直等待，可用于阻塞main函数。

## 并发安全和锁

举个例子：

```go
var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
```

上面的代码中我们开启了两个`goroutine`去累加变量x的值，这两个`goroutine`在访问和修改`x`变量的时候就会存在数据竞争，导致最后的结果与期待的不符。

> 小技巧：使用 go build、go run、go test 这些 Go 语言工具链提供的命令时，添加 -race 标识可以帮你检查 Go 语言代码是否存在资源竞争。

### 互斥锁

channel 为什么是并发安全的呢？是因为 channel 内部使用了互斥锁来保证并发的安全

互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个`goroutine`可以访问共享资源。Go语言中使用`sync`包的`Mutex`类型来实现互斥锁。 使用互斥锁来修复上面代码的问题：

```go
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
```

使用互斥锁能够保证同一时间有且只有一个`goroutine`进入临界区，其他的`goroutine`则在等待锁；当互斥锁释放后，等待的`goroutine`才可以获取锁进入临界区，多个`goroutine`同时等待一个锁时，**唤醒的策略是随机的**。

### 读写互斥锁
互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用`sync`包中的`RWMutex`类型。
读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的`goroutine`如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个`goroutine`获取写锁之后，其他的`goroutine`无论是**获取读锁还是写锁都会等待**。
读写锁示例：

```go
var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁 记住是Rlock
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
```

需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。

### sync.WaitGroup

在代码中生硬的使用`time.Sleep`肯定是不合适的，Go语言中可以使用`sync.WaitGroup`来实现并发任务的同步。 `sync.WaitGroup`有以下几个方法：

| 方法名                          | 功能                |
| ------------------------------- | ------------------- |
| (wg * WaitGroup) Add(delta int) | 计数器+delta        |
| (wg *WaitGroup) Done()          | 计数器-1            |
| (wg *WaitGroup) Wait()          | 阻塞直到计数器变为0 |

`sync.WaitGroup`内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。

我们利用`sync.WaitGroup`将上面的代码优化一下：

```go
var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}
func main() {
	wg.Add(1)
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wg.Wait()
}
```

需要注意`sync.WaitGroup`是一个**结构体**，传递的时候要**传递指针**。

### sync.Once

说在前面的话：这是一个进阶知识点。

在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

Go语言中的`sync`包中提供了一个针对只执行一次场景的解决方案–`sync.Once`。

`sync.Once`只有一个`Do`方法，其签名如下：

```go
func (o *Once) Do(f func()) {}
```

*备注：如果要执行的函数`f`需要传递参数就需要搭配**闭包**来使用。*
#### 加载配置文件示例

延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个很好的实践。因为预先初始化一个变量（比如在init函数中完成初始化）会增加程序的启动耗时，而且有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的。我们来看一个例子：

```go
var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon 被多个goroutine调用时不是并发安全的
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}
```

多个`goroutine`并发调用Icon函数时不是并发安全的，现代的编译器和CPU可能会在保证每个`goroutine`都满足串行一致的基础上自由地重排访问内存的顺序。loadIcons函数可能会被重排为以下结果：

```go
func loadIcons() {
	icons = make(map[string]image.Image)
	icons["left"] = loadIcon("left.png")
	icons["up"] = loadIcon("up.png")
	icons["right"] = loadIcon("right.png")
	icons["down"] = loadIcon("down.png")
}
```

在这种情况下就会出现即使判断了`icons`不是nil也不意味着变量初始化完成了。考虑到这种情况，我们能想到的办法就是添加互斥锁，保证初始化`icons`的时候不会被其他的`goroutine`操作，但是这样做又会引发性能问题。
使用`sync.Once`改造的示例代码如下：

```go
var icons map[string]image.Image

var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon 是并发安全的
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
```

#### 并发安全的单例模式

下面是借助`sync.Once`实现的并发安全的单例模式：

```go
package singleton

import (
    "sync"
)

type singleton struct {}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```

`sync.Once`其实内部包含一个**互斥锁和一个布尔值**，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。

## 原子操作

代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。针对**基本数据类型**我们还可以使用**原子操作**来保证并发安全，因为原子操作是Go语言提供的方法它在**用户态**就可以完成，因此性能比加锁操作更好。Go语言中原子操作由内置的标准库`sync/atomic`提供。

### atomic包

| 方法                                                         | 解释           |
| ------------------------------------------------------------ | -------------- |
| func LoadInt32(addr *int32) (val int32)                                                                                                         func LoadInt64(addr *int64) (val int64)                                                                                                        func LoadUint32(addr *uint32) (val uint32)                                                                                               func LoadUint64(addr *uint64) (val uint64)                                                                                             func LoadUintptr(addr *uintptr) (val uintptr)                                                                                                 func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer) | 读取操作       |
| func StoreInt32(addr *int32, val int32)                                                                                                                func StoreInt64(addr *int64, val int64)                                                                                                             func StoreUint32(addr *uint32, val uint32)                                                                                                                  func StoreUint64(addr *uint64, val uint64)                                                                                                                  func StoreUintptr(addr *uintptr, val uintptr)                                                                                                            func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer) | 写入操作       |
| func AddInt32(addr *int32, delta int32) (new int32)                                                                                 func AddInt64(addr *int64, delta int64) (new int64)                                                                                 func AddUint32(addr *uint32, delta uint32) (new uint32)                                                                         func AddUint64(addr *uint64, delta uint64) (new uint64)                                                                    func AddUintptr(addr *uintptr, delta uintptr) (new uintptr) | 修改操作       |
| func SwapInt32(addr *int32, new int32) (old int32)                                                                                  func SwapInt64(addr *int64, new int64) (old int64)                                                                                  func SwapUint32(addr *uint32, new uint32) (old uint32)                                                                       func SwapUint64(addr *uint64, new uint64) (old uint64)                                                                        func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)                                                                   func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer) | 交换操作       |
| func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)                                    func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)                                   func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)                                 func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)                          func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)                      func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool) | 比较并交换操作 |

### 示例

我们填写一个示例来比较下互斥锁和原子操作的性能。

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter interface {
	Inc()
	Load() int64
}

// 普通版
type CommonCounter struct {
	counter int64
}

func (c CommonCounter) Inc() {
	c.counter++
}

func (c CommonCounter) Load() int64 {
	return c.counter
}

// 互斥锁版
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

// 原子操作版
type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main() {
	c1 := CommonCounter{} // 非并发安全
	test(c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全
	test(&c2)
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
	test(&c3)
}
```

`atomic`包提供了底层的原子级内存操作，对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。






