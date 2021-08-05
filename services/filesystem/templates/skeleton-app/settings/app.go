package settings

import (
	"fmt"
	"os"
	"regexp"
	"skeleton-app/flags"
)

const RabbitServerPassword = "OeWOpyZv"
const RabbitServerLogin = "skeleton-app"
const RabbitServer = "core.140140.ru"
const RabbitServerPort = "5672"
const rabbitServerVirtualhost = "/microservices"
const rabbitServerVirtualhostTest = "/microservices-test"

const MicroserviceAuthKey = "757fce54-f4f8-4487-af02-4dac1631c1ff"

func IsDev() bool {
	var matchDev = regexp.MustCompile("^/tmp/go-build")
	return matchDev.Match([]byte(os.Args[0])) || *flags.IsDev
}

func GetVirtualHost() string {

	for _, param := range os.Args {

		if param == "dev" {
			fmt.Println("RPC use DEV environment")
			return rabbitServerVirtualhostTest
		}
	}

	if IsDev() {
		fmt.Println("RPC use DEV environment")
		return rabbitServerVirtualhostTest
	}

	fmt.Println("RPC use PRODUCTION environment")
	return rabbitServerVirtualhost
}
