package types

import (
	"gosha_v2/errors"
	"gosha_v2/settings"
	"net/http"
	"strings"
)

type Access struct {
	Find           bool
	Read           bool
	Create         bool
	Update         bool
	Delete         bool
	FindOrCreate   bool
	UpdateOrCreate bool
}

type Authenticator struct {
	Token        string
	functionType string
	urlPath      string
	ip           string
	maxPerPage   int
	userId       int
	roleIds      []int
	validator
}

func (auth *Authenticator) GetCurrentUserId() int {
	return auth.userId
}

func (auth *Authenticator) SetIp(r *http.Request) {
	auth.ip = r.Header.Get("X-Forwarded-For")
}

func (auth *Authenticator) GetIp() string {
	return auth.ip
}

func (auth *Authenticator) SetMaxPerPage(i int) {
	auth.maxPerPage = i
}
func (auth *Authenticator) GetMaxPerPage() int {
	return auth.maxPerPage
}

func (auth *Authenticator) SetCurrentUserId(id int) {
	auth.userId = id
}

func (auth *Authenticator) GetCurrentUserRoleIds() []int {
	return auth.roleIds
}

func (auth *Authenticator) IsAuthorized() bool {
	return true
}

func clearPath(s string) string {
	if strings.Count(s, "/") > 3 {
		return s[0:strings.LastIndex(s, "/")]
	}

	return s
}

func (auth *Authenticator) SetToken(r *http.Request) error {

	auth.Token = r.Header.Get("Token")

	return nil
}

func (authenticator *Authenticator) Validate(functionType string) {

	switch functionType {

	case settings.FunctionTypeFind:
		break
	case settings.FunctionTypeCreate:
		break
	case settings.FunctionTypeRead:
		break
	case settings.FunctionTypeUpdate:
		break
	case settings.FunctionTypeDelete:
		break
	case settings.FunctionTypeMultiCreate:
		break
	case settings.FunctionTypeMultiUpdate:
		break
	case settings.FunctionTypeMultiDelete:
		break
	default:
		authenticator.validator.AddValidationError("Unsupported function type: "+functionType, errors.ErrorCodeUnsupportedFunctionType, "")
		break
	}
}
