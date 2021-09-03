package types

import (
	"net/http"
	"skeleton-app/settings"
)

type NewModel struct {
	Id int

	validator
}

func (model *NewModel) Validate() {
}

type NewModelFilter struct {
	model NewModel
	list  []NewModel

	AbstractFilter
}

func GetNewModelFilter(request *http.Request, functionType string) (filter NewModelFilter, err error) {

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

func (filter *NewModelFilter) GetNewModelModel() NewModel {

	filter.model.Validate()

	return filter.model
}

func (filter *NewModelFilter) GetNewModelModelList() (data []NewModel, err error) {

	for i := range filter.list {
		filter.list[i].Validate()

		if !filter.list[i].IsValid() {
			err = filter.list[i].GetValidationError()
			return nil, err
		}
	}

	return filter.list, nil
}

func (filter *NewModelFilter) SetNewModelModel(typeModel NewModel) {

	filter.model = typeModel
}

func (filter *NewModelFilter) SetNewModelModelList(data []NewModel) {

	filter.list = data
}
