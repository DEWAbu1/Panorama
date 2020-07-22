package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var CookieName = "SalesLetter"

var Store = sessions.NewCookieStore([]byte("SalesLetter"))
var Session = sessions.NewSession(Store, CookieName)

func init() {
	// domain := os.Getenv("domain")
	domain := "localhost"
	Store.Options = &sessions.Options{
		Domain:   domain,
		Path:     "/",
		MaxAge:   3600 * 2, // 2 hours
		HttpOnly: true,
	}

}

func GetSession(r *http.Request) *sessions.Session {
	session, _ := Store.Get(r, CookieName)

	return session
}
