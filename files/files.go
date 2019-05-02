package files

import (
	"fmt"
	"log"
	"sync"

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

func GetFileList(srv *drive.Service, files *[]*drive.File, parent string) {
	dirs := GetAllDirList(srv, parent)
	limit := make(chan struct{}, 3)

	var wg sync.WaitGroup
	for _, dir := range dirs {
		wg.Add(1)
		go func(dir *drive.File) {
			limit <- struct{}{}
			defer wg.Done()

			dirQuery := fmt.Sprintf("(parents = '%s') and (trashed = false)", dir.Id)
			// 一つのアニメディレクトリは100個以下のはず
			c := srv.Files.List().Q(dirQuery).PageSize(100).Fields("nextPageToken, files(id, name, mimeType, size)")

			r, err := c.Do()
			if err != nil {
				log.Fatal(err)
			}

			for _, f := range r.Files {
				// 一つのアニメディレクトリにはもうディレクトリはないはず
				*files = append(*files, f)
			}

			<-limit
		}(dir)
	}

	wg.Wait()
}
