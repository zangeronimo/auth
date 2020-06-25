package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var basicUser = ""
var basicPass = ""

//SetBasicUser add a basic username to validade a basic login
func SetBasicUser(bu string) {
	basicUser = bu
}

//SetBasicPass add a basic password to validade a basic login
func SetBasicPass(bp string) {
	basicPass = bp
}

//basicAuth authentication with basic login, use this to validate a request for a new JWT
func basicAuth(r *http.Request) error {
	// Check basic login for basic authentication, first stap to generate a JWT
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		return fmt.Errorf("authorization failed")
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) == 2 && pair[0] == basicUser && pair[1] == basicPass {
		return nil
	}
	return fmt.Errorf("authorization failed")
}

//Login Use this func to authenticate a user and return a valid jwt
func Login(w http.ResponseWriter, r *http.Request) {
	err := basicAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SAMPLESECRET")))

	fmt.Println(tokenString, err)
	fmt.Fprintf(w, "Hello, %s - Err %s!", tokenString, err)
}
