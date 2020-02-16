package files

import "google.golang.org/api/drive/v3"

func Delete(srv *drive.Service, f *Files) error {
	return srv.Files.Delete(f.Id).Do()
}
