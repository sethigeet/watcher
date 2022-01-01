package cmd

import (
	"errors"
	"os"
)

// Exists returns a bool and an error that states whether the given
// file/directory exists or not
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err
}
