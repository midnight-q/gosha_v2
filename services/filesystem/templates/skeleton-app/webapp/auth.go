package webapp

import (
	"net/http"
	"skeleton-app/core"
	"skeleton-app/logic"
	"skeleton-app/mdl"
	"skeleton-app/settings"
	"skeleton-app/types"
)

func AuthFind(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetAuthFilter(httpRequest, settings.FunctionTypeFind)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	// Получаем список
	data, totalRecords, err := logic.AuthFind(requestDto)

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

func AuthMultiCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetAuthFilter(httpRequest, settings.FunctionTypeMultiCreate)
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

	data, err := logic.AuthMultiCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		data,
	})

	return
}

func AuthCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetAuthFilter(httpRequest, settings.FunctionTypeCreate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.AuthCreate(requestDto, core.Db)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func AuthRead(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetAuthFilter(httpRequest, settings.FunctionTypeRead)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	requestDto.PerPage = 1
	requestDto.CurrentPage = 1

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.AuthRead(requestDto)

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

func AuthMultiUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetAuthFilter(httpRequest, settings.FunctionTypeMultiUpdate)
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

	data, err := logic.AuthMultiUpdate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		data,
	})

	return
}

func AuthUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetAuthFilter(httpRequest, settings.FunctionTypeUpdate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.AuthUpdate(requestDto, core.Db)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func AuthMultiDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetAuthFilter(httpRequest, settings.FunctionTypeMultiDelete)
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

	isOk, err := logic.AuthMultiDelete(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		isOk,
	})

	return
}

func AuthDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetAuthFilter(httpRequest, settings.FunctionTypeDelete)
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

	isOk, err := logic.AuthDelete(requestDto, core.Db)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		isOk,
	})

	return
}

func AuthFindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetAuthFilter(httpRequest, settings.FunctionTypeFindOrCreate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.AuthFindOrCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseFindOrCreate{
		ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func AuthUpdateOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetAuthFilter(httpRequest, settings.FunctionTypeUpdateOrCreate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.AuthUpdateOrCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdateOrCreate{
		ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}
