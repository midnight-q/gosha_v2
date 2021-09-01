package main

import (
	"fmt"
	"net/http"
	"skeleton-app/bootstrap"
	"skeleton-app/flags"
	"skeleton-app/router"
	"skeleton-app/settings"
	"skeleton-app/wsserver"
)

func main() {

	// делаем автомиграцию
	bootstrap.FillDBTestData()

	if *flags.IsDev {
		fmt.Println("Running in DEV mode")
	} else {
		fmt.Println("Running in PROD mode")
	}

	go runWsServer()
	runHttpServer()
}

func runWsServer() {

	fmt.Println("Websocket сервер запущен :" + settings.WssPort)
	wsserver.SetMessageHandler("", router.HandleWss)
	panic(wsserver.Run("", settings.WssPort))
}

func runHttpServer() {

	fmt.Println("API сервер запущен :" + settings.ServerPort)
	panic(http.ListenAndServe("0.0.0.0:"+settings.ServerPort, router.Router()))
}
