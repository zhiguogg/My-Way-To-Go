package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func openAndCloseFile()  {
	// 只读方式打开当前目录下的main.go文件
	file,err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	file.Close()
}

func readDemo()  {
	// 只读方式打开当前目录下的main.go文件
	file,err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	defer file.Close()
	//使用Read方法读取数据
	var tem = make([]byte,128)
	read, err := file.Read(tem)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", read)
	fmt.Println(string(tem[:read]))
}

func readFileDemo()  {
	file,err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()

	//循环读取文件
	var content []byte
	var tem = make([]byte,128)
	for  {
		read, err := file.Read(tem)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break  // 跳出循环
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tem[:read]...)
	}
	fmt.Println(string(content))
}

// bufio按行读取示例
func bufioDemo()  {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for  {
		line,err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0  {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

func iotilDemo()  {
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func writeAndWriteString()  {
	file, err := os.OpenFile("fileWrite.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "我不是归人\n"
	file.Write([]byte(str)) //写入字节切片数据
	file.WriteString("就算会错过什么") //直接写入字符串数据
}

func bufioNewWrite()  {
	file, err := os.OpenFile("fileWrite.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("我不是归人\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

func iotilWriteDemo()  {
	str := "对我而言是另一天"
	err := ioutil.WriteFile("./fileWrite.txt",[]byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main()  {
	//readDemo()
	//readFileDemo()
	//bufioDemo()
	//iotilDemo()
	//writeAndWriteString()
	//bufioNewWrite()
	iotilWriteDemo()
	
}
