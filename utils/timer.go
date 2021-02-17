package utils

import (
	"time"
)

// 简易计时器，每过一定时间就发出信号
func TickTock(ch chan bool, duration time.Duration) {
	for {
		ch <- true
		time.Sleep(duration)
	}
}
