package filesystem

import (
	"go/token"
	"gosha_v2/common"
	"gosha_v2/services/utils"
	"gosha_v2/settings"
	"os"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

type Field struct {
	Name    string
	Type    string
	Comment string
}

type Model struct {
	Name     string
	Path     string
	Comment  string
	IsFilter bool
	Fields   []Field
}

func LoadDbModels(path string) (res []Model, err error) {
	dirPath := path + "/dbmodels"
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		filePath := dirPath + "/" + entry.Name()
		b, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		file, err := decorator.Parse(b)

		for _, decl := range file.Decls {
			modelDecl, isOk := decl.(*dst.GenDecl)
			if !isOk {
				continue
			}

			if modelDecl.Tok == token.TYPE {
				typeSpec := modelDecl.Specs[0].(*dst.TypeSpec)
				structType := typeSpec.Type.(*dst.StructType)

				model := Model{
					Name:     typeSpec.Name.Name,
					Path:     filePath,
					Comment:  utils.ParseComment(modelDecl.Decorations().Start.All()),
					IsFilter: false,
					Fields:   []Field{},
				}

				if len(model.Name) < 1 {
					continue
				}

				if common.CheckInArray(model.Name, settings.ServiceModelNames) {
					continue
				}

				if model.Name[:1] == strings.ToLower(model.Name[:1]) {
					continue
				}

				for _, field := range structType.Fields.List {
					if len(field.Names) != 1 {
						continue
					}
					model.Fields = append(model.Fields, Field{
						Name:    field.Names[0].Name,
						Type:    utils.ParseType(field.Type),
						Comment: utils.ParseComment(field.Decorations().Start.All()),
					})
				}

				res = append(res, model)
			}
		}
	}
	return
}

func LoadTypeModels(path string) (res []Model, err error) {
	dirPath := path + "/types"
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		filePath := dirPath + "/" + entry.Name()
		b, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		file, err := decorator.Parse(b)

		for _, decl := range file.Decls {
			modelDecl, isOk := decl.(*dst.GenDecl)
			if !isOk {
				continue
			}

			if modelDecl.Tok == token.TYPE {
				typeSpec := modelDecl.Specs[0].(*dst.TypeSpec)
				structType := typeSpec.Type.(*dst.StructType)

				model := Model{
					Name:     typeSpec.Name.Name,
					Path:     filePath,
					Comment:  utils.ParseComment(modelDecl.Decorations().Start.All()),
					IsFilter: utils.CheckIsFilter(typeSpec.Name.Name),
					Fields:   []Field{},
				}

				if len(model.Name) < 1 {
					continue
				}

				if model.Name[:1] == strings.ToLower(model.Name[:1]) {
					continue
				}

				if common.CheckInArray(model.Name, settings.ServiceModelNames) {
					continue
				}

				for _, field := range structType.Fields.List {
					if len(field.Names) != 1 {
						continue
					}
					f := Field{
						Name:    field.Names[0].Name,
						Type:    utils.ParseType(field.Type),
						Comment: utils.ParseComment(field.Decorations().Start.All()),
					}
					if len(f.Name) < 1 {
						continue
					}

					if f.Name[:1] == strings.ToLower(f.Name[:1]) {
						continue
					}
					model.Fields = append(model.Fields, f)
				}
				res = append(res, model)
			}
		}
	}
	return
}
