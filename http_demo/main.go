package main

import (
	"fmt"
	"go-web-advance/utils"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.bing.com")
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
	fmt.Println(utils.ValueJsonString(body))
}
