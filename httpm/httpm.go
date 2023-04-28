package httpm

import (
	"PonifiaUtils/logging"
	"math/rand"
	"net/http"
	"time"
)

var Logger = logging.GetLogger("httpm")

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			rec := recover()
			if rec != nil {
				w.WriteHeader(http.StatusInternalServerError)
				Logger.Print("panic: %v", rec)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Timer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		ref := rand.Intn(10000000)
		defer func(now time.Time) {
			Logger.Print("Timer: response %d: %dms %s %s", ref, time.Since(now).Milliseconds(), r.Method, r.URL.Path)
		}(now)
		Logger.Print("Timer: request: %d %s %s", ref, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
