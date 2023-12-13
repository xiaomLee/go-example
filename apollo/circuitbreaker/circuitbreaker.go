package circuitbreaker

import (
	"fmt"
	"math/rand"

	"github.com/afex/hystrix-go/hystrix"
)

func Hystrix() error {
	hystrix.Configure(map[string]hystrix.CommandConfig{
		"testttt": hystrix.CommandConfig{
			Timeout:                0,
			MaxConcurrentRequests:  5,
			RequestVolumeThreshold: 5,
			SleepWindow:            2000,
			ErrorPercentThreshold:  3,
		},
	})
	return hystrix.Do("testttt", run, fallback)
}

func InitCommandConf() {
	hystrix.Configure(nil)
}

func run() error {
	// talk to other services
	i := rand.Int31n(100)
	if i < 50 {
		fmt.Println("rand err", i)
		return fmt.Errorf("rand err:%d", i)
	}
	return nil
}

func fallback(err error) error {
	fmt.Println("handle err:", err)
	return fmt.Errorf("fallback handle err and return circuit response")
}
