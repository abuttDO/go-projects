package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	accountSid string
	authToken string
	fromPhone string
	toPhone string
	client *twilio.RestClient
)

func SendMessage(msg string) {

	params := openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(msg)

	response, err := client.Api.CreateMessage(&params)
	if err != nil {
		fmt.Printf("error creatinG and sending message: %s\n", err.Error())
		return
	}
	fmt.Printf("Message SID: %s\nMESSAGE SEND SUCCESSFULLY", *response.Sid)
}

func init() {
	
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Printf("error loading .env: %s\n", err.Error())
		os.Exit(1)
	}
	
	accountSid = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
	fromPhone = os.Getenv("FROM_PHONE")
	toPhone = os.Getenv("TO_PHONE")
	/*
	env, ok := os.LookupEnv("ACCOUNT_SID")

	if !ok {
		fmt.Println("the env var ACCOUNT_SID is not set")
	} else {
		fmt.Println("ACCOUNT_SID: ", env)
	}

	envv, ok := os.LookupEnv("AUTH_TOKEN")

	if !ok {
		fmt.Println("the env var AUTH_TOKEN is not set")
	} else {
		fmt.Println("AUTH_TOKEN: ", envv)
	}

	envvv, ok := os.LookupEnv("FROM_PHONE")

	if !ok {
		fmt.Println("the env var FROM_PHONE is not set")
	} else {
		fmt.Println("FROM_PHONE: ", envvv)
	}

	envvvv, ok := os.LookupEnv("TO_PHONE")

	if !ok {
		fmt.Println("the env var TO_PHONE is not set")
	} else {
		fmt.Println("TO_PHONE: ", envvvv)
	}
	*/
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

func main() {
	msg := fmt.Sprintf(os.Getenv("MSG"), "AHMED")
	SendMessage(msg)
}