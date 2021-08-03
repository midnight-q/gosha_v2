package types

import (
	"gosha_v2/settings"
	"net/http"
)

type Field struct {
	Id   int
	Name string
	//Field remove this line for disable generator functionality

	validator
}

func (field *Field) Validate() {
	//Validate remove this line for disable generator functionality
}

type FieldFilter struct {
	model Field
	list  []Field
	//FieldFilter remove this line for disable generator functionality

	AbstractFilter
}

func GetFieldFilter(request *http.Request, functionType string) (filter FieldFilter, err error) {

	filter.request = request
	filter.rawRequestBody, err = GetRawBodyContent(request)
	if err != nil {
		return filter, err
	}
	//filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

	//GetFieldFilter remove this line for disable generator functionality

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

func (filter *FieldFilter) GetFieldModel() Field {

	filter.model.Validate()

	return filter.model
}

func (filter *FieldFilter) GetFieldModelList() (data []Field, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *FieldFilter) SetFieldModel(typeModel Field) {

	filter.model = typeModel
}

func (filter *FieldFilter) SetFieldModelList(data []Field) {

	filter.list = data
}