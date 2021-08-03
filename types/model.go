package types

import (
	"gosha_v2/settings"
	"net/http"
)

type Model struct {
	Id   int
	Name string

	validator
}

func (model *Model) Validate() {
}

type ModelFilter struct {
	model Model
	list  []Model

	AbstractFilter
}

func GetModelFilter(request *http.Request, functionType string) (filter ModelFilter, err error) {

	filter.request = request
	filter.rawRequestBody, err = GetRawBodyContent(request)
	if err != nil {
		return filter, err
	}

	switch functionType {
	case settings.FunctionTypeMultiCreate, settings.FunctionTypeMultiUpdate, settings.FunctionTypeMultiDelete, settings.FunctionTypeMultiFindOrCreate:
		err = ReadJSON(filter.rawRequestBody, &filter.list)
		if err != nil {
			return
		}
		break
	default:
		err = ReadJSON(filter.rawRequestBody, &filter.model)
		if err != nil {
			return
		}
		break
	}

	filter.AbstractFilter, err = GetAbstractFilter(request, filter.rawRequestBody, functionType)

	return filter, err
}

func (filter *ModelFilter) GetModelModel() Model {

	filter.model.Validate()

	return filter.model
}

func (filter *ModelFilter) GetModelModelList() (data []Model, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *ModelFilter) SetModelModel(typeModel Model) {

	filter.model = typeModel
}

func (filter *ModelFilter) SetModelModelList(data []Model) {

	filter.list = data
}
