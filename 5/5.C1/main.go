package main

import "fmt"

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

func (plainText PlainText) Format() string {
	return plainText.message
}

type Bold struct {
	message string
}

func (bold Bold) Format() string {
	return fmt.Sprintf("**%s**", bold.message)
}

type Code struct {
	message string
}

func (code Code) Format() string {
	return fmt.Sprintf("`%s`", code.message)
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
