package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	var uniqueWords []string
	for _, message := range messages {
		wordsInMessage := getWords(message)
		for _, word := range wordsInMessage {
			if !wordExists(uniqueWords, word) {
				uniqueWords = append(uniqueWords, word)
			}
		}
	}

	return len(uniqueWords)
}

func getWords(sentence string) []string {
	return strings.Fields(sentence)
}

func wordExists(words []string, word string) bool {
	for _, w := range words {
		if strings.ToUpper(w) == strings.ToUpper(word) {
			return true
		}
	}

	return false
}
