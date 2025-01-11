package middleware

import (
	"errors"
	"net/http"

	"forum/utils"
)

func CheckMethod(f http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			err := errors.New("method not allowed")
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusMethodNotAllowed}, w) {
				return
			}
		}
		f(w, r)
	}
}
