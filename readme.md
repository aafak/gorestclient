Rest client utility to make the rest call

Usage:
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