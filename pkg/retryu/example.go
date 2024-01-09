package retryu

import (
	"fmt"
	retry "github.com/avast/retry-go/v4"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type MyTimer struct{}

func (t *MyTimer) After(d time.Duration) <-chan time.Time {
	fmt.Print("Timer called!")
	return time.After(d)
}

func retry1() {
	url := "http://wwwxxx.baidu.com"
	var body []byte

	if err := retry.Do(
		func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			return nil
		},
		retry.Attempts(3),                   // 重试次数，默认10次
		retry.Context(context.Background()), // Context 允许设置重试的上下文，默认是后台上下文
		retry.Delay(1*time.Second),          // 延迟设置重试之间的延迟默认为100ms
		retry.DelayType(retry.FixedDelay),   // 固定延迟，使用Delay的参数
		//retryu.DelayType(retryu.BackOffDelay), // DelayType 设置重试之间的延迟类型，默认为 BackOff
		//retryu.DelayType(retryu.RandomDelay),  // 随机延迟， 在 0 - MaxJitter 之间
		//retryu.DelayType(retryu.CombineDelay(retryu.FixedDelay, retryu.RandomDelay, retryu.BackOffDelay)), // 将所有指定的延迟合并到一起

		retry.LastErrorOnly(true),      // LastErrorOnly 仅返回最后一次的错误，默认为 false
		retry.MaxDelay(10*time.Second), // MaxDelay 设置重试之间的最大延迟默认情况下不适用
		retry.MaxJitter(1*time.Second), // MaxJitter 设置 RandomDelay 重试之间的最大随机抖动
		retry.OnRetry(func(n uint, err error) { // 每次重试都会调用 OnRetry 函数回调
			log.Printf("retryu %d: %v", n, err)
		}),
		retry.RetryIf(func(err error) bool { // RetryIf 允许你自定义重试条件，默认情况下，当 err 不为 nil 时，重试
			if err.Error() == "special error" {
				return false
			}
			return true
		}),
		retry.WithTimer(&MyTimer{}),               // WithTimer 设置重试的超时时间，默认情况下，重试不会超时
		retry.WrapContextErrorWithLastError(true), // WrapContextErrorWithLastError 允许将上下文错误与重试函数返回的最后一个错误一起返回。仅当 Attempts 设置为 0 以无限期重试以及使用上下文取消/超时时才适用
	); err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}

func retry2() {
	url := "http://example.com"

	body, err := retry.DoWithData(
		func() ([]byte, error) {
			resp, err := http.Get(url)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			return body, nil
		},
	)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func retry3() {
	url := "http://example.com"
	var body []byte

	err := retry.Do(
		func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			return nil
		},
	)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
