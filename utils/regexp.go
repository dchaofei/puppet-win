package utils

import "regexp"

var roomIDRegexp = regexp.MustCompile("@chatroom$")

func IsRoomID(id string) bool {
	return roomIDRegexp.MatchString(id)
}

var openIMRegexp = regexp.MustCompile("@openim$")

func IsIMContactId(id string) bool {
	return openIMRegexp.MatchString(id)
}
