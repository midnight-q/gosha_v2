package webapp

import (
	"net/http"
	"skeleton-app/core"
	"skeleton-app/logic"
	"skeleton-app/mdl"
	"skeleton-app/settings"
	"skeleton-app/types"
)

func NewModelWithSoftDeleteFind(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelWithSoftDeleteFilter(httpRequest, settings.FunctionTypeFind)
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
	data, totalRecords, err := logic.NewModelWithSoftDeleteFind(requestDto)

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

func NewModelWithSoftDeleteMultiCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelWithSoftDeleteFilter(httpRequest, settings.FunctionTypeMultiCreate)
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

	data, err := logic.NewModelWithSoftDeleteMultiCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model: data,
	})

	return
}

func NewModelWithSoftDeleteCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelWithSoftDeleteFilter(httpRequest, settings.FunctionTypeCreate)
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

	data, err := logic.NewModelWithSoftDeleteCreate(requestDto, core.Db)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func NewModelWithSoftDeleteRead(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelWithSoftDeleteFilter(httpRequest, settings.FunctionTypeRead)
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

	data, err := logic.NewModelWithSoftDeleteRead(requestDto)

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

func NewModelWithSoftDeleteMultiUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelWithSoftDeleteFilter(httpRequest, settings.FunctionTypeMultiUpdate)
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

	data, err := logic.NewModelWithSoftDeleteMultiUpdate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		Model: data,
	})

	return
}

func NewModelWithSoftDeleteUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelWithSoftDeleteFilter(httpRequest, settings.FunctionTypeUpdate)
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

	data, err := logic.NewModelWithSoftDeleteUpdate(requestDto, core.Db)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func NewModelWithSoftDeleteMultiDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelWithSoftDeleteFilter(httpRequest, settings.FunctionTypeMultiDelete)
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

	isOk, err := logic.NewModelWithSoftDeleteMultiDelete(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		IsSuccess: isOk,
	})

	return
}

func NewModelWithSoftDeleteDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelWithSoftDeleteFilter(httpRequest, settings.FunctionTypeDelete)
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

	isOk, err := logic.NewModelWithSoftDeleteDelete(requestDto, core.Db)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		IsSuccess: isOk,
	})

	return
}

func NewModelWithSoftDeleteFindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelWithSoftDeleteFilter(httpRequest, settings.FunctionTypeFindOrCreate)
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

	data, err := logic.NewModelWithSoftDeleteFindOrCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseFindOrCreate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}

func NewModelWithSoftDeleteUpdateOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetNewModelWithSoftDeleteFilter(httpRequest, settings.FunctionTypeUpdateOrCreate)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}

	data, err := logic.NewModelWithSoftDeleteUpdateOrCreate(requestDto)

	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseUpdateOrCreate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}
