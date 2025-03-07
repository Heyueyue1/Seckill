package limiter

import (
	"encoding/json"
	"fmt"
	"github.com/BitofferHub/gateway/internal/conf"
	"github.com/redis/go-redis/v9"
	"os"
)

var Rl *RateLimiter

func InitLimiter(routeConfigPath string, redisClient *redis.Client,
	defaultRetryTime, defaultLimitTimeout, defaultLimitRate int) error {

	routes, err := os.ReadFile(routeConfigPath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	conf.Routes = make(map[string]conf.Route, 0)
	err = json.Unmarshal(routes, &conf.Routes)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	rateLimiterConfig := RateLimiterConfig{
		Routes:              conf.Routes,
		DefaultRetryTime:    defaultRetryTime,
		DefaultLimitTimeout: defaultLimitTimeout,
		DefaultLimitRate:    defaultLimitRate,
	}
	rl, err := NewRateLimiter(rateLimiterConfig, redisClient)

	if err != nil {
		return err
	}
	Rl = rl
	return nil
}
