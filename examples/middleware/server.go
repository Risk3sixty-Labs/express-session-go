package main

import (
	"log"
	"net/http"

	m "github.com/Risk3sixty-Labs/express-session-go/middleware"
)

func handler(w http.ResponseWriter, r *http.Request) {
	session, _ := r.Context().Value(m.SessionContextKey).(m.Session)
	sid := session.SessionID
	w.Write([]byte("Session ID: " + sid))
}

func main() {
	m.SetCookieKey("sid")
	m.SetCookieSecret("r3stesting123")
	final := m.ExpressSessionMiddleware(http.HandlerFunc(handler))

	http.Handle("/", final)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	log.Print("Successfully listening on *:8080")
}
