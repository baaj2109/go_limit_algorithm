package tockenbucket

import (
	"sync"
	"time"
)

type TokenBucket struct {
	rate          int64
	capacity      int64
	tockenSize    int64
	laskTockenSec int64
	lock          sync.Mutex
}

func (b *TokenBucket) Set(rate, cap int64) {
	b.rate = rate
	b.capacity = cap
	b.tockenSize = 0
	b.laskTockenSec = time.Now().Unix()
}

func (b *TokenBucket) Allow() bool {
	b.lock.Lock()
	defer b.lock.Unlock()

	now := time.Now().Unix()
	b.tockenSize = b.tockenSize + (now-b.laskTockenSec)*b.rate
	if b.tockenSize > b.capacity {
		b.tockenSize = b.capacity
	}
	b.laskTockenSec = now
	if b.tockenSize > 0 {
		b.tockenSize--
		return true
	} else {
		return false
	}

}
