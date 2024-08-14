package middlewares

import (
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	maxRequests   = 100
	timeForbidden = 10 * time.Second
	timeToAnalyze = 1 * time.Second // Minute to test
)

type reqStats struct {
	count    int
	firstReq time.Time
	ban      time.Time
}

var raceLimitDB = make(map[string]*reqStats)
var borderMu sync.Mutex

func RaceLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			for ip, stat := range raceLimitDB {
				if time.Since(stat.firstReq) > time.Minute*1 {
					delete(raceLimitDB, ip)
				}
			}
		}()

		if locked := borderMu.TryLock(); !locked {
			http.Error(w, "Too many requests!", http.StatusRequestTimeout)
			return
		}
		defer borderMu.Unlock()

		IP := strings.Split(r.RemoteAddr, ":")[0]

		if stats, exist := raceLimitDB[IP]; exist {
			if time.Since(stats.ban) < timeForbidden {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			if time.Since(stats.firstReq) > timeToAnalyze {
				stats.count = 0
				stats.firstReq = time.Now()
			} else {
				stats.count++
				if stats.count > maxRequests {
					stats.ban = time.Now()
					http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
					return
				}
			}

		} else {
			raceLimitDB[IP] = &reqStats{count: 1, firstReq: time.Now()}
		}

		next.ServeHTTP(w, r)
	})
}
