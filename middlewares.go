package gohttplib

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func AccessMiddlewareFactory(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Secret") != secret {
				DefaultHTTP403().Write(w)
				return
			}
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func RecoverMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				HTTP500(err).Write(w)
				return
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

