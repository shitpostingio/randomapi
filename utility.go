package main

import "time"

func startCleanRequestRoutine() {
	for range time.NewTicker(10 * time.Minute).C {
		for key := range requestedPosts {
			if time.Now().Sub(requestedPosts[key].requestdate) > 10*time.Minute {
				delete(requestedPosts, key)
			}
		}
	}
}
