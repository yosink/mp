package provider

import (
	"encoding/json"
	"fmt"
	"log"
	"mini/config"
	"mini/req"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	ErrorResponse
}

func getAccessToken(client req.HttpClient, c *config.Config) (*AccessTokenResponse, error) {
	resp, err := client.Get(fmt.Sprintf(accessTokenURL, c.AppId, c.AppSecret), nil)
	if err != nil {
		log.Println("GetAccessToken get error: ", err)
		return nil, err
	}
	if resp.Code != http.StatusOK {
		log.Printf("GetAccessToken request incorrect,statusCode: %d", resp.Code)
		return nil, fmt.Errorf("bad request")
	}
	var ret AccessTokenResponse
	err = json.Unmarshal(resp.Body, &ret)
	if err != nil {
		log.Printf("GetAccessToken unmarshal error: %v", err)
		return nil, err
	}
	return &ret, nil
}

func AccessTokenHandler(client req.HttpClient, c *config.Config) (string, error) {
	var s string
	err := c.Cache.Get(accessTokenCacheKey, &s)
	if err != nil {
		return "", err
	}
	if s == "" {
		resp, err := getAccessToken(client, c)
		if err != nil {
			return "", errors.Wrapf(err, "request accessToken error")
		}
		if resp.Errcode != 0 {
			return "", errors.Wrapf(err, "response accessToken error")
		}
		err = c.Cache.Set(accessTokenCacheKey, resp.AccessToken, time.Duration(resp.ExpiresIn-100)*time.Second)
		if err != nil {
			return "", errors.Wrapf(err, "cache accessToken error")
		}
		return resp.AccessToken, nil
	}
	return s, nil
}
