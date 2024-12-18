package listener

import (
	goloaderenv "github.com/ralvarezdev/go-loader/env"
	"strings"
)

// Port is the port of the server
type Port struct {
	Port string
	Host string
}

// LoadPort load port from environment variables
func LoadPort(host string, key string) (
	*Port, error,
) {
	// Get environment variable
	port, err := goloaderenv.LoadVariable(key)
	if err != nil {
		return nil, err
	}

	// Build host string
	var hostBuilder strings.Builder
	hostBuilder.WriteString(host)
	hostBuilder.WriteString(":")
	hostBuilder.WriteString(port)

	return &Port{
		Port: port,
		Host: hostBuilder.String(),
	}, nil
}
