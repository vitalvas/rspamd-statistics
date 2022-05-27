package main

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func contentTypeMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		if req.Header.Get("content-type") != "application/msgpack" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(http.StatusText(http.StatusUnsupportedMediaType)))

			return nil
		}

		return next(w, req)
	}
}
