// RateLimiter.go
package main

import (
	"sync"
	"time"
)

type RateLimiter struct {
	interval time.Duration
	tokens   chan struct{}
	once     sync.Once
}

func NewRateLimiter(interval time.Duration, burst int) *RateLimiter {
	rl := &RateLimiter{
		interval: interval,
		tokens:   make(chan struct{}, burst),
	}

	rl.once.Do(func() {
		go func() {
			ticker := time.NewTicker(rl.interval)
			for {
				select {
				case <-ticker.C:
					select {
					case rl.tokens <- struct{}{}:
					default:
					}
				}
			}
		}()
	})

	return rl
}

func (rl *RateLimiter) Wait() {
	<-rl.tokens
}
