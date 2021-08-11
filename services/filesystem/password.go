package filesystem

import (
	"fmt"
	"gosha_v2/services/utils"

	"github.com/dave/dst"
)

func UpdatePasswordSalt(currentPath, newSalt string) error {
	fileName := currentPath + "/settings/user.go"
	file, err := readFile(fileName)
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

	return saveFile(file, fileName)
}
