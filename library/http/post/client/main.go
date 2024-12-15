package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// net/http post demo

func main() {
	// POST获取响应
	apiurl := "http://127.0.0.1:9090/post"
	data := `{"name": "枯藤 ", "age": 18}`
	contentType := "application/json"
	result, err := HTTPPost(apiurl, contentType, data)
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
	}
	fmt.Println("post results: %s", string(result))

	// HTTPPostForm 示例调用
	formURL := "http://127.0.0.1:9090/post_form"
	formData := url.Values{
		"name": {"枯藤"},
		"age":  {"18"},
	}
	result, err = HTTPPostForm(formURL, formData)
	if err != nil {
		fmt.Printf("post form failed, err:%v\n", err)
		return
	}
	fmt.Printf("post form results: %s\n", string(result))
}

// HTTPPost POST获取响应
func HTTPPost(u, contentType, body string) ([]byte, error) {

	resp, err := autoClient(u).Post(u, contentType, strings.NewReader(body))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)

		return nil, errors.New("http post err")
	}
	defer resp.Body.Close()
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("read response body err")
	}
	return p, nil
}

// HTTPPostForm POST表单并获取响应
func HTTPPostForm(u string, data url.Values) ([]byte, error) {
	c := autoClient(u)

	resp, e := c.PostForm(u, data)
	if e != nil {
		log.Printf("http post form err: %s", e)

		return nil, errors.New("http post form err")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		e = fmt.Errorf("post{%s} resp status{%d, %s}", u, resp.StatusCode, resp.Status)
		return nil, e
	}

	p, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Printf("read all content from response body err: %s", e)

		return nil, errors.New("read response body err")
	}

	return p, nil
}

// autoClient 根据url判断使用HTTPS还是HTTP
func autoClient(u string) *http.Client {
	scheme := strings.Split(u, "://")[0]

	if strings.EqualFold(scheme, "https") {
		return &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	}

	return http.DefaultClient
}
