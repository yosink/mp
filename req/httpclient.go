package req

import (
	"github.com/go-resty/resty/v2"
)

const (
	ApplicationJson      = "application/json"
	ApplicationForm      = "application/x-www-form-urlencoded"
	ApplicationMultipart = "multipart/form-data"
)

type Response struct {
	Code int
	Body []byte
}

type HttpClient interface {
	Get(url string, params map[string]string) (*Response, error)
	PostForm(url string, data map[string]interface{}) (*Response, error)
	PostJson(url string, data interface{}) (*Response, error)
}

func NewRequest() HttpClient {
	return getDefaultRequest()
}

func getDefaultRequest() HttpClient {
	return NewResty()
}

type restyClient struct {
	client *resty.Request
}

func NewResty() HttpClient {
	return &restyClient{
		client: resty.New().R(),
	}
}

func (r *restyClient) Get(url string, params map[string]string) (*Response, error) {
	resp, err := r.client.SetQueryParams(params).Get(url)
	if err != nil {
		return nil, err
	}
	return &Response{
		Code: resp.StatusCode(),
		Body: resp.Body(),
	}, nil
}
func (r *restyClient) PostForm(url string, data map[string]interface{}) (*Response, error) {
	resp, err := r.client.SetHeader("Content-Type", ApplicationForm).
		SetBody(data).Get(url)
	if err != nil {
		return nil, err
	}
	return &Response{
		Code: resp.StatusCode(),
		Body: resp.Body(),
	}, nil
}
func (r *restyClient) PostJson(url string, data interface{}) (*Response, error) {
	resp, err := r.client.SetHeader("Content-Type", ApplicationJson).
		SetBody(data).Post(url)
	if err != nil {
		return nil, err
	}
	return &Response{
		Code: resp.StatusCode(),
		Body: resp.Body(),
	}, nil
}
