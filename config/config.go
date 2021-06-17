package config

import "mini/cache"

type Config struct {
	AppId     string
	AppSecret string
	Cache     cache.Cache
}
