package core

import (
	"sync"
	"time"
)

type BanEntry struct {
	TempCount int
	BanUntil  time.Time
	Permanent bool
}

var (
	banStore = make(map[string]*BanEntry)
	banMutex sync.Mutex
)

func CheckBan(ip string) bool {
	banMutex.Lock()
	defer banMutex.Unlock()

	entry, exists := banStore[ip]
	if !exists {
		return false
	}

	if entry.Permanent {
		return true
	}

	if time.Now().Before(entry.BanUntil) {
		return true
	}

	return false
}

func ApplyBan(ip string) {
	banMutex.Lock()
	defer banMutex.Unlock()

	entry, exists := banStore[ip]

	if !exists {
		banStore[ip] = &BanEntry{
			TempCount: 1,
			BanUntil:  time.Now().Add(15 * time.Minute),
		}
		return
	}

	if entry.Permanent {
		return
	}

	entry.TempCount++

	if entry.TempCount >= 3 {
		entry.Permanent = true
		return
	}

	entry.BanUntil = time.Now().Add(15 * time.Minute)
}
