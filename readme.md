Rest client utility to make the rest call

Usage:
```
    restClient := NewRestClient()
	req, _ := restClient.BuilRequest("GET", "https://jsonplaceholder.typicode.com/posts")
	resp, err := restClient.ExecuteRequest(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)


POST request:
    restClient := NewRestClient()
	req, _ := restClient.BuilRequest("POST", "https://jsonplaceholder.typicode.com/posts", `{"title": "foo","body": "bar","userId": 1}`)
	resp, err := restClient.ExecuteRequest(req)
	fmt.Println(resp)
```