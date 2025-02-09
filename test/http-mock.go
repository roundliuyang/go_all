package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	//创建一个模拟的服务器
	server := mockServer()
	defer server.Close()
	//Get请求发往模拟服务器的地址
	resq, err := http.Get(server.URL)
	if err != nil {
		log.Fatal("创建Get失败")
	}
	defer resq.Body.Close()

	log.Println("code:", resq.StatusCode)
	data, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("body:%s\n", data)
}

func mockServer() *httptest.Server {
	//API调用处理函数
	sendJson := func(rw http.ResponseWriter, r *http.Request) {
		u := struct {
			Name string
		}{
			Name: "张三",
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		err := json.NewEncoder(rw).Encode(u)
		if err != nil {
			log.Println("JSON 编码失败:", err)
		}
	}
	//适配器转换
	return httptest.NewServer(http.HandlerFunc(sendJson))
}
