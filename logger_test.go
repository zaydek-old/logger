package logger_test

import (
	"io/ioutil"
	"sync"
	"testing"
	"time"

	"github.com/zaydek/logger"
)

func TestStress(t *testing.T) {
	l := logger.New(ioutil.Discard, func() string { return time.Now().Format("15:04:05.000000") })
	var wg sync.WaitGroup
	for x := 0; x < 1e6; x++ {
		wg.Add(1)
		go func() { l.Println(); wg.Done() }()
	}
	wg.Wait()
}
