package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	primaryLength := len(primary)
	secondaryLength := len(secondary)
	tertiaryLength := len(tertiary)

	costs := [3]int{primaryLength, primaryLength + secondaryLength, primaryLength + secondaryLength + tertiaryLength}

	return messages, costs

}
