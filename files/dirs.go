package files

import (
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
)

func GetAllDirList(srv *drive.Service, parent string) (dirs []*drive.File) {
	// get all directory in parent dir
	dirQuery := fmt.Sprintf("(parents = '%s') and (trashed = false)", parent)

	r, err := srv.Files.List().Q(dirQuery).Fields("nextPageToken, files(id, name, mimeType)").Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range r.Files {
		dirs = append(dirs, f)
	}

	if r.NextPageToken != "" {
		n := GetDirList(srv, &dirs, parent, r.NextPageToken)

		for n != "" {
			n = GetDirList(srv, &dirs, parent, n)
		}
	}

	return dirs
}

func GetDirList(srv *drive.Service, dirs *[]*drive.File, parent, npt string) (nextPageToken string) {
	dirQuery := fmt.Sprintf("(parents = '%s') and (trashed = false)", parent)

	r, err := srv.Files.List().Q(dirQuery).Fields("nextPageToken, files(id, name, mimeType)").PageToken(npt).Do()
	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range r.Files {
		*dirs = append(*dirs, folder)
	}

	if r.NextPageToken != "" {
		return r.NextPageToken
	}

	return ""
}
