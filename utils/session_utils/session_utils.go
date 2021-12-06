package session_utils

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func SetSession(w http.ResponseWriter, req *http.Request, sName string, path string) {
	// get cookie
	c, err := req.Cookie(sName)
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:    sName,
			Value:   sID.String(),
			Expires: time.Now().Add(60 * time.Minute),
			Path:    path,
		}

	}
	http.SetCookie(w, c)
}

func CheckLogin(req *http.Request, sName string) bool {
	t, _ := req.Cookie(sName)
	if t == nil {
		return false
	}
	return true
}
