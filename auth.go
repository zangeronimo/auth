package auth

import (
	"fmt"
)

//Hello is a basic function for test
func Hello() string {
	return "Hello world."
}

//BasicRequest populate this struct with BasicAuth username and password geted in client request
type BasicRequest struct {
	reqUsername string
	reqPassword string
}

//BasicLocal create this struct with Basic username and password requested in your project
type BasicLocal struct {
	locUsername string
	locPassword string
}

//BasicAuth authentication with basic login, use this to validate a request for a new JWT
func BasicAuth(basicRequest BasicRequest, basicLocal BasicLocal) (bool, error) {
	if (basicRequest.reqUsername == basicLocal.locUsername) &&
		(basicRequest.reqPassword == basicLocal.locPassword) {
		return true, nil
	}
	return false, fmt.Errorf("authorization failed")
}
