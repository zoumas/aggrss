// Package env provides a Env struct that groups all the environment variables the service depends on
package env

import (
	"fmt"
	"os"
)

// Env groups all the environment variables the service depends on
type Env struct {
	// The port for the Web Server to bind to
	Port string
}

func Load() (*Env, error) {
	port, set := os.LookupEnv("PORT")
	if !set {
		return nil, envNotSet("PORT")
	}

	return &Env{
		Port: port,
	}, nil
}

func envNotSet(name string) error {
	return fmt.Errorf("%s environment variable is not set", name)
}
