package generator

import (
    "skeleton-app/types"
    "math/rand"
    "strings"
)

func GenResource() types.Resource {

	return types.Resource{
		Id:   rand.Intn(100500),
		Name: strings.Title(Babbler2.Babble()),
		Code: strings.Title(Babbler2.Babble()),
		TypeId: rand.Intn(100500),
	}
}

func GenListResource() (list []types.Resource) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenResource())
	}

	return
}
