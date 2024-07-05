# gorestclient

[![Go Reference](https://pkg.go.dev/badge/github.com/aafak/gorestclient.svg)](https://pkg.go.dev/github.com/aafak/gorestclient)

A Go client library for interacting with RESTful APIs.

## Features

- Easy to use
- Supports JSON

## Installation

```sh
go get -u github.com/aafak/gorestclient
```

# Usage:

```
package main

import (
	"fmt"
	"net/http"

	"github.com/aafak/gorestclient"
)

func main() {
	restClient := gorestclient.NewRestClient()
	req, err := restClient.BuildRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Printf("Failed to build request, Err : %v\n", err)
	}
	resp, err := restClient.ExecuteRequest(req)
	if err != nil {
		fmt.Printf("Failed to send request, Err : %v, resp: %v, \n", err, resp)
	}
	defer resp.Body.Close()
	fmt.Println(resp)


POST request:
	req, _ := restClient.BuildRequest("POST", "https://jsonplaceholder.typicode.com/posts", `{"title": "foo","body": "bar","userId": 1}`)
	resp, err := restClient.ExecuteRequest(req)
	fmt.Println(resp)
```