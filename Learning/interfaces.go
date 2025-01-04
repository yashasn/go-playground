package main

import (
	"fmt"
	"time"
)

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	name string
	time time.Time
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, your birthday is on %s", bm.name, bm.time)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

// Interface casting using switch
func sendMessage(msg message) {
	switch v := msg.(type) {
	case sendingReport:
		fmt.Println("It's an email-report")
		fmt.Println(v.getMessage())
	case birthdayMessage:
		fmt.Println("It's a birthday message")
		fmt.Println(v.getMessage())
	default:
		fmt.Println("Unresolved Type")
	}

}

func interfaceTest() {
	sendMessage(birthdayMessage{
		name: "John Doe",
		time: time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})

	sendMessage(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
}
