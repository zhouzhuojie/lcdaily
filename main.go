package main

import (
	"fmt"
	"net/http"

	"github.com/jasonlvhit/gocron"
)

func main() {
	// schedule
	go scheduledDailyEmail()

	// serve http
	fmt.Println("serving at port 8080")
	http.HandleFunc("/send_random_leetcode_email", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	sendRandomLeetcodeEmail()
}

func sendRandomLeetcodeEmail() {
	l, _ := NewRandomLeetcodePage()
	SendMail(l.ToHTML())
	fmt.Println("Email sent!")
}

func scheduledDailyEmail() {
	gocron.Every(1).Day().At("15:15").Do(sendRandomLeetcodeEmail)
	<-gocron.Start()
}
