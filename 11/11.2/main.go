package main

import "strings"

func removeProfanity(message *string) {
	if strings.Contains(*message, "fubb") {
		*message = strings.Replace(*message, "fubb", "****", 1)
	}

	if strings.Contains(*message, "shiz") {
		*message = strings.Replace(*message, "shiz", "****", 1)
	}

	if strings.Contains(*message, "witch") {
		*message = strings.Replace(*message, "witch", "*****", 1)
	}
}
