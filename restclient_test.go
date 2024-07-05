package gorestclient

import (
	"fmt"
	"testing"
)

func TestNewRestClient(t *testing.T) {
	restClient := NewRestClient()
	if restClient == nil {
		t.Errorf("NewRestClient() = %v; want RestClientInterface", restClient)
	}
}

func TestNewRestClientWithMalformCerts(t *testing.T) {
	restClient, err := NewRestClientWithCerts("malformcerts/malcert.pem", "malformcerts/malkey.pem")
	if err == nil {
		t.Errorf("NewRestClient() = %v; want Err", restClient)
	}
}

func TestNewRestClientWithCerts(t *testing.T) {
	restClient, _ := NewRestClientWithCerts("malformcerts/cert.pem", "malformcerts/privkey.pem")
	if restClient == nil {
		t.Errorf("NewRestClient() = %v; want RestClientInterface", restClient)
	}
}

func TestExecuteReq(t *testing.T) {
	restClient := NewRestClient()
	req, _ := restClient.BuildRequest("GET", "https://jsonplaceholder.typicode.com/posts")
	resp, err := restClient.ExecuteRequest(req)
	if err != nil {
		t.Errorf("ExecuteRequest() error = %v; want nil", err)
	}
	if resp == nil {
		t.Errorf("ExecuteRequest() = %v; want *http.Response", resp)
	}
}

func TestExecutePostReq(t *testing.T) {
	restClient := NewRestClient()
	req, _ := restClient.BuildRequest("POST", "https://jsonplaceholder.typicode.com/posts", `{"title": "foo","body": "bar","userId": 1}`)
	req.Header.Add("Content-Type", "application/json")
	resp, err := restClient.ExecuteRequest(req)
	fmt.Println(resp)
	if err != nil {
		t.Errorf("ExecuteRequest() error = %v; want nil", err)
	}
	if resp == nil {
		t.Errorf("ExecuteRequest() = %v; want *http.Response", resp)
	}
}
