package filesystem

import (
	"fmt"
	"gosha_v2/settings"
	"gosha_v2/utils"
	"io/ioutil"
	"os"
	"strings"
)

func CreateGoMod(path string) error {
	skeletonGoModPath := path + settings.SkeletonGoModFilename
	newGoModPath := path + settings.GoModFilename

	if _, err := os.Stat(newGoModPath); err == nil {
		return fmt.Errorf("go.mod already exist")
	}

	b, err := os.ReadFile(skeletonGoModPath)
	if err != nil {
		return err
	}

	result := strings.Replace(string(b), settings.SkeletonAppName, utils.GetNameForNewApp(path), -1)

	err = ioutil.WriteFile(newGoModPath, []byte(result), os.ModePerm)
	if err != nil {
		return err
	}

	return os.Remove(skeletonGoModPath)
}
