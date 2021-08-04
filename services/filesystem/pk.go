package filesystem

import (
	"bytes"
	"go/token"
	"gosha_v2/utils"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

func ChangePKToUuid(currentPath string) error {

	entriesDbmodels, err := os.ReadDir(currentPath + "dbmodels/")
	if err != nil {
		return err
	}

	err = replacePkToUuid(entriesDbmodels, currentPath, "dbmodels/")
	if err != nil {
		return err
	}

	entriesTypes, err := os.ReadDir(currentPath + "types/")
	if err != nil {
		return err
	}

	err = replacePkToUuid(entriesTypes, currentPath, "types/")
	if err != nil {
		return err
	}
	return nil
}


func replacePkToUuid(entries []os.DirEntry, path, dir string) error {
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".go" {
			continue
		}
		fileName := path + dir +entry.Name()
		b, err := os.ReadFile(fileName)
		if err != nil {
			return err
		}

		file, err := decorator.Parse(b)
		if err != nil {
			return err
		}

		isReplace := false
		for _, decl := range file.Decls {
			modelDecl, isOk := decl.(*dst.GenDecl)
			if !isOk {
				continue
			}

			if modelDecl.Tok == token.TYPE {
				typeSpec := modelDecl.Specs[0].(*dst.TypeSpec)
				structType := typeSpec.Type.(*dst.StructType)
				for _, field := range structType.Fields.List {
					if len(field.Names) != 1 {
						continue
					}

					if strings.ToLower(field.Names[0].Name) == "id" {
						field.Type = utils.GetUuidType()
						isReplace = true
					}
				}
			}
		}

		if isReplace {
			utils.AddUuidImportIfNotExist(file)
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
	}

	return nil
}
