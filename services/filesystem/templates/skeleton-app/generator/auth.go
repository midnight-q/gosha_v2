package generator

import (
	"math/rand"
	"skeleton-app/types"
)

func GenAuth() types.Auth {

	return types.Auth{
		Id:   rand.Intn(100500),
	}
}

func GenListAuth() (list []types.Auth) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenAuth())
	}

	return
}
