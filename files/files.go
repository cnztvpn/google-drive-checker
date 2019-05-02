package files

import (
	"fmt"

	"golang.org/x/sync/errgroup"

	"google.golang.org/api/drive/v3"
)

func getFileList(srv *drive.Service, resultFiles *[]*drive.File, dirs []*drive.File) error {
	limit := make(chan struct{}, 3) // limit of parallel Google Drive API call

	eg := errgroup.Group{}
	for _, dir := range dirs {
		eg.Go(func() error {
			limit <- struct{}{}

			dirQuery := fmt.Sprintf("(parents = '%s') and (trashed = false)", dir.Id)
			// 一つのアニメディレクトリは100個以下のはず = no pagination
			c := srv.Files.List().Q(dirQuery).PageSize(100).Fields("nextPageToken, files(id, name, mimeType, size)")

			r, err := c.Do()
			if err != nil {
				return err
			}

			for _, f := range r.Files {
				// 一つのアニメディレクトリにはもうディレクトリはないはず
				*resultFiles = append(*resultFiles, f)
			}

			<-limit

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func GetFileListById(srv *drive.Service, resultFiles *[]*drive.File, parent string) error {
	dirs := GetAllDirList(srv, parent)

	err := getFileList(srv, resultFiles, dirs)
	if err != nil {
		return err
	}

	return nil
}

func GetFileListByDirs(srv *drive.Service, resultFiles *[]*drive.File, dirs []*drive.File) error {
	return getFileList(srv, resultFiles, dirs)
}
