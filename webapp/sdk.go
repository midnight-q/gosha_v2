package webapp

import (
	"gosha_v2/logic"
	"gosha_v2/mdl"
	"gosha_v2/settings"
	"gosha_v2/types"
	"net/http"
)

func SdkCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto, err := types.GetSdkFilter(httpRequest, settings.FunctionTypeCreate)
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

	data, err := logic.SdkCreate(requestDto)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		Model: ApplyFieldsFilterToData(requestDto.GetFields(), data),
	})

	return
}
