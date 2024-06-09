package utils

import "regexp"

var roomIDRegexp = regexp.MustCompile("@chatroom$")

func IsRoomID(id string) bool {
	return roomIDRegexp.MatchString(id)
}
