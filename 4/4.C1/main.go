package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType membershipType) User {
	user := User{}

	user.Name = name
	user.Type = membershipType

	if membershipType == TypePremium {
		user.MessageCharLimit = 1000
	} else {
		user.MessageCharLimit = 100
	}

	return user
}

func (user User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= user.MessageCharLimit {
		return message, true
	}
	return "", false
}
