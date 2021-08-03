package types

import (
    "net/http"
    "skeleton-app/settings"
)

type Language struct {
    Id   int
    Name int
	Code string

	validator
}

func (language *Language) Validate()  {
}

type LanguageFilter struct {
    model Language
    list  []Language

	AbstractFilter
}

func GetLanguageFilter(request *http.Request, functionType string) (filter LanguageFilter, err error) {

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

    return  filter, err
}


func (filter *LanguageFilter) GetLanguageModel() Language {

    filter.model.Validate()

    return  filter.model
}

func (filter *LanguageFilter) GetLanguageModelList() (data []Language, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = filter.list[k].GetValidationError()
            break
        }
    }

    return  filter.list, nil
}

func (filter *LanguageFilter) SetLanguageModel(typeModel Language) {

    filter.model = typeModel
}

func (filter *LanguageFilter) SetLanguageModelList(data []Language) {

	filter.list = data
}
