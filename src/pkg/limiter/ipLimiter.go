package limiter

import (
	"golang.org/x/time/rate"
	"sync"
)

type IPLimiter struct {
	ips   map[string]*rate.Limiter
	mutex *sync.RWMutex
	rate  rate.Limit
	burst int
}

func NewIpLimiter(r rate.Limit, burst int) *IPLimiter {
	return &IPLimiter{
		ips:   make(map[string]*rate.Limiter),
		mutex: &sync.RWMutex{},
		rate:  r,
		burst: burst,
	}
}

func (i *IPLimiter) AddIp(ip string) *rate.Limiter {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	limiter := rate.NewLimiter(i.rate, i.burst)

	i.ips[ip] = limiter
	return limiter
}

func (i *IPLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mutex.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mutex.Unlock()
		return i.AddIp(ip)
	}
	i.mutex.Unlock()
	return limiter
}
