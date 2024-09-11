package main

import (
	"fmt"
	"os"
)

func main() {
	fprintfDemo()
}

func createDemo() {
	file, err := os.Create("demo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
}

func fprintfDemo() {
	n1, err := fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n1)
	fileObj, err := os.OpenFile("./demo.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "words"

	n2, err := fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n2)
}
