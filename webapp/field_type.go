package webapp

import (
	"gosha_v2/logic"
	"gosha_v2/mdl"
	"gosha_v2/settings"
	"gosha_v2/types"
	"net/http"
)

func FieldTypeFind(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetFieldTypeFilter(httpRequest, settings.FunctionTypeFind)
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
	data, totalRecords, err := logic.FieldTypeFind(requestDto)

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
