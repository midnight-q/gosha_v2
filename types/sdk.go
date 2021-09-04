package types

import (
	"gosha_v2/settings"
	"net/http"
)

type Sdk struct {
	// Js sdk with vuex store for gosha_v1 compatibility
	JsTypes bool
	Js      bool
	Ts      bool
	Go      bool
	Dart    bool

	Namespace string
	validator
}

func (sdk *Sdk) Validate() {

}

type SdkFilter struct {
	model Sdk
	list  []Sdk

	AbstractFilter
}

func GetSdkFilter(request *http.Request, functionType string) (filter SdkFilter, err error) {

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

func (filter *SdkFilter) GetSdkModel() Sdk {

	filter.model.Validate()

	return filter.model
}

func (filter *SdkFilter) GetSdkModelList() (data []Sdk, err error) {

	for k := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *SdkFilter) SetSdkModel(typeModel Sdk) {

	filter.model = typeModel
}

func (filter *SdkFilter) SetSdkModelList(data []Sdk) {

	filter.list = data
}
