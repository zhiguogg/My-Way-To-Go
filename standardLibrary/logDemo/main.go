package main

import (
	"fmt"
	"log"
	"os"
)

func logDemo()  {
	logFile,err := os.OpenFile("./logDemo.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("简单使用log")
	v := "我不是归人"
	log.SetPrefix("golang:")
	log.Printf("这是%s的日志\n",v)
	//log.Fatalln("这是一条会触发fatal的日志")
	//log.Panicln("这是一条会触发panic的日志。")
}

func logDemo1()  {
	logger := log.New(os.Stdout, "<我不是归人>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}

func main()  {

	logDemo1()
}