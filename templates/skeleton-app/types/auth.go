package types

import (
    "net/http"
    "skeleton-app/settings"
)

type Auth struct {
    
    Email     string
    Password  string
    Token     string
    UserId   int
    Id   int
    
    validator
}

func (auth *Auth) Validate()  {
}

type AuthFilter struct {
    model Auth
    list []Auth

    AbstractFilter
}

func GetAuthFilter(request *http.Request, functionType string) (filter AuthFilter, err error) {

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


func (filter *AuthFilter) GetAuthModel() Auth {

    filter.model.Validate()

    return  filter.model
}

func (filter *AuthFilter) GetAuthModelList() (data []Auth, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = filter.list[k].GetValidationError()
            break
        }
    }

    return  filter.list, nil
}

func (filter *AuthFilter) SetAuthModel(typeModel Auth) {

    filter.model = typeModel
}

func (filter *AuthFilter) SetAuthModelList(data []Auth) {

	filter.list = data
}
