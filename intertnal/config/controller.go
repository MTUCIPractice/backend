package config

import "fmt"

type Controller struct {
	BindAddress string
	BindPort    int
}

func (c Controller) GetBindAddress() string {
	return fmt.Sprintf("%s:%d", c.BindAddress, c.BindPort)
}
