package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var secret = []byte("very-secret-key")

type Claims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Exp  int64  `json:"exp"`
}

func GenerateToken(name, role string) (string, error) {
	claims := Claims{
		Name: name,
		Role: role,
		Exp:  time.Now().Add(24 * time.Hour).Unix(),
	}

	payload, _ := json.Marshal(claims)
	encoded := base64.RawURLEncoding.EncodeToString(payload)

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(encoded))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	return encoded + "." + signature, nil
}

func ValidateToken(token string) (*Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return nil, errors.New("invalid token")
	}

	payload, sig := parts[0], parts[1]

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(payload))
	expected := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	if sig != expected {
		return nil, errors.New("invalid signature")
	}

	data, _ := base64.RawURLEncoding.DecodeString(payload)
	var claims Claims
	json.Unmarshal(data, &claims)

	if time.Now().Unix() > claims.Exp {
		return nil, errors.New("token expired")
	}

	return &claims, nil
}
