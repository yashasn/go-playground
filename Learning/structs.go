package main

import (
	"fmt"
)

// 3 types of structs - Nested, Embeded and Anonymous

type name struct {
	firstName string
	lastName  string
}

//name is embedded into User-  the types of name can be directly accessed

type User struct {
	name
	number int
}

// media is anonymous struct
type SendMessage struct {
	message    string
	sender     User
	receipient User
	media      struct {
		format string
		size   int
	}
}

func structFunction() {

	firstMessage := SendMessage{
		message: "Hi Yashas",
		sender: User{
			name: name{
				firstName: "James",
				lastName:  "Low",
			},
			number: 8833,
		},
		receipient: User{
			name: name{
				firstName: "Yashas",
				lastName:  "N",
			},
			number: 8843,
		},
		media: struct {
			format string
			size   int
		}{
			format: "text",
			size:   6,
		},
	}

	fmt.Printf("Message is %s  from %s to %s and size is %d. \n",
		firstMessage.message, firstMessage.sender.firstName, firstMessage.receipient.firstName, firstMessage.media.size)
}
