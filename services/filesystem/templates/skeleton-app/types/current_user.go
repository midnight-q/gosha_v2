package types

import (
	"net/http"
	"skeleton-app/settings"
)

type CurrentUser struct {
	Id int

	validator
}

func (currentUser *CurrentUser) Validate() {
}

type CurrentUserFilter struct {
	model CurrentUser
	list  []CurrentUser

	AbstractFilter
}

func GetCurrentUserFilter(request *http.Request, functionType string) (filter CurrentUserFilter, err error) {

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

func (filter *CurrentUserFilter) GetCurrentUserModel() CurrentUser {

	filter.model.Validate()

	return filter.model
}

func (filter *CurrentUserFilter) GetCurrentUserModelList() (data []CurrentUser, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *CurrentUserFilter) SetCurrentUserModel(typeModel CurrentUser) {

	filter.model = typeModel
}

func (filter *CurrentUserFilter) SetCurrentUserModelList(data []CurrentUser) {

	filter.list = data
}
