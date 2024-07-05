package gorestclient

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

type RestClientInterface interface {
	ExecuteRequest(req *http.Request) (*http.Response, error)
	BuildRequest(method, url string, body ...string) (*http.Request, error)
}

type RestClient struct {
	Client *http.Client
}

func NewRestClient() RestClientInterface {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
	}
	client := &http.Client{Transport: tr}
	return &RestClient{
		Client: client,
	}
}

func NewRestClientWithCerts(certFilePath, privateKeyFilePath string) (RestClientInterface, error) {
	pubCert, err := os.ReadFile(certFilePath)
	if err != nil {
		fmt.Printf("reading cert file failed with error: %s", err.Error())
		return nil, err
	}

	privKey, err := os.ReadFile(privateKeyFilePath)
	if err != nil {
		fmt.Printf("reading private key file failed with error: %s", err.Error())
		return nil, err
	}

	cert, err := tls.X509KeyPair(pubCert, privKey)
	if err != nil {
		fmt.Printf("falied to load certificate pair error: %s", err.Error())
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
		},
	}

	client := &http.Client{Transport: tr}
	// client := &http.Client{Transport: tr, Timeout: timeout}
	restClient := &RestClient{
		Client: client,
	}
	return restClient, nil
}

func (rc *RestClient) ExecuteRequest(req *http.Request) (*http.Response, error) {
	fmt.Printf("Executing %s request %s\n", req.Method, req.URL)
	resp, err := rc.Client.Do(req)
	if err != nil {
		fmt.Printf("%s request %s failed, response:%v\n", req.Method, req.URL, resp)
		return nil, err
	}
	return resp, nil
}

func (r *RestClient) BuildRequest(method, url string, body ...string) (*http.Request, error) {

	if len(body) > 0 && body[0] != "" {
		reqbodyPayload := bytes.NewBuffer([]byte(body[0]))
		req, err := http.NewRequest(method, url, reqbodyPayload)
		if err != nil {
			return nil, err
		}
		return req, nil
	} else {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
		return req, nil
	}
}

// func main() {
// 	restClient := NewRestClient()
// 	req, _ := restClient.BuilRequest("GET", "https://jsonplaceholder.typicode.com/posts")
// 	resp, err := restClient.ExecuteRequest(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	fmt.Println(resp.Status)
// 	fmt.Println(resp.StatusCode)

// }
