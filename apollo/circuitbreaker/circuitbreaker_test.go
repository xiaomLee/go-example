package circuitbreaker

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestHystrix(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 11; i++ {
		wg.Add(1)
		go func(int2 int) {
			defer wg.Done()
			if err := Hystrix(); err != nil {
				t.Error(err)
			}
		}(i)
		time.Sleep(1 * time.Second)
	}
	wg.Wait()

	fmt.Printf("请求完毕\n")
}
