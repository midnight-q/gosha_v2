package generator

import (
	"skeleton-app/types"
	"math/rand"
)

func GenRoleResource() types.RoleResource {

	return types.RoleResource{
		Id:   rand.Intn(100500),
		RoleId: rand.Intn(100500),
		ResourceId: rand.Intn(100500),
		Find: (rand.Intn(100500) % 2 > 0),
		Read: (rand.Intn(100500) % 2 > 0),
		Create: (rand.Intn(100500) % 2 > 0),
		Update: (rand.Intn(100500) % 2 > 0),
		Delete: (rand.Intn(100500) % 2 > 0),
		FindOrCreate: (rand.Intn(100500) % 2 > 0),
		UpdateOrCreate: (rand.Intn(100500) % 2 > 0),
	}
}

func GenListRoleResource() (list []types.RoleResource) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenRoleResource())
	}

	return
}
