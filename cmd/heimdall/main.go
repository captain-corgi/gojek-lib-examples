package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

func main() {
	// Create a new HTTP client with a default timeout
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Use the clients GET method to create and execute the request
	res, err := client.Get("http://google.com", nil)
	if err != nil {
		panic(err)
	}

	// Heimdall returns the standard *http.Response object
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

}
