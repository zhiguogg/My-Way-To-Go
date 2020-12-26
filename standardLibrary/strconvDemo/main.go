package main

import (
	"fmt"
	"strconv"
)

const (
	i = iota
)

func atoiDemo()  {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}
}
func itoaDemo()  {
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"
}

func parseDemo()  {
	//a,err := strconv.ParseBool("0")
	//if err != nil {
	//	fmt.Println("conversion failed")
	//}
	//fmt.Printf("type is %T,value is %v\n",a,a)

	//b,err := strconv.ParseInt("-123",0,0)
	//if err != nil {
	//	fmt.Println("conversion failed")
	//}
	//fmt.Printf("type is %T,value is %v\n",b,b)

	//c,err := strconv.ParseUint("123",0,0)
	//if err != nil {
	//	fmt.Println("conversion failed")
	//}
	//fmt.Printf("type is %T,value is %v\n",c,c)


	d, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println("conversion failed")
	}
	fmt.Printf("type is %T,value is %v\n",d,d)


}

func formatDemo()  {
	a := strconv.FormatBool(true)
	fmt.Println(a)

	b := strconv.FormatInt(-5,2)
	fmt.Println(b)

	c := strconv.FormatUint(5,2)
	fmt.Println(c)

	d := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(d)
}

func main()  {
	//atoiDemo()
	//itoaDemo()
	//parseDemo()
	formatDemo()
}