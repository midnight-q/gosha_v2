package utils

import (
	"gosha_v2/errors"
	"gosha_v2/types"
	"regexp"
	"strings"
)

func CheckIsFilter(name string) bool {
	return strings.Contains(name, "Filter")
}

func IsIdField(filedName string) bool {
	if strings.ToLower(filedName) == "id" {
		return true
	}
	return false
}

func FilterServiceModels(models []types.Model) (res []types.Model) {

	for _, model := range models {
		if !model.IsServiceModel {
			res = append(res, model)
		}
	}

	return
}

func CreateModelPath(prefix, modelName string) string {
	if len(prefix) < 1 {
		return modelName
	}
	return prefix + "." + modelName
}

func GetModel(name string, models []types.Model) (res types.Model, IsFilter bool, err error) {
	for _, model := range models {
		if model.Name == name {
			return model, false, nil
		}
		if model.Filter != nil && model.Filter.Name == name {
			return model, true, nil
		}
	}

	return types.Model{}, false, errors.New("Not found model with name: " + name)
}

func GetFileNameFromModelName(name string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func IsFieldExistInModel(fieldName string, model types.Model, isFilter bool) bool {
	if isFilter {
		for _, field := range model.Filter.Fields {
			if field.Name == fieldName {
				return true
			}
		}
	} else {
		for _, field := range model.Fields {
			if field.Name == fieldName {
				return true
			}
		}
	}
	return false
}

func GetFilePath(currentPath, dir, modelName string) string {
	return currentPath + dir + GetFileNameFromModelName(modelName) + ".go"
}
