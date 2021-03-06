package types

import (
	"net/http"
	"skeleton-app/settings"
)

type User struct {
	Id          int
	Email       string
	FirstName   string
	IsActive    bool
	LastName    string
	MobilePhone string
	Password    string

	validator
}

func (user *User) Validate() {
}

type UserFilter struct {
	model User
	list  []User

	AbstractFilter
}

func GetUserFilter(request *http.Request, functionType string) (filter UserFilter, err error) {

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

func (filter *UserFilter) GetUserModel() User {

	filter.model.Validate()

	return filter.model
}

func (filter *UserFilter) GetUserModelList() (data []User, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *UserFilter) SetUserModel(typeModel User) {

	filter.model = typeModel
}

func (filter *UserFilter) SetUserModelList(data []User) {

	filter.list = data
}
