package webapp

import (
	"gosha_v2/logic"
	"gosha_v2/mdl"
	"gosha_v2/settings"
	"gosha_v2/types"
	"net/http"
)

func FieldMultiCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetFieldFilter(httpRequest, settings.FunctionTypeMultiCreate)
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

	data, err := logic.FieldMultiCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model: data,
	})

	return
}

func FieldCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetFieldFilter(httpRequest, settings.FunctionTypeCreate)
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

	data, err := logic.FieldCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func FieldMultiUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetFieldFilter(httpRequest, settings.FunctionTypeMultiUpdate)
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

	data, err := logic.FieldMultiUpdate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		Model: data,
	})

	return
}

func FieldUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetFieldFilter(httpRequest, settings.FunctionTypeUpdate)
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

	data, err := logic.FieldUpdate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func FieldMultiDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetFieldFilter(httpRequest, settings.FunctionTypeMultiDelete)
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

	isOk, err := logic.FieldMultiDelete(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		IsSuccess: isOk,
	})

	return
}

func FieldDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetFieldFilter(httpRequest, settings.FunctionTypeDelete)
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

	isOk, err := logic.FieldDelete(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		IsSuccess: isOk,
	})

	return
}
