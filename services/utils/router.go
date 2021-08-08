package utils

import (
	"bytes"
	"gosha_v2/types"
	"strings"
)

func GenerateRoutesForModels(models []types.Model) (res []types.Model, err error) {
	for _, model := range models {
		if model.IsServiceModel {
			res = append(res, model)
			continue
		}
		baseRoute := GenerateRouteForModel(model.Name)

		model.HttpRoutes.Find = baseRoute
		model.HttpRoutes.Create = baseRoute
		model.HttpRoutes.MultiCreate = baseRoute + "/list"
		model.HttpRoutes.Read = baseRoute + "/{id}"
		model.HttpRoutes.Update = baseRoute + "/{id}"
		model.HttpRoutes.MultiUpdate = baseRoute + "/list"
		model.HttpRoutes.Delete = baseRoute + "/{id}"
		model.HttpRoutes.MultiDelete = baseRoute + "/list"
		model.HttpRoutes.FindOrCreate = baseRoute
		model.HttpRoutes.UpdateOrCreate = baseRoute

		res = append(res, model)
	}
	return
}

func GenerateRouteForModel(name string) string {
	return "/api/v1/" + GetFirstLowerCase(name)
}

func GetFirstLowerCase(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
