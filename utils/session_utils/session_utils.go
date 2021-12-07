package session_utils

import (
	"net/http"
	"time"
)

func SetSession(w http.ResponseWriter, req *http.Request, sName string, path string, value string) {
	c, err := req.Cookie(sName)
	if err != nil {
		c = &http.Cookie{
			Name:    sName,
			Value:   value,
			Expires: time.Now().Add(60 * time.Minute),
			Path:    path,
		}

	}
	http.SetCookie(w, c)
}

func ClearSession(w http.ResponseWriter, sName string) {
	c := &http.Cookie{
		Name:   sName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
}

func CheckSession(req *http.Request, sName string) bool {
	t, _ := req.Cookie(sName)
	if t == nil {
		return false
	}
	return true
}
