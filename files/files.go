package files

import (
	"fmt"

	"golang.org/x/sync/errgroup"
	"google.golang.org/api/drive/v3"
)

type Files struct {
	// expand drive.File
	drive.File

	FullPath string
}

func getFileList(srv *drive.Service, resultFiles *[]*drive.File, dirs []*drive.File) error {
	limit := make(chan struct{}, 1) // limit of parallel Google Drive API call

	eg := errgroup.Group{}
	for _, dir := range dirs {
		d := dir

		eg.Go(func() error {
			limit <- struct{}{}
			defer func() {
				<-limit
			}()

			dirQuery := fmt.Sprintf("(parents = '%s') and (trashed = false) and createdTime > '%s'", d.Id, CreatedTime)
			// 一つのアニメディレクトリは20個以下のはず = no pagination
			c := srv.Files.List().Q(dirQuery).PageSize(20).Fields("nextPageToken, files(id, name, mimeType, size)")

			r, err := c.Do()
			if err != nil {
				return err
			}

			for _, f := range r.Files {
				// 一つのアニメディレクトリにはもうディレクトリはないはず
				*resultFiles = append(*resultFiles, f)
			}

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
