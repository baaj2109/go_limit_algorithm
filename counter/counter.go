package counter

import (
	"sync"
	"time"
)

type LimitRate struct {
	rate  int
	begin time.Time
	cycle time.Duration
	count int
	lock  sync.Mutex
}

func (limit *LimitRate) Allow() bool {
	limit.lock.Lock()
	defer limit.lock.Unlock()

	if limit.count == limit.rate-1 {
		now := time.Now()
		if now.Sub(limit.begin) >= limit.cycle {
			limit.Reset(now)
			return true
		}
		return false
	} else {
		limit.count++
		return true
	}
}

func (limit *LimitRate) Reset(begin time.Time) {
	limit.begin = begin
	limit.count = 0
}

func (limit *LimitRate) Set(rate int, cycle time.Duration) {
	limit.rate = rate
	limit.begin = time.Now()
	limit.cycle = cycle
	limit.count = 0
}
