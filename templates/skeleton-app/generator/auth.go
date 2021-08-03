package generator

import (
	"skeleton-app/types"
	"math/rand"
)

func GenAuth() types.Auth {

	return types.Auth{
		Id:   rand.Intn(100500),
		//Auth remove this line for disable generator functionality
	}
}

func GenListAuth() (list []types.Auth) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenAuth())
	}

	return
}
