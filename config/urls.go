package config

import (
	"PowerShare/helper"
)

// ports
const DefaultPortStr = ":5000"
var Port int = 5000

func PortStr() string{
	port, err := helper.GetPortString(Port)
	if err != nil {
		return DefaultPortStr
	}
	return port
}
