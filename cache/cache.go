package cache

import "time"

const (
	// DefaultExpireTime 默认过期时间
	DefaultExpireTime = time.Hour * 24
	// DefaultNotFoundExpireTime 结果为空时的过期时间 1分钟, 常用于数据为空时的缓存时间(缓存穿透)
	DefaultNotFoundExpireTime = time.Minute
	// NotFoundPlaceholder .
	NotFoundPlaceholder = "*"
)

var Client Cache

type Cache interface {
	Set(key string, val interface{}, expiration time.Duration) error
	Get(key string, val interface{}) error
	Del(keys ...string) error
}

func Set(key string, val interface{}, expiration time.Duration) error {
	return Client.Set(key, val, expiration)
}

func Get(key string, val interface{}) error {
	return Client.Get(key, val)
}

func Del(keys ...string) error {
	return Client.Del(keys...)
}
