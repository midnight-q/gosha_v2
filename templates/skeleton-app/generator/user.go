package generator

import (
	"skeleton-app/types"
	"math/rand"
)

func GenUser() types.User {

	return types.User{
		Id:   rand.Intn(100500),
	}
}

func GenListUser() (list []types.User) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenUser())
	}

	return
}
