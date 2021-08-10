package utils

import (
	"gosha_v2/errors"
	"gosha_v2/types"
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
