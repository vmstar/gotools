package ratelimiter

import (
	"sync"
	"time"
)

//copy from https://github.com/vitessio/vitess/blob/master/go/ratelimiter/ratelimiter.go
type RateLimiter struct {
	maxCount int
	interval time.Duration

	mu       sync.Mutex
	curCount int
	lastTime time.Time
}

func NewRateLimiter(maxCount int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		maxCount: maxCount,
		interval: interval,
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	if time.Since(rl.lastTime) < rl.interval {
		if rl.curCount > 0 {
			rl.curCount--
			return true
		}
		return false
	}
	rl.curCount = rl.maxCount - 1
	rl.lastTime = time.Now()
	return true
}
