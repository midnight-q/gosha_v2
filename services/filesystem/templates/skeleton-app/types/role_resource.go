package types

import (
	"net/http"
	"skeleton-app/settings"
)

type RoleResource struct {
	Id             int
	RoleId         int
	ResourceId     int
	Find           bool
	Read           bool
	Create         bool
	Update         bool
	Delete         bool
	FindOrCreate   bool
	UpdateOrCreate bool

	validator
}

func (roleResource *RoleResource) Validate() {
}

type RoleResourceFilter struct {
	model RoleResource
	list  []RoleResource

	AbstractFilter
}

func GetRoleResourceFilter(request *http.Request, functionType string) (filter RoleResourceFilter, err error) {

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

func (filter *RoleResourceFilter) GetRoleResourceModel() RoleResource {

	filter.model.Validate()

	return filter.model
}

func (filter *RoleResourceFilter) GetRoleResourceModelList() (data []RoleResource, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *RoleResourceFilter) SetRoleResourceModel(typeModel RoleResource) {

	filter.model = typeModel
}

func (filter *RoleResourceFilter) SetRoleResourceModelList(data []RoleResource) {

	filter.list = data
}
