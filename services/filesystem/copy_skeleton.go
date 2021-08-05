package filesystem

import (
	"fmt"
	"gosha_v2/settings"
	"io"
	"io/fs"
	"os"
	"strings"
)

func CopySkeletonApp(workDir string) (err error) {
	dirs, files, err := getDirAndFileLists(settings.SkeletonAppPath, templateFS)
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

func getDirAndFileLists(path string, fs fs.ReadDirFS) (dirsRes []string, filesRes []string, err error) {

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
			dirs, files, err := getDirAndFileLists(path+"/"+entry.Name(), fs)
			if err != nil {
				return nil, nil, err
			}
			dirsRes = append(dirsRes, dirs...)
			filesRes = append(filesRes, files...)
		} else {
			filesRes = append(filesRes, path+"/"+entry.Name())
		}
	}
	return
}

func getNewFileName(path, workDir string) string {
	return workDir + strings.TrimPrefix(path, settings.SkeletonAppPath)
}
