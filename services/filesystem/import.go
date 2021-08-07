package filesystem

import (
	"gosha_v2/services/utils"
	"gosha_v2/settings"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func FindAndReplaceImports(currentPath string) error {
	newAppName := utils.GetNameForNewApp(currentPath)

	_, files, err := getDirAndFileLists(currentPath, nil)
	if err != nil {
		return err
	}
	for _, file := range files {
		if filepath.Ext(file) != ".go" {
			continue
		}
		err = replaceImports(file, newAppName)
	}

	return nil
}

func replaceImports(file string, name string) error {
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	result := strings.Replace(string(b), settings.SkeletonAppName, name, -1)

	return ioutil.WriteFile(file, []byte(result), os.ModePerm)
}
