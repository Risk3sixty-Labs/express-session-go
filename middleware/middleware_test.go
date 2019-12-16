package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExpressSessionMiddlwareWithNoCookie(t *testing.T) {
	// Create a request to pass to our handler
	reqWithoutCookie, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := getHandler()

	SetCookieKey("risk3sixty")
	SetCookieSecret("r3stesting123")
	finalHandler := ExpressSessionMiddleware(handler)
	finalHandler.ServeHTTP(rr, reqWithoutCookie)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v: %s", status, http.StatusOK, rr.Body)
	}

	sid := rr.Body
	if sid.String() != "" {
		t.Error("session value is not set to the correct value", sid)
	}
}

func TestExpressSessionMiddlwareWithCookie(t *testing.T) {
	// valid signature in cookie
	rr1 := validKeyTester(t, "s:NRpU1tN8G65fcVBAFdkl1JY9QQdd9Tjx.no3F096gtdJepWOkBdHoDo25Si8jF%2BEp0PqNfgB7IQY")
	if status := rr1.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v: %s", status, http.StatusOK, rr1.Body)
	}

	sid := rr1.Body
	if sid.String() != "NRpU1tN8G65fcVBAFdkl1JY9QQdd9Tjx" {
		t.Error("session value is not set to the correct value", sid)
	}

	// invalid signature in cookie
	rr2 := validKeyTester(t, "s:NRpU1tN8G65fcVBAFdkl1JY9QQdd9Tjx.invalidsig")
	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v: %s", status, http.StatusOK, rr1.Body)
	}

	sid = rr2.Body
	if sid.String() != "" {
		t.Error("session value is not set to an empty value", sid)
	}
}

func validKeyTester(t *testing.T, cookieValue string) *httptest.ResponseRecorder {
	// Create a request to pass to our handler
	reqWithCookie, err := http.NewRequest("GET", "/", nil)
	sessionCookie := http.Cookie{
		Name:  "risk3sixty",
		Value: cookieValue,
	}
	reqWithCookie.AddCookie(&sessionCookie)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := getHandler()

	SetCookieKey("risk3sixty")
	SetCookieSecret("r3stesting123")
	finalHandler := ExpressSessionMiddleware(handler)
	finalHandler.ServeHTTP(rr, reqWithCookie)

	return rr
}

func getHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := r.Context().Value(SessionContextKey).(Session)
		sid := session.SessionID
		w.Write([]byte(sid))
	})
}
