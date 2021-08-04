package utils

import (
	"go/token"

	"github.com/dave/dst"
)

func GetUuidImport() *dst.ImportSpec {
	return &dst.ImportSpec{
		Path: &dst.BasicLit{
			Kind:  token.STRING,
			Value: WrapString("github.com/google/uuid"),
		},
	}
}

func AddUuidImportIfNotExist(file *dst.File) {
	if CheckImport("github.com/google/uuid", file) {
		return
	}
	isAdded := false
	for _, decl := range file.Decls {
		genDecl, isOk := decl.(*dst.GenDecl)
		if !isOk {
			continue
		}
		if genDecl.Tok != token.IMPORT {
			continue
		}

		genDecl.Specs = append(genDecl.Specs, GetUuidImport())
		isAdded = true
	}

	if !isAdded {
		//file.Decls = append(file.Decls, &dst.GenDecl{
		//	Tok:    token.IMPORT,
		//	Specs:  []dst.Spec{GetUuidImport()},
		//})
	}
}

func CheckImport(path string, file *dst.File) bool {
	for _, decl := range file.Decls {
		genDecl, isOk := decl.(*dst.GenDecl)
		if !isOk {
			continue
		}
		if genDecl.Tok != token.IMPORT {
			continue
		}
		for _, spec := range genDecl.Specs {
			importSpec, isOk := spec.(*dst.ImportSpec)
			if !isOk {
				continue
			}
			if importSpec.Path.Value == path {
				return true
			}
		}
	}
	return false
}
