package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

const (
	issuer         = "huck.com.ar"
	secretKey      = "changeme"
	accessTokenKey = "access_token"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type Claims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

func (c *Claims) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return err
	}

	if !c.VerifyEmail() {
		return ErrInvalidEmail
	}

	return nil
}

func (c *Claims) VerifyEmail() bool {
	if len(c.Email) == 0 {
		return false
	}

	return true
}

func main() {
	r := chi.NewRouter()

	r.Get("/", Serve)
	r.Post("/sign", Sign)

	http.ListenAndServe(":8080", r)
}

func Sign(w http.ResponseWriter, r *http.Request) {
	expiresAt := time.Now().Add(1 * time.Hour)
	c := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
		},
		Email: "marcos@huck.com.ar",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &c)

	signed, err := token.SignedString([]byte(secretKey))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: Signing access token failed. Error: %v", http.StatusText(http.StatusInternalServerError), err), http.StatusInternalServerError)
		return
	}

	ck := http.Cookie{
		Name:     accessTokenKey,
		Value:    signed,
		Domain:   "localhost",
		Expires:  expiresAt,
		HttpOnly: true,
	}

	http.SetCookie(w, &ck)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func Serve(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie(accessTokenKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: Getting cookie failed, error: %v", http.StatusText(http.StatusUnauthorized), err), http.StatusUnauthorized)
		return
	}

	err = validateAccessToken(ck.Value)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: Invalid access token", http.StatusText(http.StatusUnauthorized)), http.StatusUnauthorized)
		return
	}

	data, err := json.Marshal(map[string]interface{}{
		"status": "OK",
		"data":   "Welcome!",
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("%s: Marshaling response failed, error: %v", http.StatusText(http.StatusInternalServerError), err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	n, err := w.Write(data)
	if n != len(data) || err != nil {
		http.Error(w, fmt.Sprintf("%s: Writing HTTP response failed", http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
		return
	}
}

func validateAccessToken(accessToken string) error {
	parser := new(jwt.Parser)

	_, err := parser.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if err := token.Claims.Valid(); err != nil {
			return nil, err
		}
		return []byte(secretKey), nil
	})

	return err

}
