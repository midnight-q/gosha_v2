package main

import (
	"fmt"
	"gosha_v2/router"
	"gosha_v2/services/filesystem"
	"gosha_v2/settings"
	"gosha_v2/types"
	"net/http"
	"os"
)

func main() {
	err := filesystem.AddParserInFilter(types.Field{
		Name:      "FilterField",
		Type:      "int",
		ModelName: "Auth",
		IsPointer: false,
		IsArray:   false,
	}, "/home/alex/Projects/gosha-test/types/auth.go")

	fmt.Println(err)
	os.Exit(42)
	runHttpServer()
}

func runHttpServer() {
	fmt.Println("API сервер запущен :" + settings.ServerPort)
	panic(http.ListenAndServe("0.0.0.0:"+settings.ServerPort, router.Router()))
}
