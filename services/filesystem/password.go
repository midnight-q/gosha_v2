package filesystem

import (
	"bytes"
	"fmt"
	"gosha_v2/utils"
	"io/fs"
	"io/ioutil"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

func UpdatePasswordSalt(currentPath, newSalt string) error {
	fileName := currentPath + "/settings/user.go"
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	file, err := decorator.Parse(b)
	if err != nil {
		return err
	}

	isFind := false

	for _, decl := range file.Decls {
		constDecl, isOk := decl.(*dst.GenDecl)
		if !isOk {
			continue
		}
		if len(constDecl.Specs) != 1 {
			continue
		}

		spec, isOk := constDecl.Specs[0].(*dst.ValueSpec)
		if !isOk {
			continue
		}
		if len(spec.Names) != 1 {
			continue
		}
		switch spec.Names[0].Name {
		case "PASSWORD_SALT":
			if len(spec.Values) != 1 {
				continue
			}
			val, isOk := spec.Values[0].(*dst.BasicLit)
			if !isOk {
				continue
			}
			val.Value = utils.WrapString(newSalt)
			isFind = true
		default:
			continue
		}
	}
	if !isFind {
		return fmt.Errorf("not found PASSWORD_SALT")
	}

	buf := bytes.NewBuffer([]byte{})
	err = decorator.Fprint(buf, file)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, buf.Bytes(), fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
