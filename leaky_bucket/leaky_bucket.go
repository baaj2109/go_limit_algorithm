package leakybucket

import (
	"math"
	"sync"
	"time"
)

type LeakyBucket struct {
	rate       float64
	capacity   float64
	water      float64
	lastLeakMs int64
	lock       sync.Mutex
}

func (leaky *LeakyBucket) Allow() bool {
	leaky.lock.Lock()
	defer leaky.lock.Unlock()

	now := time.Now().UnixNano() / 1e6
	leakyWater := leaky.water - (float64(now-leaky.lastLeakMs) * leaky.rate / 1000)
	leakyWater = math.Max(0, leakyWater)
	leaky.lastLeakMs = now
	if leaky.water+1 <= leaky.capacity {
		leaky.water++
		return true
	} else {
		return false
	}
}

func (leaky *LeakyBucket) Set(rate, capacity float64) {
	leaky.rate = rate
	leaky.capacity = capacity
	leaky.water = 0
	leaky.lastLeakMs = time.Now().UnixNano() / 1e6
}
