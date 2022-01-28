package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type httpbin struct {
	Args struct {
	} `json:"args"`
	Data  string `json:"data"`
	Files struct {
	} `json:"files"`
	Form struct {
	} `json:"form"`
	Headers struct {
		AcceptEncoding string `json:"Accept-Encoding"`
		ContentLength  string `json:"Content-Length"`
		Host           string `json:"Host"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceID   string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	JSON   interface{} `json:"json"`
	Method string      `json:"method"`
	Origin string      `json:"origin"`
	URL    string      `json:"url"`
}

var httpBinChan = make(chan httpbin)

func main() {

	httpClient := http.DefaultClient

	go func() {
		rb := httpbin{}
		reqBody := strings.NewReader("\"location\":\"india\"")
		req, _ := http.NewRequest(http.MethodGet, "https://httpbin.org/anything/{anything}", reqBody)
		resp, _ := httpClient.Do(req)
		defer resp.Body.Close()
		dataBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(dataBytes, &rb)
		httpBinChan <- rb
	}()

	go func() {
		rb := httpbin{}
		reqBody := strings.NewReader("\"location\":\"singapore\"")
		req, _ := http.NewRequest(http.MethodGet, "https://httpbin.org/anything/{anything}", reqBody)
		resp, _ := httpClient.Do(req)
		defer resp.Body.Close()
		dataBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(dataBytes, &rb)
		httpBinChan <- rb
	}()

	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case data := <-httpBinChan:
			fmt.Println(data.Data)
			return
		default:
			fmt.Println("no value received")
		}
	}
}
