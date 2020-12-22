package main

import "fmt"

func ifDemo() {
	if i := 6; i > 10 {

		fmt.Println("i>10")

	} else if i > 5 && i <= 10 {

		fmt.Println("5<i<=10")

	} else {

		fmt.Println("i<=5")

	}
}

func switchDemo() {
	switch i := 6; {

	case i > 10:

		fmt.Println("i>10")

	case i > 5 && i <= 10:

		fmt.Println("5<i<=10")

	default:

		fmt.Println("i<=5")

	}
}

func testSwitch3() {
	switch n := 7; n { //和if之前加一个执行语句多像 不能有var
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func switchDemo4() {
	age := 30
	switch { // 不能有表达式
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

func forDetails1() {
	numbers1 := [...]int{1, 2, 3, 4, 5, 6}
	for i := 10; i < len(numbers1); i++ {
		if i == 3 {
			numbers1[i] += 10
		}
	}
	fmt.Println(numbers1)
}

func forDetails2()  {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := 10; i < len(numbers1); i++ {
		if i == 3 {
			numbers1[i] += 10
		}
	}
	fmt.Println(numbers1)
}

func rangeDetails1() {
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
}

func rangeDetails2()  {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] += 10
		}
	}
	fmt.Println(numbers1)
}

func rangeDetails3()  {
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
}


func forAndChange() []int {
	a := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		a[i] = a[i] + 1
	}
	return a
}

func forRangeAndChange(a []int) []int {
	for k, v := range a {
		a[k] = v + 1
	}
	return a
}

func main() {
	//fmt.Println(forAndChange())
	//a := []int{1, 2, 3, 4, 5}
	//fmt.Println(forRangeAndChange(a))

	//forDetails1()
	//forDetails2()
	rangeDetails1()
	rangeDetails2()
	rangeDetails3()
}
