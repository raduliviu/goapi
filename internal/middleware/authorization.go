package middleware

import (
	"errors"
	"net/http"

	"github.com/raduliviu/goapi/api"
	"github.com/raduliviu/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		// var loginDetails *tools.LoginDetails
		loginDetails := (*database).GetUserLoginDetails(username)
		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}