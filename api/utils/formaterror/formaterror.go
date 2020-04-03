package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "username") {
		return errors.New("Username Already Taken")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Invalid Password")
	}
	return errors.New("Error occured")
}
