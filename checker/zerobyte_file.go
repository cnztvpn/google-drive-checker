package checker

import (
	"errors"
	"fmt"

	"google.golang.org/api/drive/v3"
)

func ZeroByteFile(f *drive.File) error {
	if f.Size == 0 {
		err := fmt.Sprintf("Zerobyte file detect: %s", f.Name)
		return errors.New(err)
	}

	return nil
}
