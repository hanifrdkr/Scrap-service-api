package util

import "strings"

func SplitFullname(fullname string) (first, middle, last string) {
	parts := strings.Split(fullname, " ")

	var firstName, middleName, lastName string

	if len(parts) <= 0 {
		firstName = ""
		middleName = ""
		lastName = ""
	}

	if len(parts) > 0 {
		firstName = parts[0]
	}

	if len(parts) > 1 {
		middleName = parts[1]
	}

	if len(parts) > 2 {
		lastName = strings.Join(parts[2:], " ")
	}

	return firstName, middleName, lastName
}
