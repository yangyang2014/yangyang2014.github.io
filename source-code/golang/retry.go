package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
}

// 重试策略接口
// 打印日志
type RetryStrategy interface {
	ShouldRetry(err error) (bool, time.Duration)
	GetRetryCount() int
	StopRetry() bool
	PrintLog(err error)
}

// 要求 err！=nil 时，重试   , 写函数时需要考虑考
var ApiCall func() (interface{}, error)

// 策略有哪些
// 指数退避+随机的策略
// 固定时间退避策略
// 最大重试次数

func retry(ApiCall func() (interface{}, error), strategy RetryStrategy) (interface{}, error) {
	response, err := ApiCall()
	if err == nil {
		return response, nil
	}
	for retry, duration := strategy.ShouldRetry(err); retry; retry, duration = strategy.ShouldRetry(err) {
		time.Sleep(duration)
		response, err = ApiCall()
		if err == nil {
			break
		}
	}
	return response, err
}
