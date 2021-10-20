package types

import (
	"gosha_v2/settings"
	"net/http"
	"strconv"
)

type HttpMethods struct {
	Find           bool
	Create         bool
	MultiCreate    bool
	Read           bool
	Update         bool
	MultiUpdate    bool
	Delete         bool
	MultiDelete    bool
	FindOrCreate   bool
	UpdateOrCreate bool
}

type Model struct {
	Name        string
	Fields      []Field
	IsTypeModel bool
	IsDbModel   bool
	PkType      string
	CommentType string
	CommentDb   string
	Filter      *Model
	TypePath    string
	DbPath      string
	HttpRoutes  struct {
		Find           string
		Create         string
		MultiCreate    string
		Read           string
		Update         string
		MultiUpdate    string
		Delete         string
		MultiDelete    string
		FindOrCreate   string
		UpdateOrCreate string
	}

	HttpMethods struct {
		Find           bool
		Create         bool
		MultiCreate    bool
		Read           bool
		Update         bool
		MultiUpdate    bool
		Delete         bool
		MultiDelete    bool
		FindOrCreate   bool
		UpdateOrCreate bool
	}
	IsServiceModel     bool
	IsSoftDeleteEnable bool
	validator
}

func (model *Model) Validate() {
}

type ModelFilter struct {
	model Model
	list  []Model

	IsShowServiceModels bool

	AbstractFilter
}

func GetModelFilter(request *http.Request, functionType string) (filter ModelFilter, err error) {

	filter.request = request
	filter.rawRequestBody, err = GetRawBodyContent(request)
	if err != nil {
		return filter, err
	}

	filter.IsShowServiceModels, _ = strconv.ParseBool(request.FormValue("IsShowServiceModels"))

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

func (filter *ModelFilter) GetModelModel() Model {

	filter.model.Validate()

	return filter.model
}

func (filter *ModelFilter) GetModelModelList() (data []Model, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *ModelFilter) SetModelModel(typeModel Model) {

	filter.model = typeModel
}

func (filter *ModelFilter) SetModelModelList(data []Model) {

	filter.list = data
}
