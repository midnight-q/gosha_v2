package generator

import (
	"skeleton-app/types"
	"math/rand"
)

func GenCurrentUser() types.CurrentUser {

	return types.CurrentUser{
		Id:   rand.Intn(100500),
	}
}

func GenListCurrentUser() (list []types.CurrentUser) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenCurrentUser())
	}

	return
}
