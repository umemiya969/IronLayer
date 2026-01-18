package core

import "time"

func StartCleanup() {
	go func() {
		for {
			time.Sleep(30 * time.Second)

			now := time.Now()

			for k, v := range rateStore {
				if now.After(v.Expires) {
					delete(rateStore, k)
				}
			}
		}
	}()
}
