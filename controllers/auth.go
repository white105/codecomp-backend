package controllers

import (
	"codecomp-backend/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	clientID     = "9d29c4e6e9d459681fc4"
	clientSecret = "fbd6c6db981856ff0048922f7333645a99212450"
)

//User Github OAuth controller
func GithubOAuthController(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("could not parse query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	code := r.FormValue("code")
	log.Println("code from github ", code)

	// Next, lets for the HTTP request to call the github oauth enpoint
	// to get our access token
	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		log.Fatalf("could not create HTTP request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	// We set this header since we want the response
	// as JSON
	req.Header.Set("accept", "application/json")

	// Send out the HTTP request
	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("could not send HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	var t responses.OAuthAccessResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		log.Fatalf("could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	log.Println("access token", t.AccessToken)

	// Get user info
	userInfoURL := fmt.Sprintf("https://api.github.com/user")
	userReq, err := http.NewRequest(http.MethodGet, userInfoURL, nil)
	if err != nil {
		log.Fatalf("could not create HTTP user get info request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	userReq.Header.Add("Authorization", "token "+t.AccessToken)

	// Send HTTP request to get user info
	resp, err := client.Do(userReq)
	if err != nil {
		log.Fatalf("could not send HTTP get user info request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	// User info
	log.Println(string(result))

}
