package core

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

type RateEntry struct {
	Count   int
	Expires time.Time
}

var (
	rateStore = make(map[string]*RateEntry)
	rateMutex sync.Mutex
)

func RateLimit(ctx *RequestContext, cfg *Config) {

	for path, rule := range cfg.Rules.RateLimit {

		if !strings.HasPrefix(ctx.Path, path) {
			continue
		}

		limit, duration := parseRule(rule)

		key := ctx.IP + ":" + path

		rateMutex.Lock()

		entry, exists := rateStore[key]

		if !exists || time.Now().After(entry.Expires) {
			rateStore[key] = &RateEntry{
				Count:   1,
				Expires: time.Now().Add(duration),
			}
			rateMutex.Unlock()
			return
		}

		entry.Count++

		if entry.Count > limit {
			ctx.Score += 10
		}

		rateMutex.Unlock()
	}
}

func parseRule(rule string) (int, time.Duration) {
	// example: "5/m"
	parts := strings.Split(rule, "/")

	n, _ := strconv.Atoi(parts[0])

	switch parts[1] {
	case "s":
		return n, time.Second
	case "m":
		return n, time.Minute
	case "h":
		return n, time.Hour
	default:
		return n, time.Minute
	}
}
