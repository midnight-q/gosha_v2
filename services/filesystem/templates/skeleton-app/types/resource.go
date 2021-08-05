package types

import (
	"net/http"
	"skeleton-app/settings"
)

type Resource struct {
	Id     int
	Name   string
	Code   string
	TypeId int

	validator
}

func (resource *Resource) Validate() {
}

type ResourceFilter struct {
	model Resource
	list  []Resource

	AbstractFilter
}

func GetResourceFilter(request *http.Request, functionType string) (filter ResourceFilter, err error) {

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

func (filter *ResourceFilter) GetResourceModel() Resource {

	filter.model.Validate()

	return filter.model
}

func (filter *ResourceFilter) GetResourceModelList() (data []Resource, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *ResourceFilter) SetResourceModel(typeModel Resource) {

	filter.model = typeModel
}

func (filter *ResourceFilter) SetResourceModelList(data []Resource) {

	filter.list = data
}
