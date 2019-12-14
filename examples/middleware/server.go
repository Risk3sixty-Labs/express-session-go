package main

import (
	"log"
	"net/http"

	m "github.com/Risk3sixty-Labs/express-session-go/middleware"
)

func handler(w http.ResponseWriter, r *http.Request) {
	sid, _ := r.Context().Value(m.SessionContextKey).(string)
	w.Write([]byte("Session ID: " + sid))
}

func main() {
	final := m.ExpressSessionMiddleware("sid", "r3stesting123", func(sessionID string) interface{} {
		return sessionID
	}, http.HandlerFunc(handler))

	http.Handle("/", final)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	log.Print("Successfully listening on *:8080")
}
