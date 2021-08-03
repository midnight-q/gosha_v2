package generator

import (
	"math/rand"
	"skeleton-app/types"
	"strings"
)

func GenRole() types.Role {

	return types.Role{
		Id:   rand.Intn(100500),
		Name: strings.Title(Babbler2.Babble()),
		Description: strings.Title(Babbler2.Babble()),
	}
}

func GenListRole() (list []types.Role) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenRole())
	}

	return
}