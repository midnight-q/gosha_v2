package filesystem

import (
	"fmt"
	"go/token"
	"gosha_v2/common"
	"gosha_v2/errors"
	"gosha_v2/services/utils"
	"gosha_v2/settings"
	"os"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

type Field struct {
	Name        string
	Type        string
	Comment     string
	SourceModel string
}

type Model struct {
	Name              string
	Path              string
	Comment           string
	IsFilter          bool
	Fields            []Field
	IsServiceModel    bool
	CompositionModels []string
}

func LoadDbModels(path string) (res []Model, err error) {
	dirPath := path + "/dbmodels"
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	rawRes := []Model{}

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
					model.IsServiceModel = true
				}

				for _, field := range structType.Fields.List {
					if len(field.Names) < 1 {
						model.CompositionModels = append(model.CompositionModels, utils.ParseType(field.Type))
						continue
					}
					if field.Names[0].Name[:1] == strings.ToLower(field.Names[0].Name[:1]) {
						continue
					}
					model.Fields = append(model.Fields, Field{
						Name:    field.Names[0].Name,
						Type:    utils.ParseType(field.Type),
						Comment: utils.ParseComment(field.Decorations().Start.All()),
					})
				}

				rawRes = append(rawRes, model)
			}
		}
	}
	return LoadCompositionFields(rawRes)
}

func LoadTypeModels(path string) (res []Model, err error) {
	dirPath := path + "/types"
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	rawRes := []Model{}

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

				if common.CheckInArray(model.Name, settings.ServiceModelNames) {
					model.IsServiceModel = true
				}

				for _, field := range structType.Fields.List {
					if len(field.Names) < 1 {
						model.CompositionModels = append(model.CompositionModels, utils.ParseType(field.Type))
						continue
					}
					if field.Names[0].Name[:1] == strings.ToLower(field.Names[0].Name[:1]) {
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
				rawRes = append(rawRes, model)
			}
		}
	}

	return LoadCompositionFields(rawRes)
}

func LoadCompositionFields(models []Model) (res []Model, err error) {
	for _, model := range models {
		if len(model.CompositionModels) > 0 {
			for _, compositionModelName := range model.CompositionModels {
				fields, err := loadFields(compositionModelName, models, "")
				if err != nil {
					return nil, err
				}
				model.Fields = append(model.Fields, fields...)
			}
		}
		res = append(res, model)
	}
	return
}

func loadFields(modelName string, models []Model, prefix string) (res []Field, err error) {
	model, isFind := FindModel(modelName, models)
	if !isFind {
		return nil, errors.New(fmt.Sprintf("Not found composition model name=%s", modelName))
	}
	for _, field := range model.Fields {
		field.SourceModel = utils.CreateModelPath(prefix, modelName)
		res = append(res, field)
	}
	for _, name := range model.CompositionModels {
		fields, err := loadFields(name, models, utils.CreateModelPath(prefix, modelName))
		if err != nil {
			return nil, err
		}
		res = append(res, fields...)
	}

	return
}

func FindModel(name string, models []Model) (Model, bool) {
	for _, model := range models {
		if model.Name == name {
			return model, true
		}
	}
	return Model{}, false
}
