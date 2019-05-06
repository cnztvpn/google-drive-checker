package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	ParentId         string
	CredJson         string
	SlackHookURL     string
	SlackChannelName string
}

func Initiallize() (*Config, error) {
	c := Config{}
	piEnv := "GD_CHECKER_PARENT_ID"
	cjEnv := "GD_CHECKER_CRED_JSON"
	shuEnv := "GD_CHECKER_SLACK_HOOK_URL"
	shcnEnv := "GD_CHECKER_SLACK_CHANNEL_NAME"

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

	if os.Getenv(shuEnv) == "" {
		return nil, errors.New(fmt.Sprintf("%s must be set", shuEnv))
	} else {
		c.SlackHookURL = os.Getenv(shuEnv)
	}

	if os.Getenv(shcnEnv) == "" {
		return nil, errors.New(fmt.Sprintf("%s must be set", shcnEnv))
	} else {
		c.SlackChannelName = os.Getenv(shcnEnv)
	}

	return &c, nil
}
