package filesystem

import (
	"fmt"
	"gosha_v2/services/utils"
	"gosha_v2/settings"
	"io"
	"io/fs"
	"os"
	"strings"
)

func CopySkeletonApp(workDir string) (err error) {
	dirs, files, err := getDirAndFileLists(settings.SkeletonAppPath, templateFS, true)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		newDir := getNewFileName(dir, workDir)
		if _, err := os.Stat(newDir); os.IsNotExist(err) {
			err := os.Mkdir(newDir, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	for _, file := range files {
		err = copyFile(file, getNewFileName(file, workDir))
		if err != nil {
			return err
		}
	}

	return nil
}

func copyFile(src string, dst string) error {
	source, err := templateFS.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("file %s already exists", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	buf := make([]byte, settings.BufferSizeForCopy)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}

func getDirAndFileLists(path string, fs fs.ReadDirFS, isIgnoreNewModel bool) (dirsRes []string, filesRes []string, err error) {

	var entries []os.DirEntry
	if fs == nil {
		entries, err = os.ReadDir(path)
	} else {
		entries, err = fs.ReadDir(path)
	}

	if err != nil {
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			dirsRes = append(dirsRes, path+"/"+entry.Name())
			dirs, files, err := getDirAndFileLists(path+"/"+entry.Name(), fs, isIgnoreNewModel)
			if err != nil {
				return nil, nil, err
			}
			dirsRes = append(dirsRes, dirs...)
			filesRes = append(filesRes, files...)
		} else {
			if isIgnoreNewModel && strings.Contains(entry.Name(), "new_model") {
				continue
			}
			filesRes = append(filesRes, path+"/"+entry.Name())
		}
	}
	return
}

func getNewFileName(path, workDir string) string {
	return workDir + strings.TrimPrefix(path, settings.SkeletonAppPath)
}

func CopyNewModelFile(currentPath, dir, newName, appName string, withSoftDelete bool) (err error) {
	modelData := GetBaseModelData(withSoftDelete)

	source, err := templateFS.ReadFile(settings.SkeletonAppPath + dir + modelData.FileName)
	if err != nil {
		return err
	}
	text := string(source)
	text = strings.Replace(text, modelData.ModelName, newName, -1)

	text = strings.Replace(text, settings.SkeletonAppName, appName, -1)

	return os.WriteFile(utils.GetFilePath(currentPath, dir, newName), []byte(text), fs.ModePerm)
}

type ModelData struct {
	FileName  string
	ModelName string
}

func GetBaseModelData(withSoftDelete bool) (res ModelData) {
	if withSoftDelete {
		return ModelData{
			FileName:  settings.NewModelWithSoftDeleteFileName,
			ModelName: settings.NewModelWithSoftDeleteName,
		}
	}
	return ModelData{
		FileName:  settings.NewModelFileName,
		ModelName: settings.NewModelName,
	}
}
