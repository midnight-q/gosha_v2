package generator

import (
	"skeleton-app/types"
	"math/rand"
)

func GenUserRole() types.UserRole {

	return types.UserRole{
		Id:   rand.Intn(100500),
		UserId: rand.Intn(100500),
		RoleId: rand.Intn(100500),
		//UserRole remove this line for disable generator functionality
	}
}

func GenListUserRole() (list []types.UserRole) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenUserRole())
	}

	return
}
