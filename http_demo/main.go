package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	for {
		time.Sleep(time.Millisecond * 100)
		go func() {
			demo()
		}()
	}
}

func demo() {
	resp, err := http.Get("https://feel.wechen.xyz/index.php/Wechat/Customer/detail?id=3")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
