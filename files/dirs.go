package files

import (
	"fmt"

	"google.golang.org/api/drive/v3"
)

var (
	CreatedTime = "2019-01-01T12:00:00"
)

func GetAllDirList(srv *drive.Service, parent string) (dirs []*Files, err error) {
	// get all directory in parent dir
	n, err := GetDirList(srv, &dirs, parent, "")
	if err != nil {
		return nil, err
	}

	for n != "" {
		n, err = GetDirList(srv, &dirs, parent, n)
		if err != nil {
			return nil, err
		}

	}

	return dirs, err
}

func GetDirList(srv *drive.Service, dirs *[]*Files, parent, npt string) (nextPageToken string, err error) {
	dirQuery := fmt.Sprintf("(parents = '%s') and (trashed = false) and createdTime > '%s'", parent, CreatedTime)

	r, err := srv.Files.List().Q(dirQuery).Fields("nextPageToken, files(id, name, mimeType)").PageToken(npt).Do()
	if err != nil {
		return "", err
	}

	for _, folder := range r.Files {
		*dirs = append(*dirs, &Files{*folder, folder.Name})
	}

	if r.NextPageToken != "" {
		return r.NextPageToken, nil
	}

	return "", nil
}
