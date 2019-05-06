package checker

import (
	"errors"
	"fmt"

	"github.com/whywaita/google-drive-checker/files"
)

func DuplicateDirName(fs []*files.Files) error {
	// check same name and diff id (= is invalid)
	nameToId := map[string]string{}
	detected := map[string][]string{}

	for _, f := range fs {
		if _, ok := nameToId[f.Name]; ok {
			// detect duplicate folder name
			detected[f.Name] = append(detected[f.Name], f.Id)
		} else {
			// new value
			nameToId[f.Name] = f.Id
		}

	}

	if len(detected) != 0 {
		fmt.Println("detected!")
		for folderName, ids := range detected {
			fmt.Printf("%v: %v\n", folderName, ids)
		}

		return errors.New("detected")
	}

	return nil
}
