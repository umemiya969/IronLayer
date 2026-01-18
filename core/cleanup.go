package core

import "time"

func StartBanCleanup() {
	go func() {
		for {
			time.Sleep(1 * time.Minute)

			now := time.Now()
			for ip, b := range banStore {
				if !b.Permanent && now.After(b.BanUntil) {
					delete(banStore, ip)
				}
			}
		}
	}()
}
