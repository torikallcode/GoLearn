package middleware

import (
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"golearn/internal/response"
)

type bucket struct {
	start time.Time
	count int
}

// RateLimit is a small fixed-window limiter. Good enough for one API instance.
// ponytail: in-memory only; swap for Redis when the API scales horizontally.
func RateLimit(perMinute int) gin.HandlerFunc {
	var mu sync.Mutex
	buckets := map[string]bucket{}
	return func(c *gin.Context) {
		key := strconv.FormatUint(UserID(c), 10)
		if key == "0" {
			key = c.ClientIP()
		}
		now := time.Now()
		mu.Lock()
		b, ok := buckets[key]
		if !ok || now.Sub(b.start) >= time.Minute {
			b = bucket{start: now}
		}
		b.count++
		buckets[key] = b
		allowed := b.count <= perMinute
		mu.Unlock()
		if !allowed {
			response.Fail(c, 429, "rate_limited", "Terlalu banyak permintaan. Coba lagi sebentar.")
			return
		}
		c.Next()
	}
}
