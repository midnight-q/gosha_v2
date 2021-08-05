package webapp

import (
	"gosha_v2/logic"
	"gosha_v2/mdl"
	"gosha_v2/settings"
	"gosha_v2/types"
	"net/http"
)

func ApplicationFind(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetApplicationFilter(httpRequest, settings.FunctionTypeFind)
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
	data, totalRecords, err := logic.ApplicationFind(requestDto)

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

func ApplicationMultiCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetApplicationFilter(httpRequest, settings.FunctionTypeMultiCreate)
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

	data, err := logic.ApplicationMultiCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model: data,
	})

	return
}

func ApplicationCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetApplicationFilter(httpRequest, settings.FunctionTypeCreate)
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

	data, err := logic.ApplicationCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func ApplicationRead(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetApplicationFilter(httpRequest, settings.FunctionTypeRead)
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

	data, err := logic.ApplicationRead(requestDto)

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

func ApplicationMultiUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetApplicationFilter(httpRequest, settings.FunctionTypeMultiUpdate)
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

	data, err := logic.ApplicationMultiUpdate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		Model: data,
	})

	return
}

func ApplicationUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetApplicationFilter(httpRequest, settings.FunctionTypeUpdate)
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

	data, err := logic.ApplicationUpdate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func ApplicationMultiDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetApplicationFilter(httpRequest, settings.FunctionTypeMultiDelete)
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

	isOk, err := logic.ApplicationMultiDelete(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		IsSuccess: isOk,
	})

	return
}

func ApplicationDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetApplicationFilter(httpRequest, settings.FunctionTypeDelete)
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

	isOk, err := logic.ApplicationDelete(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		IsSuccess: isOk,
	})

	return
}

func ApplicationFindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetApplicationFilter(httpRequest, settings.FunctionTypeFindOrCreate)
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

	data, err := logic.ApplicationFindOrCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseFindOrCreate{
		ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func ApplicationUpdateOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetApplicationFilter(httpRequest, settings.FunctionTypeUpdateOrCreate)
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

	data, err := logic.ApplicationUpdateOrCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdateOrCreate{
		ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}
