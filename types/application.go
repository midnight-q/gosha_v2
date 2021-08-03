package types

import (
	"gosha_v2/settings"
	"net/http"
)

type Application struct {
	Id           int
	Name         string
	Email        string
	Password     string
	UseUuidPk    bool
	DatabaseType int

	validator
}

func (application *Application) Validate() {
}

type ApplicationFilter struct {
	model Application
	list  []Application

	AbstractFilter
}

func GetApplicationFilter(request *http.Request, functionType string) (filter ApplicationFilter, err error) {

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

func (filter *ApplicationFilter) GetApplicationModel() Application {

	filter.model.Validate()

	return filter.model
}

func (filter *ApplicationFilter) GetApplicationModelList() (data []Application, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *ApplicationFilter) SetApplicationModel(typeModel Application) {

	filter.model = typeModel
}

func (filter *ApplicationFilter) SetApplicationModelList(data []Application) {

	filter.list = data
}
