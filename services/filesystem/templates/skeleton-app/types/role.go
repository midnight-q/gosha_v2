package types

import (
	"net/http"
	"skeleton-app/settings"
)

type Role struct {
	Id          int
	Name        string
	Description string

	validator
}

func (role *Role) Validate() {
}

type RoleFilter struct {
	model Role
	list  []Role

	AbstractFilter
}

func GetRoleFilter(request *http.Request, functionType string) (filter RoleFilter, err error) {

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

func (filter *RoleFilter) GetRoleModel() Role {

	filter.model.Validate()

	return filter.model
}

func (filter *RoleFilter) GetRoleModelList() (data []Role, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *RoleFilter) SetRoleModel(typeModel Role) {

	filter.model = typeModel
}

func (filter *RoleFilter) SetRoleModelList(data []Role) {

	filter.list = data
}
