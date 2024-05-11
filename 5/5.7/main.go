package main

import (
	"fmt"
)

func (e email) cost() int {
	bodyLength := len(e.body)

	if e.isSubscribed {
		return bodyLength * 2
	}

	return bodyLength * 5
}

func (e email) format() string {
	var isSubscribed string

	if e.isSubscribed {
		isSubscribed = "Subscribed"
	} else {
		isSubscribed = "Not Subscribed"
	}

	formattedString := fmt.Sprintf("'%s' | %s", e.body, isSubscribed)

	return formattedString
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
