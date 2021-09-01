package webapp

import (
	"net/http"
	"skeleton-app/core"
	"skeleton-app/logic"
	"skeleton-app/mdl"
	"skeleton-app/settings"
	"skeleton-app/types"
)

func NewModelFind(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelFilter(httpRequest, settings.FunctionTypeFind)
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
	data, totalRecords, err := logic.NewModelFind(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseFind{
		List:  ApplyFieldsFilterToData(requestDto.GetFields(), data),
		Total: totalRecords,
	})

	return
}

func NewModelMultiCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelFilter(httpRequest, settings.FunctionTypeMultiCreate)
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

	data, err := logic.NewModelMultiCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model: data,
	})

	return
}

func NewModelCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelFilter(httpRequest, settings.FunctionTypeCreate)
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

	data, err := logic.NewModelCreate(requestDto, core.Db)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func NewModelRead(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelFilter(httpRequest, settings.FunctionTypeRead)
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

	data, err := logic.NewModelRead(requestDto)

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
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func NewModelMultiUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelFilter(httpRequest, settings.FunctionTypeMultiUpdate)
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

	data, err := logic.NewModelMultiUpdate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		Model: data,
	})

	return
}

func NewModelUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelFilter(httpRequest, settings.FunctionTypeUpdate)
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

	data, err := logic.NewModelUpdate(requestDto, core.Db)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func NewModelMultiDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelFilter(httpRequest, settings.FunctionTypeMultiDelete)
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

	isOk, err := logic.NewModelMultiDelete(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		IsSuccess: isOk,
	})

	return
}

func NewModelDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelFilter(httpRequest, settings.FunctionTypeDelete)
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

	isOk, err := logic.NewModelDelete(requestDto, core.Db)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		IsSuccess: isOk,
	})

	return
}

func NewModelFindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelFilter(httpRequest, settings.FunctionTypeFindOrCreate)
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

	data, err := logic.NewModelFindOrCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseFindOrCreate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func NewModelUpdateOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelFilter(httpRequest, settings.FunctionTypeUpdateOrCreate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.NewModelUpdateOrCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdateOrCreate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}
