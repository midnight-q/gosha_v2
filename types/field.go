package types

import (
	"gosha_v2/settings"
	"net/http"
)

type Field struct {
	// Field name
	Name string
	// Field type
	Type string
	// Field comment from typeModel
	CommentType string
	// Field comment from dbModel
	CommentDb string
	// Flag - field exist in typeModel
	IsTypeField bool
	// Flag - field exist in dbModel
	IsDbField bool
	// Path for parent model in composition
	SourceModel string

	// Model name for create field
	ModelName string
	// Flag for creating func
	IsPointer bool
	// Flag for creating func
	IsArray bool

	validator
}

func (field *Field) Validate() {
}

type FieldFilter struct {
	model Field
	list  []Field

	AbstractFilter
}

func GetFieldFilter(request *http.Request, functionType string) (filter FieldFilter, err error) {

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
