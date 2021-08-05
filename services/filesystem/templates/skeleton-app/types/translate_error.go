package types

import (
	"net/http"
	"net/url"
	"skeleton-app/settings"
	"strconv"
)

type TranslateError struct {
	Id           int
	Code         int
	LanguageCode string
	Translate    string

	validator
}

func (translateError *TranslateError) Validate() {
}

type TranslateErrorFilter struct {
	model      TranslateError
	list       []TranslateError
	ErrorCodes []int

	AbstractFilter
}

func GetTranslateErrorFilter(request *http.Request, functionType string) (filter TranslateErrorFilter, err error) {

	filter.request = request
	filter.rawRequestBody, err = GetRawBodyContent(request)
	if err != nil {
		return filter, err
	}

	arrErrorCodes, _ := url.ParseQuery(request.URL.RawQuery)
	for _, f := range arrErrorCodes["ErrorCodes[]"] {
		v, _ := strconv.Atoi(f)
		filter.ErrorCodes = append(filter.ErrorCodes, v)
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

func (filter *TranslateErrorFilter) GetTranslateErrorModel() TranslateError {

	filter.model.Validate()

	return filter.model
}

func (filter *TranslateErrorFilter) GetTranslateErrorModelList() (data []TranslateError, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *TranslateErrorFilter) SetTranslateErrorModel(typeModel TranslateError) {

	filter.model = typeModel
}

func (filter *TranslateErrorFilter) SetTranslateErrorModelList(data []TranslateError) {

	filter.list = data
}
