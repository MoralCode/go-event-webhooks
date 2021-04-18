package helpers

import "net/http"
import (
	"bytes"
	"errors"
	"io/ioutil"
)

type MockClient struct {
    DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
    if m.DoFunc != nil {
        return m.DoFunc(req)
    }
    return &http.Response{}, nil
}


func SetupSuccessResponse(client *MockClient, response_str string) {
	client.DoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(response_str)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
}


func SetupMirrorResponse(client *MockClient) {
	client.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       req.Body,
		}, nil
	}
}


func SetupServerError(client *MockClient) {
	client.DoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New(
			"Error from web server",
		)
	}
}