package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	var returnedUser []User

	res, err := http.Get(url)
	if err != nil {
		return returnedUser, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&returnedUser); err != nil {
		return returnedUser, err
	}
	return returnedUser, err
}
