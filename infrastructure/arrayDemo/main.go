package main

import "fmt"

func declareArray() {
	var testArray1 [3]int                //数组会初始化为int类型的零值
	var testArray2 [3]int = [3]int{7, 4} // 可以省略 前面的[3]int
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray1, testArray2, cityArray)

	a := [...]int{0: 2, 2: 4, 5: 1}
	fmt.Println(a)
}

func traverse() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func changeArray() {
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}

func twoDimensionalArray() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}

func traverseArray2() {
	a := [2][3]string{
		{"安徽", "合肥", "大蜀山"},
		{"江苏", "南京", "雨花台"},
	}
	for _, value := range a {
		for _, value := range value { //注意这两个value的作用域，最好起不一样的名字
			fmt.Println(value)
		}
	}
}

func main() {
	//declareArray()
	//traverse()
	//changeArray()
	//traverseArray2()
}
