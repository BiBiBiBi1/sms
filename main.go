package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type client struct {
	Username string
	Password string
}

var (
	accountSid string
	authToken  string
	fromPhone  string
	toPhone    string
)

func connect() {
	user := http.Client{}
	resp, err := user.Get("https://api.turbosms.ua")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func SendMessage(msg string) {

	params := client.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(msg)

	response, err := client.CreateMessage(&params)
	if err != nil {
		fmt.Printf("error creating and sending message: %s\n", err.Error())
		return
	}
	fmt.Printf("Message SID: %s\n", *response.Sid)
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error loading .env: %s\n", err.Error())
		os.Exit(1)
	}

	accountSid = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
	fromPhone = os.Getenv("FROM_PHONE")
	toPhone = os.Getenv("TO_PHONE")

	// создать структуру cleint, с данными user name & password.
}

func main() {

	emp := []client{}
	emp = append(emp, client{Username: accountSid, Password: authToken})
	fmt.Printf("name is: %s\n")

	msg := fmt.Sprintf(os.Getenv("MSG"), "Dima")
	SendMessage(msg)

}
