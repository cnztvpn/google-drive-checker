package main

import (
	"log"
	"os"
	"testing"

	"github.com/whywaita/google-drive-checker/files"

	"github.com/whywaita/google-drive-checker/api"
	"github.com/whywaita/google-drive-checker/checker"
	"github.com/whywaita/google-drive-checker/config"
	"google.golang.org/api/drive/v3"
)

var (
	con = &config.Config{}
	srv = &drive.Service{}
	fs  = &[]*drive.File{}
)

func TestMain(m *testing.M) {
	before()
	os.Exit(m.Run())
}

func before() {
	c, err := config.Initiallize()
	if err != nil {
		log.Fatal(err)
	}
	client := api.GetClient(c.CredJson)

	s, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	err = files.GetFileListById(srv, fs, con.ParentId)
	if err != nil {
		log.Fatalf("Unable to get all file List: %v", err)
	}

	con = c
	srv = s
}

func TestZeroByteFile(t *testing.T) {
	for _, f := range *fs {
		err := checker.ZeroByteFile(f)
		if err != nil {
			// notify slack
			log.Println(err)
		}
	}
}
