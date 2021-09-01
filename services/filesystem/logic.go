package filesystem

import (
	"go/token"
	"gosha_v2/services/utils"
	"strings"

	"github.com/dave/dst"
)

func ClearLogic(currentPath, modelName string) (err error) {
	path := currentPath + "/logic/" + utils.GetFileNameFromModelName(modelName) + ".go"
	file, err := readFile(path)
	if err != nil {
		return err
	}

	var newBody []dst.Decl

	for _, decl := range file.Decls {

		funcDecl, isOk := decl.(*dst.FuncDecl)
		if !isOk {
			genDecl, isOk := decl.(*dst.GenDecl)
			if !isOk {
				newBody = append(newBody, decl)
				continue
			}
			if genDecl.Tok != token.IMPORT {
				newBody = append(newBody, decl)
				continue
			}
			var newSpec = []dst.Spec{}
			for _, spec := range genDecl.Specs {
				importSpec, isOk := spec.(*dst.ImportSpec)
				if !isOk {
					continue
				}
				if strings.Contains(importSpec.Path.Value, "/types") || strings.Contains(importSpec.Path.Value, "/gorm") {
					newSpec = append(newSpec, importSpec)
				}
			}
			genDecl.Specs = newSpec
			newBody = append(newBody, genDecl)
			continue
		}
		if strings.HasPrefix(funcDecl.Name.Name, "Assign") {
			continue
		}

		funcDecl.Body.List = []dst.Stmt{&dst.ReturnStmt{
			Decs: dst.ReturnStmtDecorations{
				NodeDecs: utils.GetNewLineDecorations(),
			},
		}}
		newBody = append(newBody, funcDecl)
	}

	file.Decls = newBody

	return saveFile(file, path)
}
