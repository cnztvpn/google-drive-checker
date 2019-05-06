package checker

import (
	"github.com/whywaita/google-drive-checker/files"
)

func DuplicateDirName(fs []*files.Files) map[string][]string {
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
		return detected
	}

	return nil
}
