package tokens

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func registrationsHandler(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	if req.FormValue("username") == "" || req.FormValue("password") == "" {

		fmt.Fprintf(w, "Please enter a valid username and password.\r\n")

	} else {

		response, err := registerUser(req.FormValue("username"), req.FormValue("password"))

		if err != nil {
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprintf(w, response)
		}
	}

}

func authenticationsHandler(w http.ResponseWriter, req *http.Request) {

	username, password, ok := req.BasicAuth()

	if ok {

		tokenDetails, err := generateToken(username, password)

		if err != nil {
			fmt.Fprintf(w, err.Error())
		} else {

			enc := json.NewEncoder(w)
			enc.SetIndent("", "  ")
			enc.Encode(tokenDetails)
		}
	} else {

		fmt.Fprintf(w, "You require a username/password to get a token.\r\n")
	}

}

func testResourceHandler(w http.ResponseWriter, req *http.Request) {

	authToken := strings.Split(req.Header.Get("Authorization"), "Bearer ")[1]

	userDetails, err := validateToken(authToken)

	if err != nil {

		fmt.Fprintf(w, err.Error())

	} else {

		username := fmt.Sprint(userDetails["username"])

		fmt.Fprintf(w, "Welcome, "+username+"\r\n")
	}

}
