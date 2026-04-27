package middleware

import (
	"errors"
	"net/http"

	"github.com/promise111/goapi/api"
	"github.com/prommise111/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")

		if token == "" || username == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
		}
	})
}
