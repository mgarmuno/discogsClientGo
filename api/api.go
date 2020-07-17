package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mgarmuno/discogsClientGo/model"

	"github.com/mgarmuno/discogsClientGo/file"
)

const (
	discogsURL = "https://api.discogs.com"
	userInfo   = "/oauth/identity"
	token      = "token"
)

// UpdateUserInfo updates the user info from Discogs
func UpdateUserInfo() {
	fmt.Println("User info request")
	req, err := http.NewRequest("GET", discogsURL+userInfo, nil)
	if err != nil {
		log.Println("Error creating GET HTTP request:", err)
	}
	query := req.URL.Query()
	query.Add(token, file.GetToken())
	req.URL.RawQuery = query.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error executing GET HTTP request:", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
	}
	fmt.Println(string(bodyBytes))
	var user = new(model.User)
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Println("Error decoding response:", err)
	}

	fmt.Println("id:", user.DiscogsID)
	fmt.Println("username:", user.UserName)
	fmt.Println("profile:", user.ProfileURL)
}
