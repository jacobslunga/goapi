package middleware

import (
  "errors"
  "net/http"

  "github.com/jacobslunga/goapi/api"
  "github.com/jacobslunga/goapi/internal/tools"
  log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    token := r.Header.Get("Authorization")

    if username == "" || token == "" {
      log.Error(UnAuthorizedError)
      api.RequestErrorHandler(w, UnAuthorizedError)
      return
    }

    var database *tools.DatabaseInterface
    database, err := tools.NewDatabase()

    if err != nil {
      api.InternalErrorHandler(w, UnAuthorizedError)
      return
    }

    var loginDetails *tools.LoginDetails
    loginDetails = (*database).GetUserLoginDetails(username)

    if (loginDetails == nil || (token != (*loginDetails).Authtoken)) {
      log.Error(UnAuthorizedError) 
      api.RequestErrorHandler(w, UnAuthorizedError)
      return
    }

    next.ServeHTTP(w, r)
  })
}

