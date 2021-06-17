package provider

import (
	"encoding/json"
	"fmt"
	"log"

	"mini/config"
	"mini/req"
	"net/http"
)

type Auth interface {
	Code2Session(jsCode string) (*AuthResponse, error)
	GetPaidUnionId(openid string) (*AuthResponse, error)
}

type AuthResponse struct {
	Openid     string `json:"openid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	Unionid    string `json:"unionid,omitempty"`
	ErrorResponse
}

type auth struct {
	conf   *config.Config
	client req.HttpClient
}

func NewAuth(c *config.Config, client req.HttpClient) Auth {
	return &auth{
		conf:   c,
		client: client,
	}
}

func (a *auth) Code2Session(jsCode string) (*AuthResponse, error) {
	resp, err := a.client.Get(fmt.Sprintf(sessionURL, a.conf.AppId, a.conf.AppSecret, jsCode), nil)
	if err != nil {
		log.Println("Code2Session get error: ", err)
		return nil, err
	}
	if resp.Code != http.StatusOK {
		log.Printf("Code2Session request incorrect,statusCode: %d", resp.Code)
		return nil, fmt.Errorf("bad request")
	}
	var ret AuthResponse
	err = json.Unmarshal(resp.Body, &ret)
	if err != nil {
		log.Printf("Code2Session unmarshal error: %v", err)
		return nil, err
	}
	return &ret, nil
}

func (a *auth) GetPaidUnionId(openid string) (*AuthResponse, error) {
	token, err := AccessTokenHandler(a.client, a.conf)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(paidUnionIdURL, token, openid)
	fmt.Println("paid url: ", url)
	resp, err := a.client.Get(url, nil)
	if err != nil {
		log.Println("GetPaidUnionId get error: ", err)
		return nil, err
	}
	if resp.Code != http.StatusOK {
		log.Printf("GetPaidUnionId request incorrect,statusCode: %d", resp.Code)
		return nil, fmt.Errorf("bad request")
	}
	var ret AuthResponse
	err = json.Unmarshal(resp.Body, &ret)
	if err != nil {
		log.Printf("GetPaidUnionId unmarshal error: %v", err)
		return nil, err
	}
	return &ret, nil
}
