package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jasonlvhit/gocron"
)

var (
	// userEmail and userPassword are used for gmail smtp server
	userEmail    = os.Getenv("userEmail")
	userPassword = os.Getenv("userPassword")

	// serverHost and serverPort are used for generating the url
	// of hitting the server
	serverHost  = os.Getenv("serverHost")
	serverPort  = os.Getenv("serverPort")
	serverRoute = "/send_random_leetcode_email"

	// utc time of sending emails every day
	timeOfSendingEmail = "15:15"
)

func main() {
	// schedule daily email
	go scheduledDailyEmail()

	// serve http for adhoc email
	fmt.Printf("serving at %s:%s%s\n", serverHost, serverPort, serverRoute)
	http.HandleFunc(serverRoute, func(w http.ResponseWriter, r *http.Request) {
		sendRandomLeetcodeEmail()
	})
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
}

func sendRandomLeetcodeEmail() {
	l, _ := NewRandomLeetcodePage()
	SendMail(l.ToHTML(serverHost, serverPort))
	fmt.Println("Email sent!")
}

func scheduledDailyEmail() {
	gocron.Every(1).Day().At(timeOfSendingEmail).Do(sendRandomLeetcodeEmail)
	<-gocron.Start()
}
