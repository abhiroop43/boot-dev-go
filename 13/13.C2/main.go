package main

import "sync"

func processMessages(messages []string) []string {
	var wg sync.WaitGroup
	processedMessagesChan := make(chan string, len(messages))

	for _, message := range messages {
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
			processedMessagesChan <- msg + "-processed"
		}(message)
	}

	go func() {
		wg.Wait()
		close(processedMessagesChan)
	}()

	var processedMessages []string
	for msg := range processedMessagesChan {
		processedMessages = append(processedMessages, msg)
	}

	return processedMessages
}
