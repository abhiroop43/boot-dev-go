package main

import (
	"errors"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	nameLength := len(names)
	phoneNumbersLength := len(phoneNumbers)

	if nameLength != phoneNumbersLength {
		return nil, errors.New("invalid sizes")
	}

	userMap := make(map[string]user)

	for i := 0; i < nameLength; i++ {
		userMap[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
