package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (auth authenticationInfo) getBasicAuth() string {
	retVal := fmt.Sprintf("Authorization: Basic %s:%s", auth.username, auth.password)
	return retVal
}
