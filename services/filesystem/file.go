package filesystem

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

func saveFile(file *dst.File, path string) error {
	buf := bytes.NewBuffer([]byte{})
	err := decorator.Fprint(buf, file)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, buf.Bytes(), fs.ModePerm)
}

func readFile(path string) (*dst.File, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return decorator.Parse(b)
}
