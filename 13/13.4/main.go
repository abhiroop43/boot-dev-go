package main

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))

	for i, email := range emails {
		ch <- email
		if i == len(emails)-1 {
			close(ch)
		}
	}

	return ch
}
