package utils

import (
	"go/token"
	"gosha_v2/settings"

	"github.com/dave/dst"
)

func AddImportIfNeeded(name string, file *dst.File) {
	AddImportIfNotExist(file, GetImport(name))
}

func GetImport(name string) *dst.ImportSpec {
	name = prepareImportName(name)
	return &dst.ImportSpec{
		Path: &dst.BasicLit{
			Kind:  token.STRING,
			Value: WrapString(name),
		},
	}
}

func prepareImportName(name string) string {
	importPath, isExist := settings.ImportMap[name]
	if isExist {
		return importPath
	}
	return name
}

func AddImportIfNotExist(file *dst.File, importSpec *dst.ImportSpec) {
	if CheckImport(importSpec.Path.Value, file) {
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

		genDecl.Specs = append(genDecl.Specs, importSpec)
		isAdded = true
	}

	if !isAdded {
		file.Decls = append([]dst.Decl{&dst.GenDecl{
			Tok:   token.IMPORT,
			Specs: []dst.Spec{importSpec},
		}}, file.Decls...)
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
