package cookie

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/url"
	"strings"
)

// Cookie is the full cookie populated from an express session
type Cookie string

// CheckAndGetSession takes a full cookie value and returns
// the session ID if it checks out to be valid and secure, or
// an error otherwise.
func (c Cookie) CheckAndGetSession(secret string) (sid string, err error) {
	sid, sig, err := c.GetSessionAndSignature()
	if err != nil {
		return "", err
	}

	if !IsValid(secret, sid, sig) {
		return "", errors.New("session ID and signature are not valid")
	}

	return sid, nil
}

// GetSessionAndSignature takes a full cookie value populated from
// a express session and returns the session ID and signature.
func (c Cookie) GetSessionAndSignature() (sessionID string, signature string, err error) {
	sC := string(c)
	sC, err = url.QueryUnescape(sC)
	if err != nil {
		return "", "", err
	}

	if len(sC) <= 2 {
		return "", "", errors.New("cookie is unexpectedly short")
	}

	sIDAndSig := strings.Split(sC, ":")[1]
	idAndSig := strings.Split(sIDAndSig, ".")

	return idAndSig[0], idAndSig[1], nil
}

// IsValid checks that the session and signature are valid
// based on the session secret used to create the session ID and
// signature.
func IsValid(secret string, sessionID string, sig string) bool {
	sessionIDBytes := []byte(sessionID)
	secretBytes := []byte(secret)

	sha256Signer := hmac.New(sha256.New, secretBytes)
	sha256Signer.Write(sessionIDBytes)
	calculatedSignature := sha256Signer.Sum(nil)
	calculatedBase64Signature := strings.TrimRight(base64.StdEncoding.EncodeToString(calculatedSignature), "=")

	return sig == calculatedBase64Signature
}
