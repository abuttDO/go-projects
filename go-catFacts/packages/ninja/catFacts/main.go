package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type CatFact struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}


func GetCatFact() string{
	url := "https://catfact.ninja/fact"

	var catFact CatFact

	err := GetJson(url, &catFact)
	if err != nil {
		fmt.Printf("Error getting cat facts: %s\n", err.Error())
	} else {
		fmt.Printf("A Super interesting Cat Fact: %s\n", catFact.Fact)
	}

	return catFact.Fact

}

func Main() map[string]interface{} {

	client = &http.Client{Timeout: 10 * time.Second}

	msg := make(map[string]interface{})

	msg["body"] = GetCatFact()

	return msg

}