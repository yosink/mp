package mini

import (
	"mini/cache"
	"mini/config"
	"mini/provider"
	"testing"

	"github.com/go-redis/redis/v8"
)

func newApp() *Application {
	rdsClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "secret",
		DB:       0,
	})
	app := New(&config.Config{
		AppId:     "wx738b0d8ef76f77d7",
		AppSecret: "dc2d9495bc5d43a66ff73b5972176499",
		Cache:     cache.NewRedisCache(rdsClient, ""),
	})
	return app
}

func TestApplication_Auth(t *testing.T) {
	// get miniprogram app
	app := newApp()
	// get a provider
	auth := app.Auth()
	// call a service
	resp, err := auth.Code2Session("021mLt0w3a68zW21XJ2w3mJL1Q2mLt0i")
	if err != nil {
		t.Errorf("GetAccessToken error: %v", err)
	}
	t.Logf("resp: %+v", resp)
}

func TestApplication_Auth_PaidUnionId(t *testing.T) {
	app := newApp()
	auth := app.Auth()
	resp, err := auth.GetPaidUnionId("oErfb4iaFRTImCHLEHbak2oP0_Lo")
	if err != nil {
		t.Errorf("auth.GetPaidUnionId error: %v", err)
	}
	t.Logf("resp: %+v", resp)
}

func TestApplication_Analysis(t *testing.T) {
	app := newApp()
	req := provider.DailyRetainRequest{
		BeginDate: "20210616",
		EndDate:   "20210617",
	}
	resp, err := app.Analysis().GetDailyRetain(req)
	if err != nil {
		t.Errorf("Analysis().GetDailyRetain error: %v", err)
	}
	t.Logf("resp: %+v", resp)
}
