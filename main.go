package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/whywaita/google-drive-checker/api"
	"github.com/whywaita/google-drive-checker/checker"
	"github.com/whywaita/google-drive-checker/files"
	"google.golang.org/api/drive/v3"
)

type Config struct {
	parentId string
	credJson string
}

func initiallize() (*Config, error) {
	c := Config{}
	piEnv := "GD_CHECKER_PARENT_ID"
	cjEnv := "GD_CHECKER_CRED_JSON"

	if os.Getenv(piEnv) == "" {
		return nil, errors.New(fmt.Sprintf("%s must be set", piEnv))
	} else {
		c.parentId = os.Getenv(piEnv)
	}

	if os.Getenv(cjEnv) == "" {
		return nil, errors.New(fmt.Sprintf("%s must be set", cjEnv))
	} else {
		c.credJson = os.Getenv(cjEnv)
	}

	return &c, nil
}

func main() {
	c, err := initiallize()
	if err != nil {
		log.Fatal(err)
	}
	client := api.GetClient(c.credJson)

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	fs := &[]*drive.File{}
	files.GetFileList(srv, fs, c.parentId)

	for _, f := range *fs {
		err := checker.ZerobyteFile(f)
		if err != nil {
			// notify slack
			log.Println(err)
		}
	}
}
