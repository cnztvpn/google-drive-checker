package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	ParentId string
	CredJson string
}

func Initiallize() (*Config, error) {
	c := Config{}
	piEnv := "GD_CHECKER_PARENT_ID"
	cjEnv := "GD_CHECKER_CRED_JSON"

	if os.Getenv(piEnv) == "" {
		return nil, errors.New(fmt.Sprintf("%s must be set", piEnv))
	} else {
		c.ParentId = os.Getenv(piEnv)
	}

	if os.Getenv(cjEnv) == "" {
		return nil, errors.New(fmt.Sprintf("%s must be set", cjEnv))
	} else {
		c.CredJson = os.Getenv(cjEnv)
	}

	return &c, nil
}
