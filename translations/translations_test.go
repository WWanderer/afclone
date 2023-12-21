package translations

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"testing"
)

func Test_setupTranslations(t *testing.T) {
	var doF = func(path string, info fs.DirEntry, err error) error {

		// first thing to do, check error. and decide what to do about it
		if err != nil {
			fmt.Printf("error [%v] at a path [%q]\n", err, path)
			return err
		}

		fmt.Printf("path: %v\n", path)

		// find out if it's a dir or file, if file, print info
		if info.IsDir() {
			fmt.Printf("is dir.\n")
		} else {
			fmt.Printf("  dir: [%v]\n", filepath.Dir(path))
			fmt.Printf("  file name [%v]\n", info.Name())
			fmt.Printf("  extenion: [%v]\n", filepath.Ext(path))
		}

		return nil
	}
	dir := "../../resources/translations"
	err := filepath.WalkDir(dir, doF)

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
	}
}
