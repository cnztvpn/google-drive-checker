package files

import (
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
)

var (
	CreatedTime = "2019-01-01T12:00:00"
)

func GetAllDirList(srv *drive.Service, parent string) (dirs []*Files) {
	// get all directory in parent dir
	n := GetDirList(srv, &dirs, parent, "")

	for n != "" {
		n = GetDirList(srv, &dirs, parent, n)

	}

	return dirs
}

func GetDirList(srv *drive.Service, dirs *[]*Files, parent, npt string) (nextPageToken string) {
	dirQuery := fmt.Sprintf("(parents = '%s') and (trashed = false) and createdTime > '%s'", parent, CreatedTime)

	r, err := srv.Files.List().Q(dirQuery).Fields("nextPageToken, files(id, name, mimeType)").PageToken(npt).Do()
	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range r.Files {
		*dirs = append(*dirs, &Files{*folder, folder.Name})
	}

	if r.NextPageToken != "" {
		return r.NextPageToken
	}

	return ""
}
