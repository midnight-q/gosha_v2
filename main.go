package main

import (
	"fmt"
	"gosha_v2/logic"
	"gosha_v2/router"
	"gosha_v2/settings"
	"gosha_v2/types"
	"net/http"
	"os"
)

func main() {
	f := types.ApplicationFilter{}
	f.SetApplicationModel(types.Application{
		Name:         "example-app",
		Email:        "test@mail.com",
		Password:     "qwerty12",
		UseUuidPk:    true,
		DatabaseType: 1,
	})
	_, err := logic.ApplicationCreate(f)
	fmt.Println(err)
	os.Exit(42)

	runHttpServer()
}

func runHttpServer() {
	fmt.Println("API сервер запущен :" + settings.ServerPort)
	panic(http.ListenAndServe("0.0.0.0:"+settings.ServerPort, router.Router()))
}
