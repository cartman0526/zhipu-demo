package main

import (
	"fmt"
	"io"
	"net/http"
	"zhipu-demo/pkg/zhipu"
)

func main() {
	token, err := zhipu.GenerateToken("", 10)
	if err != nil {
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://open.bigmodel.cn/api/paas/v4/chat/completions", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(body))
}
