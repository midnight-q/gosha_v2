package main

import (
	"fmt"
	"gosha_v2/router"
	"gosha_v2/settings"
	"net/http"
)

func main() {
	runHttpServer()
}

func runHttpServer() {
	fmt.Println("API сервер запущен :" + settings.ServerPort)
	panic(http.ListenAndServe("0.0.0.0:"+settings.ServerPort, router.Router()))
}
