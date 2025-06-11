package main

import (
	"errors"
	"net/http"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	username := r.FormValue("username")
	user, ok := users[username]
	if !ok {
		return AuthError
	}

	// Get the Session Token from the cookie
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		return AuthError
	}

	//Get CSRF token from headers
	csrf := r.Header.Get("X-CSRF-TOKEN")
	if csrf != user.CSRFToken || csrf == "" {
		return AuthError
	}

	return nil
}
