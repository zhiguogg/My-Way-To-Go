package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("hello goroutine")
}

func helloGoroutine(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func helloGoroutineDemo() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go helloGoroutine(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func goMaxPROCSDemo() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

func chanDemo() {
	ch := make(chan int, 10)
	ch <- 10
	x := <-ch
	fmt.Println(x)
	close(ch)
}

func recover(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func noBufferChan() {
	ch := make(chan int)
	go recover(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

func chanIsClose()  {
	ch1 := make(chan int )
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i:=0;i<100;i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for  {
			i,ok := <- ch1  // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	for i := range ch2{ // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

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

var x int64
func add() {
	for i := 0;i<100;i++ {
		x = x+1
	}
	wg.Done()
}

func sumDemo()  {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

var lock sync.Mutex

func addMutex()  {
	for i := 0;i<100;i++ {
		//加锁
		lock.Lock()
		x = x+1
		lock.Unlock()
	}
	wg.Done()
}

func sumMutexDemo()  {
	wg.Add(3)
	go addMutex()
	go addMutex()
	go addMutex()
	wg.Wait()
	fmt.Println(x)
}


var rwLock sync.RWMutex

func write()  {
	// 加写锁
	rwLock.Lock()
	//lock.Lock()
	x = x +1
	time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
	rwLock.Unlock()
	//lock.Unlock()
	wg.Done()
}

func read()  {
	// 加读锁
	//lock.Lock()
	rwLock.RLock()
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwLock.RUnlock()
	//lock.Unlock()
	wg.Done()
}

func readAndWrite()  {
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

func helloWaitGroup()  {
	defer wg.Done()
	fmt.Println("hello goroutine")
}

func helloWaitGroupTest()  {
	wg.Add(2)
	go helloWaitGroup()
	go helloWaitGroup()
	fmt.Println("done!")
	wg.Wait()

}

func race(){

	cond :=sync.NewCond(&sync.Mutex{})

	var wg sync.WaitGroup

	wg.Add(11)

	for i:=0;i<10; i++ {

		go func(num int) {

			defer  wg.Done()

			fmt.Println(num,"号已经就位")

			cond.L.Lock()

			cond.Wait()//等待发令枪响

			fmt.Println(num,"号开始跑……")

			cond.L.Unlock()

		}(i)

	}

	//等待所有goroutine都进入wait状态

	time.Sleep(2*time.Second)

	go func() {

		defer  wg.Done()

		fmt.Println("裁判已经就位，准备发令枪")

		fmt.Println("比赛开始，大家准备跑")

		cond.Broadcast()//发令枪响

	}()

	//防止函数提前返回退出

	wg.Wait()

}

func main() {
	//go hello()
	//fmt.Println("main goroutine done!")
	//time.Sleep(time.Second)

	//helloGoroutineDemo()

	//goMaxPROCSDemo()

	//chanDemo()
	//noBufferChan()
	//chanIsClose()

	//sumMutexDemo()

	//readAndWrite()

	//helloWaitGroupTest()
	race()
}
