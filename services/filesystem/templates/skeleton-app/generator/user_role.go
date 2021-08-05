package generator

import (
	"math/rand"
	"skeleton-app/types"
)

func GenUserRole() types.UserRole {

	return types.UserRole{
		Id:     rand.Intn(100500),
		UserId: rand.Intn(100500),
		RoleId: rand.Intn(100500),
	}
}

func GenListUserRole() (list []types.UserRole) {

	for i := 0; i < rand.Intn(5)+2; i++ {
		list = append(list, GenUserRole())
	}

	return
}
