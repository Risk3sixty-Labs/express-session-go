package middleware

import (
	"context"
	"net/http"

	c "github.com/Risk3sixty-Labs/express-session-go/cookie"
)

type sessionParser func(string) interface{}
type contextKey string

// SessionContextKey is a const representing
// the session key we're populating in the request context.
const SessionContextKey contextKey = "session"

// ExpressSessionMiddleware is HTTP middleware to
func ExpressSessionMiddleware(cookieKey string, cookieSecret string, parser sessionParser, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieKey)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		cookieValue := c.Cookie(cookie.Value)
		sessionID, err := cookieValue.CheckAndGetSession(cookieSecret)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctxWithSession := context.WithValue(r.Context(), SessionContextKey, parser(sessionID))
		rWithSession := r.WithContext(ctxWithSession)
		next.ServeHTTP(w, rWithSession)
	})
}
