package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	demo4()
}

func demo1() {
	file, err := os.Open("../log_demo/demo.log")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

func demo2() {
	file, err := os.Open("../log_demo/demo.log")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	scanner := bufio.NewScanner(file) // MaxScanTokenSize = 64 * 1024
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("read file failed, err:", err)
	}
}

func demo3() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	str := "str1"
	write, err := file.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(write)
	writeString, err := file.WriteString("str2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(writeString)
}

func demo4() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writeString, err := writer.WriteString("test" + strconv.Itoa(i) + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(writeString)
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

}
