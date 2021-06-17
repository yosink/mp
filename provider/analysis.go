package provider

import (
	"encoding/json"
	"fmt"
	"log"
	"mini/config"
	"mini/req"
	"net/http"
)

type Analysis interface {
	GetDailyRetain(data DailyRetainRequest) (*DailyRetainResponse, error)
}

type KVResopnse struct {
	Key   int `json:"key,omitempty"`
	Value int `json:"value,omitempty"`
}

type DailyRetainRequest struct {
	BeginDate string `json:"begin_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type DailyRetainResponse struct {
	RefDate    string     `json:"ref_date,omitempty"`
	VisitUvNew KVResopnse `json:"visit_uv_new,omitempty"`
	VisitUv    KVResopnse `json:"visit_uv,omitempty"`
}

type analysis struct {
	conf   *config.Config
	client req.HttpClient
}

func NewAnalysis(c *config.Config, client req.HttpClient) Analysis {
	return &analysis{c, client}
}

func (a *analysis) GetDailyRetain(data DailyRetainRequest) (*DailyRetainResponse, error) {
	token, err := AccessTokenHandler(a.client, a.conf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("access_token: %v\n", token)
	resp, err := a.client.PostJson(
		fmt.Sprintf(dailyRetainURL, token),
		data,
	)
	if err != nil {
		log.Println("GetDailyRetain get error: ", err)
		return nil, err
	}
	if resp.Code != http.StatusOK {
		log.Printf("GetDailyRetain request incorrect,statusCode: %d", resp.Code)
		return nil, fmt.Errorf("bad request")
	}
	var ret DailyRetainResponse
	err = json.Unmarshal(resp.Body, &ret)
	if err != nil {
		log.Printf("GetDailyRetain unmarshal error: %v", err)
		return nil, err
	}
	return &ret, nil
}
