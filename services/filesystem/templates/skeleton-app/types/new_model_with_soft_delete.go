package types

import (
	"net/http"
	"skeleton-app/settings"
)

type NewModelWithSoftDelete struct {
	Id int

	validator
}

func (model *NewModelWithSoftDelete) Validate() {
}

type NewModelWithSoftDeleteFilter struct {
	model NewModelWithSoftDelete
	list  []NewModelWithSoftDelete

	AbstractFilter
}

func GetNewModelWithSoftDeleteFilter(request *http.Request, functionType string) (filter NewModelWithSoftDeleteFilter, err error) {

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

func (filter *NewModelWithSoftDeleteFilter) GetNewModelWithSoftDeleteModel() NewModelWithSoftDelete {

	filter.model.Validate()

	return filter.model
}

func (filter *NewModelWithSoftDeleteFilter) GetNewModelWithSoftDeleteModelList() (data []NewModelWithSoftDelete, err error) {

	for i := range filter.list {
		filter.list[i].Validate()

		if !filter.list[i].IsValid() {
			err = filter.list[i].GetValidationError()
			return nil, err
		}
	}

	return filter.list, nil
}

func (filter *NewModelWithSoftDeleteFilter) SetNewModelWithSoftDeleteModel(typeModel NewModelWithSoftDelete) {

	filter.model = typeModel
}

func (filter *NewModelWithSoftDeleteFilter) SetNewModelWithSoftDeleteModelList(data []NewModelWithSoftDelete) {

	filter.list = data
}
