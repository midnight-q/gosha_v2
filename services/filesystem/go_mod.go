package filesystem

import (
	"fmt"
	"gosha_v2/services/utils"
	"gosha_v2/settings"
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

func FindAppName(path string) (name string) {
	goModPath := path + settings.GoModFilename

	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		return ""
	}

	b, err := os.ReadFile(goModPath)
	if err != nil {
		return ""
	}
	file := string(b)

	lines := strings.Split(file, "\n")

	rawName := strings.Replace(lines[0], "module", "", 1)

	name = strings.TrimSpace(rawName)

	return
}
