package helper

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPortString(t *testing.T) {
	cases := []struct {
		input int
		expectedPort string
		expectedError error
	}{
		{8080, ":8080", nil},
		{443, ":443", nil},

		{-1, "", errors.New("invalid port -1; must be between 0 an 65535")},
		{65536, "", errors.New("invalid port 65536; must be between 0 an 65535")},

	}

	for _, c := range cases {
		actualPort, actualError := GetPortString(c.input)

		assert.Equal(t, c.expectedPort, actualPort, "wrong conversion")
		assert.Equal(t, c.expectedError, actualError, "wrong error")
	}
}
