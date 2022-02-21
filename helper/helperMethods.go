package helper

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// GetPortString converts the port specified as int to a string and validates it in the process
func GetPortString(port int) (string, error) {
	var err error
	var portStr string
	if port<0 || port>65535 {
		err = fmt.Errorf("invalid port %d; must be between 0 an 65535", port)
	} else{
		portStr = ":"+strconv.Itoa(port)
	}

	return portStr, err
}

// CreateDir creates the specified directory with permissions 0755
func CreateDir(path string) error{
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			return errDir
		}
	}

	return nil
}

// Exists returns if the specified directory exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
