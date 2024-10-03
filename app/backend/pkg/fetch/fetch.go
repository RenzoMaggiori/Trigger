package fetch

import (
	"io"
	"log"
	"net/http"
	"time"
)

type FetchRequest struct {
	Method  string
	Url     string
	Body    io.Reader
	Headers map[string]string
}

func NewFetchRequest(method string, url string, body io.Reader, headers map[string]string) *FetchRequest {
	return &FetchRequest{
		Method:  method,
		Url:     url,
		Body:    body,
		Headers: headers,
	}
}

func Fetch(client *http.Client, request *FetchRequest) (*http.Response, error) {
	req, err := http.NewRequest(request.Method, request.Url, request.Body)
	if err != nil {
		return nil, err
	}

	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}

	start := time.Now()
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Fetching %s %s [%d] %s", req.Method, req.URL.Path, res.StatusCode, time.Since(start))
	return res, nil
}
