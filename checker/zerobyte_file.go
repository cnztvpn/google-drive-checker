package checker

import (
	"errors"
	"fmt"

	"github.com/whywaita/google-drive-checker/files"
)

func ZeroByteFile(f *files.Files) error {
	if f.Size == 0 {
		err := fmt.Sprintf("Zerobyte file detect: %s", f.FullPath)
		return errors.New(err)
	}

	return nil
}
