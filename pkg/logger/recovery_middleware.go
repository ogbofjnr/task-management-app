package logger

import (
	"errors"
	"github.com/ogbofjnr/maze/responses"
	"net/http"
	"runtime/debug"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("unknown error")
				}
				trace := debug.Stack()

				Logger.Error(err.Error()+string(trace))
				responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), Logger)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
