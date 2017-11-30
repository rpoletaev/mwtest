package autch

import (
	"encoding/base64"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
	authRequiredHeader  = "Basic realm=Authorization Required"
)

type authService struct {
	accounts map[string]struct{}
}

func New() *authService {
	return &authService{
		accounts: make(map[string]struct{}),
	}
}

func (as *authService) SetBasicAuth(login, password string) {
	lpPair := login + ":" + password
	encPair := base64.StdEncoding.EncodeToString([]byte(lpPair))
	encStr := "Basic " + encPair
	as.accounts[encStr] = struct{}{}
}

func (as *authService) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authData := r.Header.Get(authorizationHeader)

		_, ok := as.accounts[authData]
		if !ok {
			sendAuthRequired(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func sendAuthRequired(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", authRequiredHeader)
	w.WriteHeader(http.StatusUnauthorized)
}
