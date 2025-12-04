package nansen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"golang.org/x/time/rate"
)

const DefaultURL = "https://api.nansen.ai/api/v1"
const CustomKeyHeader = "apiKey"

// TODO: find correct rate limit
const RateLimitMilliseconds = 120

type Client struct {
	ApiUrl             string
	ApiKey             string
	RetaLimitRemaining int64
	Limiter            *rate.Limiter
	c                  *http.Client
}

func NewClient(url, key string) *Client {
	return &Client{
		ApiUrl:  url,
		ApiKey:  key,
		Limiter: rate.NewLimiter(rate.Every(RateLimitMilliseconds*time.Millisecond), 1),
		c:       http.DefaultClient,
	}
}

func (c *Client) Url(endpoint string) string {
	return fmt.Sprintf("%s%s", c.ApiUrl, endpoint)
}

func (c *Client) doCall(ctx context.Context, req *Request, response any) (*http.Response, error) {
	// LIMITER
	err := c.Limiter.Wait(ctx)
	if err != nil {
		return nil, err
	}
	if reflect.TypeOf(response).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("response struct is not a pointer")
	}
	httpRequest, err := req.NewHttpRequest(ctx, c.ApiKey)
	if err != nil {
		return nil, fmt.Errorf("api call %v() on %v: %v", req.Method, req.Endpoint, err.Error())
	}
	httpResponse, err := c.c.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("api call %v() on %v: %v", req.Method, httpRequest.URL.String(), err.Error())
	}

	bodyBytes, err := io.ReadAll(httpResponse.Body)
	// parsing error
	if err != nil {
		return nil, fmt.Errorf(
			"call %v() on %v status code: %v. could not decode body to response: %v",
			req.Method,
			httpRequest.URL.String(),
			httpResponse.StatusCode,
			err.Error())
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf(
			"call %v() on %v status code: %v.raw body %v",
			req.Method,
			httpRequest.URL.String(),
			httpResponse.StatusCode,
			string(bodyBytes))
	}

	err = json.Unmarshal(bodyBytes, response)

	if err != nil {
		return nil, fmt.Errorf(
			"call %v() on %v status code: %v. could not decode body to response model: %v",
			req.Method,
			httpRequest.URL.String(),
			httpResponse.StatusCode,
			err.Error())
	}
	if response == nil {
		// if we have some http error, return it
		return nil, fmt.Errorf("call %v() on %v status code: %v. response missing",
			req.Method,
			httpRequest.URL.String(),
			httpResponse.StatusCode)
	}

	return httpResponse, nil
}

type Request struct {
	Body     any
	Endpoint string
	Method   string
}

func NewRequest(endpoint string, body any, methods ...string) *Request {
	method := http.MethodGet
	if len(methods) != 0 {
		method = methods[0]
	}
	return &Request{
		Body:     body,
		Endpoint: endpoint,
		Method:   method,
	}
}

func (r *Request) NewHttpRequest(ctx context.Context, token string) (*http.Request, error) {
	var bodyReader io.Reader
	if r.Method != http.MethodGet {
		byteBody, err := json.Marshal(r.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(byteBody)
	}
	request, err := http.NewRequest(r.Method, r.Endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(CustomKeyHeader, token)
	//Custom headers
	request = request.WithContext(ctx)

	return request, nil
}
