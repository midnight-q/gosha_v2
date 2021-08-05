package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"skeleton-app/dbmodels"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

type WebTest struct {
	Name            string
	Request         *http.Request
	ResponseCode    int
	TestFunc        func(tt WebTest) (*httptest.ResponseRecorder, error)
	ResultValidator func(t *testing.T, response *httptest.ResponseRecorder)
}

func TestFunction(t *testing.T, tt WebTest) {

	response, err := tt.TestFunc(tt)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if response.Code != tt.ResponseCode {
		t.Error("Fail test", tt.Name, "expect", tt.ResponseCode, "got", response.Code, "response:", response.Body)
	}

	if tt.ResultValidator != nil {
		tt.ResultValidator(t, response)
	}
}

func SendRequest(routeUrl string, request *http.Request, webFunction func(w http.ResponseWriter, httpRequest *http.Request), method string) *httptest.ResponseRecorder {

	response := httptest.NewRecorder()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(routeUrl, webFunction).Methods(method)
	router.ServeHTTP(response, request)

	return response
}

func getRequest(method, route string, model interface{}, auth dbmodels.Auth) *http.Request {

	var jsonData []byte

	switch model.(type) {
	case int:
		break
	default:
		jsonData, _ = json.Marshal(model)
		break
	}

	r, _ := http.NewRequest(method, route, strings.NewReader(string(jsonData))) // URL-encoded payload
	r.Header.Add("Token", auth.Token)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(string(jsonData))))

	return r
}

func GetCreateAdminRequest(route string, model interface{}) *http.Request {

	auth := GetAdminAuth(true)
	return getRequest(http.MethodPost, route, model, auth)
}

func GetUpdateAdminRequest(route string, model interface{}) *http.Request {
	auth := GetAdminAuth(true)
	return getRequest(http.MethodPut, route, model, auth)
}

func GetCreateUserRequest(route string, model interface{}) *http.Request {

	auth := GetNonAdminUser(true)
	return getRequest(http.MethodPost, route, model, auth)
}

func GetCreateNonAuthorizedUserRequest(route string, model interface{}) *http.Request {
	return getRequest(http.MethodPost, route, model, dbmodels.Auth{})
}
func GetUpdateNonAuthorizedUserRequest(route string, model interface{}) *http.Request {
	return getRequest(http.MethodPut, route, model, dbmodels.Auth{})
}

func GetDeleteNonAuthorizedUserRequest(route string) *http.Request {
	return getRequest(http.MethodDelete, route, nil, dbmodels.Auth{})
}

func GetReadAdminRequest(route string) *http.Request {
	auth := GetAdminAuth(true)
	return getRequest(http.MethodGet, route, nil, auth)
}

func GetFindAdminRequest(route string, page, perPage int) *http.Request {

	auth := GetAdminAuth(true)
	req := getRequest(http.MethodGet, route, nil, auth)

	req.Form = url.Values{}

	req.Form.Add("CurrentPage", strconv.Itoa(page))
	req.Form.Add("PerPage", strconv.Itoa(perPage))
	req.Form.Add("Order[]", "id")
	req.Form.Add("OrderDirection[]", "desc")
	req.Form.Encode()

	return req
}

func GetReadNonAuthorizedUserRequest(route string) *http.Request {
	return getRequest(http.MethodGet, route, nil, dbmodels.Auth{})
}

func GetDeleteAdminRequest(route string) *http.Request {
	auth := GetAdminAuth(true)
	return getRequest(http.MethodDelete, route, nil, auth)
}
