package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	var suggestedFriends []string

	// get user's friends
	userFriends := friendships[username]

	// loop through user's friends
	for _, friend := range userFriends {

		// get the friends list of currently selected friend
		friendsOfFriend := friendships[friend]

		for _, friendOfFriend := range friendsOfFriend {
			hisFriendships := friendships[friendOfFriend]

			// if the new friend is the same as original user, then skip
			if friendOfFriend == username {
				continue
			}

			// check if the username's friend is in the list and also if this friend is not already added
			if checkStringExists(hisFriendships, friend) && !checkStringExists(suggestedFriends, friendOfFriend) {
				suggestedFriends = append(suggestedFriends, friendOfFriend)
			}
		}
	}

	return suggestedFriends
}

func checkStringExists(collection []string, str string) bool {

	for _, value := range collection {
		if value == str {
			return true
		}
	}

	return false
}
