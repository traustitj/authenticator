package authenticator

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"strings"
)

type Authenticator struct {
	Login  string
	Logout string
	Secret string
	Token  string
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{"/login", "/logout", "secret", "auth"}
}

func (a *Authenticator) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var hashKey = []byte(a.Secret)
	var blockKey = []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	var s = securecookie.New(hashKey, blockKey)

	if !strings.Contains(r.URL.Path, "admin") {
		next(rw, r)
		return
	}

	var illegalCookie bool
	if cookie, cookieerr := r.Cookie("user"); cookieerr == nil {
		results := make(map[string]string)
		tampered := s.Decode("user", cookie.Value, &results)
		if tampered != nil {
			http.Error(rw, "Unauthorized", 401)
			illegalCookie = true
		} else {
			illegalCookie = false
		}
	} else {
		illegalCookie = false
	}

	if illegalCookie == false {
		_, err := r.Cookie("user")
		if err != nil {
			http.Redirect(rw, r, a.Login, 401)
			cookie := http.Cookie{Name: "redirect", Value: r.URL.Path, Path: "/"}
			http.SetCookie(rw, &cookie)
		} else {
			next(rw, r)
		}
		return
	}
	next(rw, r)
}
