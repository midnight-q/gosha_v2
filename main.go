package main

import (
	"fmt"
	"gosha_v2/router"
	"gosha_v2/settings"
	"net/http"
)

func main() {
	//f := types.ModelFilter{}
	//fmt.Println(logic.ModelFind(f))
	//os.Exit(42)
	runHttpServer()
}

func runHttpServer() {
	fmt.Println("API сервер запущен :" + settings.ServerPort)
	panic(http.ListenAndServe("0.0.0.0:"+settings.ServerPort, router.Router()))
}
