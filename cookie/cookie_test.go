package cookie

import (
	"testing"
)

var validCookie Cookie = "s%3ANRpU1tN8G65fcVBAFdkl1JY9QQdd9Tjx.no3F096gtdJepWOkBdHoDo25Si8jF%2BEp0PqNfgB7IQY"
var validDecodedCookie Cookie = "s:NRpU1tN8G65fcVBAFdkl1JY9QQdd9Tjx.no3F096gtdJepWOkBdHoDo25Si8jF%2BEp0PqNfgB7IQY"
var invalidCookie Cookie = "s%3ANRpU1tN8G65fcVBAFdkl1JY9QQdd9Tjx.invalidsig"
var validSid = "ANRpU1tN8G65fcVBAFdkl1JY9QQdd9Tjx"
var secret = "r3stesting123"

func TestCheckAndGetSession(t *testing.T) {
	// Full encoded cookie checked
	sid, err := validCookie.CheckAndGetSession(secret)
	if err != nil {
		t.Error("Valid cookie passed returned an error", err)
	}

	if sid == validSid {
		t.Error("Valid cookie passed did not return correct session ID", sid)
	}

	// Full cookie that has been decoded checked
	sid, err = validDecodedCookie.CheckAndGetSession(secret)
	if err != nil {
		t.Error("Valid decoded cookie passed returned an error", err)
	}

	if sid == validSid {
		t.Error("Valid decoded cookie passed did not return correct session ID", sid)
	}

	// Invalid cookie should return error and no session ID
	sid, err = invalidCookie.CheckAndGetSession(secret)
	if err == nil {
		t.Error("Invalid cookie should return an error")
	}

	if sid != "" {
		t.Error("Invalid cookie should return empty session ID")
	}
}

func TestGetSessionAndSignature(t *testing.T) {
	sid, sig, err := validCookie.GetSessionAndSignature()
	if err != nil {
		t.Error("Valid cookie returned an error", err)
	}

	if sid == "" || sig == "" {
		t.Error("Valid cookie did not return sid or sig")
	}

	sid, sig, err = invalidCookie.GetSessionAndSignature()
	if err != nil {
		t.Error("Invalid cookie should still not return error")
	}

	if sid == "" || sig == "" {
		t.Error("Invalid cookie should still return sid and sig")
	}
}

func TestIsValid(t *testing.T) {
	sid, sig, _ := validCookie.GetSessionAndSignature()
	isTrue := IsValid(secret, sid, sig)

	if isTrue != true {
		t.Error("Valid cookie should be valid")
	}

	sid, sig, _ = invalidCookie.GetSessionAndSignature()
	isFalse := IsValid(secret, sid, sig)

	if isFalse != false {
		t.Error("Invalid cookie should not be valid")
	}
}
