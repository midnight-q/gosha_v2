package webapp

import (
	"gosha_v2/logic"
	"gosha_v2/mdl"
	"gosha_v2/settings"
	"gosha_v2/types"
	"net/http"
)

func ModelFind(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetModelFilter(httpRequest, settings.FunctionTypeFind)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	// Получаем список
	data, totalRecords, err := logic.ModelFind(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseFind{
		ApplyFieldsFilterToData(requestDto.GetFields(), data),
		totalRecords,
	})

	return
}

func ModelMultiCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetModelFilter(httpRequest, settings.FunctionTypeMultiCreate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.ModelMultiCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model:data,
	})

	return
}

func ModelCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetModelFilter(httpRequest, settings.FunctionTypeCreate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.ModelCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model:ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func ModelRead(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetModelFilter(httpRequest, settings.FunctionTypeRead)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}

	requestDto.PerPage = 1
	requestDto.CurrentPage = 1

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.ModelRead(requestDto)

	// Создаём структуру ответа
	if err != nil {
		code := http.StatusBadRequest
		if err.Error() == "Not found" {
			code = http.StatusNotFound
		}
		ErrResponse(w, err, code, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseRead{
		ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func ModelMultiUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetModelFilter(httpRequest, settings.FunctionTypeMultiUpdate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.ModelMultiUpdate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		Model:data,
	})

	return
}

func ModelUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetModelFilter(httpRequest, settings.FunctionTypeUpdate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.ModelUpdate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		Model:ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func ModelMultiDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetModelFilter(httpRequest, settings.FunctionTypeMultiDelete)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	isOk, err := logic.ModelMultiDelete(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		IsSuccess:isOk,
	})

	return
}

func ModelDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetModelFilter(httpRequest, settings.FunctionTypeDelete)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	isOk, err := logic.ModelDelete(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		IsSuccess:isOk,
	})

	return
}

func ModelFindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetModelFilter(httpRequest, settings.FunctionTypeFindOrCreate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.ModelFindOrCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseFindOrCreate{
		ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func ModelUpdateOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetModelFilter(httpRequest, settings.FunctionTypeUpdateOrCreate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.ModelUpdateOrCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdateOrCreate{
		ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}
