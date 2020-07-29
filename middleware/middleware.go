package middleware

import (
	"context"
	"log"
	"net/http"

	c "github.com/whatl3y/express-session-go/cookie"
	s "github.com/whatl3y/express-session-go/store"
)

type contextKey string

// Session is the struct that stores all relavant session data
// that a user can access via a request context
type Session struct {
	SessionID    string
	SessionData  interface{}
	SessionStore s.BaseStore
}

// SessionParser takes a session ID and returns a processed
// session that we want to store in our request context
type SessionParser func(string) (interface{}, error)

// SessionContextKey represents all of the session data
const SessionContextKey contextKey = "session"

var mOptions struct {
	CookieKey    string
	CookieSecret string
	Logger       *log.Logger
	Store        s.BaseStore
}

// SetCookieKey sets the cookie key
func SetCookieKey(key string) {
	mOptions.CookieKey = key
}

// SetCookieSecret sets the cookie secret for the signature
func SetCookieSecret(secret string) {
	mOptions.CookieSecret = secret
}

// SetLogger sets the logger to log info while parsing the session
func SetLogger(logger *log.Logger) {
	mOptions.Logger = logger
}

// SetStore sets the session store
func SetStore(store s.BaseStore) {
	mOptions.Store = store
}

// ExpressSessionMiddleware is HTTP middleware to retrieve and
// optionally parse/process session information that was originally
// populated from ExpressJS (express-session)
func ExpressSessionMiddleware(next http.Handler) http.Handler {
	if mOptions.Store == nil {
		memoryStore := make(s.MemoryStore)
		mOptions.Store = &memoryStore
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(mOptions.CookieKey)
		if err != nil {
			next.ServeHTTP(w, r)
			log.Print("Error getting cookie", err)
			return
		}

		cookieValue := c.Cookie(cookie.Value)
		sessionID, err := cookieValue.CheckAndGetSession(mOptions.CookieSecret)
		if err != nil {
			next.ServeHTTP(w, r)
			log.Print("Error getting session ID", err)
			return
		}

		parser := storeSessionParser(mOptions.Store)
		parsedSession, err := parser(sessionID)
		if err != nil {
			next.ServeHTTP(w, r)
			log.Print("Error getting parsed session", err)
			return
		}

		sessionData := Session{sessionID, parsedSession, mOptions.Store}
		ctxWithSession := context.WithValue(r.Context(), SessionContextKey, sessionData)
		rWithSession := r.WithContext(ctxWithSession)

		next.ServeHTTP(w, rWithSession)
	})
}

func storeSessionParser(store s.BaseStore) SessionParser {
	return func(sid string) (interface{}, error) {
		extendedStore, isExtended := store.(s.ExtendedStore)
		session, err := store.Get(sid)
		if session != nil && isExtended {
			extendedStore.Touch(sid)
		}

		return session, err
	}
}
