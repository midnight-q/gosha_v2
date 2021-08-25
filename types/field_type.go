package types

import (
	"gosha_v2/settings"
	"net/http"
)

type FieldType struct {
	Name                 string
	Code                 string
	IsAvailableForFilter bool

	validator
}

func (field *FieldType) Validate() {
}

type FieldTypeFilter struct {
	model FieldType
	list  []FieldType

	AbstractFilter
}

func GetFieldTypeFilter(request *http.Request, functionType string) (filter FieldTypeFilter, err error) {

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

func (filter *FieldTypeFilter) GetFieldTypeModel() FieldType {

	filter.model.Validate()

	return filter.model
}

func (filter *FieldTypeFilter) GetFieldTypeModelList() (data []FieldType, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *FieldTypeFilter) SetFieldTypeModel(typeModel FieldType) {

	filter.model = typeModel
}

func (filter *FieldTypeFilter) SetFieldTypeModelList(data []FieldType) {

	filter.list = data
}
