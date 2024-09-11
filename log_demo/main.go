package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	demo4()
}

func demo1() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}

func demo2() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[hello]")
	log.Println("这是一条很普通的日志。")

	log.SetPrefix("[no]")
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("")
	log.Println("这是一条很普通的日志。")

}
func demo3() {

	logFile, err := os.OpenFile("./demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[test]")
	log.Println("这是一条很普通的日志。")

}

func demo4() {
	logFile, err := os.OpenFile("./demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	demo4Logger := log.New(logFile, "demo4", log.Llongfile|log.Lmicroseconds|log.Ldate)
	demo4Logger.Println("hello")
	demo4Logger.Println("what")
}
