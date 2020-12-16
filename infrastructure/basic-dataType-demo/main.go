package main

import "fmt"

func byteDemo()  {
	var a byte = 65
	var b uint8 = 65
	fmt.Println(a)
	fmt.Printf("a的值： %c \n", a)
	fmt.Println(b)
	fmt.Printf("a的值： %c \n", b)

	var aa byte = 'A'
	var bb byte = 'B'
	fmt.Println(aa)
	fmt.Printf("a的值： %c \n", aa)
	fmt.Println(bb)
	fmt.Printf("a的值： %c \n", bb)
}

func runeDemo()  {
	a := '中'
	fmt.Println(a)
	fmt.Printf("a 的值: %c",  a)
}

func stringDemo()  {
	a := "\\r\\n"   // 因为\r \n是转义字符
	b := `\r\n`
	fmt.Println(a)
	fmt.Println(b)

	var aa string = `ssss
yyyyyyy`
	var bb string = "wwww\n yyy"
	fmt.Println(aa)
	fmt.Println(bb)
}

func changeString()  {
	//纯粹英文
	s := "zhiguogg"
	// 强制类型转换
	byteS := []byte(s)
	byteS[0] = 'c'
	//fmt.Println(string(byteS))

	//纯粹中文
	s1 := "我不是归人"
	runeS := []rune(s1)
	runeS[0] = '你'
	//fmt.Println(string(runeS))

	//混合
	s2 := "hello,中国"
	byteS2 := []byte(s2)
	//fmt.Println(byteS2)
	byteS2[0] = 'l'
	byteS2[9] = 228
	byteS2[10] = 184
	byteS2[11] = 173
	//我们把第一个英文字符改成 l，把后面代表'国'的三个字节更改成和'中'一样 预测为 lello,中中
	//fmt.Println(string(byteS2))

	s3 := "hello,中国"
	runeS3 := []rune(s3)
	fmt.Println(runeS3)
	runeS3[0] = 'l'
	runeS3[7] = '中'
	fmt.Println(string(runeS3))



}

func main()  {
	//var a int = 0b1100
	//fmt.Println(a)

	//byteDemo()
	//runeDemo()
	//stringDemo()

	changeString()
}
