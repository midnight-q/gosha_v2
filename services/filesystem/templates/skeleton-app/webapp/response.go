package webapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"skeleton-app/common"
	"skeleton-app/errors"
	"skeleton-app/logic"
	"skeleton-app/mdl"
	"skeleton-app/types"
	"strings"

	dynamicStruct "github.com/ompluscator/dynamic-struct"
)

type FilterInterface interface {
	IsDebug() bool
	GetLanguageId() int
}

func ApplyFieldsFilterToData(fields []string, data interface{}) (result interface{}) {
	if len(fields) < 1 {
		return data
	}
	b, _ := json.Marshal(&data)

	resMap := fieldPreprocessing(fields)

	result = buildStruct(resMap, data)
	err := json.Unmarshal(b, &result)

	if err != nil {
		fmt.Printf("ApplyFieldsFilterToData error: %+v", err.Error())
		return data
	}

	return
}

// Формирование карты полей
func fieldPreprocessing(fields []string) map[string]interface{} {
	res := map[string]interface{}{}

	nextFields := map[string][]string{}

	for _, field := range fields {
		if len(field) < 1 {
			continue
		}
		validFieldName := strings.Title(field)
		if !strings.Contains(validFieldName, ".") {
			res[validFieldName] = true
		} else {
			arr := strings.Split(validFieldName, ".")
			nextFields[arr[0]] = append(nextFields[arr[0]], strings.Join(arr[1:], "."))
		}
	}
	for k, v := range nextFields {
		res[k] = fieldPreprocessing(v)
	}
	return res
}

// Построение стуктуры на основе карты полей
func buildStruct(structMap map[string]interface{}, data interface{}) interface{} {
	var i *interface{}
	isArr := isArray(data)
	builder := dynamicStruct.NewStruct()
	for k, v := range structMap {
		if nextMap, isOk := v.(map[string]interface{}); isOk {
			builder.AddField(k, buildStruct(nextMap, getDataFromStructByFieldName(data, k)), "")
		} else {
			builder.AddField(k, i, "")
		}
	}
	if isArr {
		return builder.Build().NewSliceOfStructs()
	}
	return builder.Build().New()
}

//Получаем вложенные данные из структуры
func getDataFromStructByFieldName(data interface{}, field string) interface{} {
	var i *interface{}

	r := reflect.Indirect(reflect.ValueOf(data))
	if r.Kind() == reflect.Slice {
		el := reflect.Indirect(reflect.New(r.Type().Elem()))
		f := el.FieldByName(field)
		if !f.IsValid() {
			return i
		}
		return reflect.New(f.Type()).Interface()
	} else {
		if !r.IsValid() || r.IsZero() || !(r.Kind() == reflect.Struct) && r.IsNil() {
			return i
		}
		f := reflect.Indirect(r).FieldByName(field)
		if !(r.Kind() == reflect.Struct) && f.IsNil() || f.IsZero() || !f.IsValid() {
			return i
		}
		return f.Interface()
	}
}

func isArray(data interface{}) (res bool) {
	rt := reflect.Indirect(reflect.ValueOf(data))
	switch rt.Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	default:
		return false
	}
}

func Bad(w http.ResponseWriter, requestDto FilterInterface, err error) {
	ErrResponse(w, err, http.StatusBadRequest, requestDto)
}

func AuthErr(w http.ResponseWriter, requestDto FilterInterface) {

	ErrResponse(w, GetAuthErrTpl(common.MyCaller()), http.StatusForbidden, requestDto)
}

func GetAuthErrTpl(operation string) errors.ErrorWithCode {

	return errors.NewErrorWithCode(
		fmt.Sprintf("Invalid authorize in %s", operation),
		errors.ErrorCodeInvalidAuthorize,
		"Token")
}

func ErrResponse(w http.ResponseWriter, err error, status int, filter FilterInterface) {

	response := types.APIError{}
	response.Error = true

	switch e := err.(type) {
	case errors.ValidatorError:
		for _, errWithCode := range e.Errors() {
			newError := types.Error{
				Field:     errWithCode.GetField(),
				ErrorCode: errWithCode.ErrorCode(),
			}
			if filter.IsDebug() {
				newError.ErrorDebug = errWithCode.Error()
			}
			response.Errors = append(response.Errors, newError)
		}
		break

	case errors.ValidatorErrorInterface:
		for _, errWithCode := range e.Errors() {
			newError := types.Error{
				Field:     errWithCode.GetField(),
				ErrorCode: errWithCode.ErrorCode(),
			}
			if filter.IsDebug() {
				newError.ErrorDebug = errWithCode.Error()
			}
			response.Errors = append(response.Errors, newError)
		}
		break

	case errors.ErrorWithCode:
		newError := types.Error{
			Field:     e.GetField(),
			ErrorCode: e.ErrorCode(),
		}
		if filter.IsDebug() {
			newError.ErrorDebug = e.Error()
		}
		response.Errors = append(response.Errors, newError)
		break

	default:
		newError := types.Error{}
		if filter.IsDebug() {
			newError.ErrorDebug = e.Error()
		}
		response.Errors = append(response.Errors, newError)
		break
	}

	var errCodes []int
	for _, e := range response.Errors {
		errCodes = append(errCodes, e.ErrorCode)
	}
	errCodes = common.UniqueIntArray(errCodes)
	f := types.TranslateErrorFilter{}
	f.LanguageId = filter.GetLanguageId()
	if f.LanguageId < 1 {
		f.LanguageId = errors.DefaultErrorLanguageId
	}
	f.ErrorCodes = errCodes
	f.CurrentPage = 1
	f.PerPage = len(errCodes)
	translates, _, err := logic.TranslateErrorFind(f)
	if err != nil {
		fmt.Println("TranslateErrorFind err = ", err)
	}

	for i, err := range response.Errors {
		for _, translate := range translates {
			if err.ErrorCode == translate.Code {
				response.Errors[i].ErrorMessage = translate.Translate
				break
			}
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)

	return
}

func ValidResponse(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	switch data.(type) {
	case mdl.ResponseCreate:
		w.WriteHeader(http.StatusCreated)
		break
	default:
		w.WriteHeader(http.StatusOK)
		break
	}
	json.NewEncoder(w).Encode(data)

	return
}
