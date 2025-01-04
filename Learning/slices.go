package main

import (
	"errors"
	"fmt"
)

const (
	freePlan = "free"
	proPlan  = "pro"
)

func MessagesWithPlan(plan string) ([]string, error) {
	allMessages := Messages()
	if plan == proPlan {
		return allMessages[:], nil
	}
	if plan == freePlan {
		return allMessages[0:1], nil
	}
	return nil, errors.New("unsupported plan")
}

func Messages() [3]string {
	return [3]string{
		"First free message",
		"Second  message",
		"Third  message",
	}
}

//variadic functions

func PrintMessages(messages ...string) {
	for i := 0; i < len(messages); i++ {
		fmt.Printf("sending : %s .\n", messages[i])
	}
}

//variadic functions

func VariadicExample(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func TestSlices() {
	messages, err := MessagesWithPlan("PRO")
	if err != nil {
		fmt.Println(err)
		return
	}

	PrintMessages(messages...)

	//can pass any number of arguments
	VariadicExample(1, 2, 3, 4)
}
