package pg_core_validators

import (
	"errors"
	"freecreate/lib/logger"
)

func ValidateEmail(email string) error {
	if email == "" {
		err := errors.New("Email cannot be empty!")
		logger.Log(err)
		return err
	}

	return nil
}
