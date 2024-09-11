package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sscanDemo()
}

func demo1() {
	var (
		name    string
		age     int
		married bool
	)
	scan, err := fmt.Scan(&name, &age, &married)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(scan)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

func demo2() {
	var (
		name    string
		age     int
		married bool
	)
	scanf, err := fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(scanf)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

func fscanDemo() {
	file, err := os.Open("../fmt_demo/demo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var s string
	fscan, err := fmt.Fscan(file, &s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fscan)
	fmt.Printf("%#v\n", s)
}

func sscanDemo() {
	s := "sfe jfjsfsafas OK:1"
	var s2 string
	var n int
	sscanf, err := fmt.Sscanf(s, "sfe %s OK:%d", &s2, &n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sscanf)
	fmt.Printf("%#v %#v\n", s2, n)
}
