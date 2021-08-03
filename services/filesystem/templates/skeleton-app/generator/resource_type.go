package generator

import (
	"math/rand"
	"skeleton-app/types"
	"strings"
)

func GenResourceType() types.ResourceType {

	return types.ResourceType{
		Id:   rand.Intn(100500),
		Name: strings.Title(Babbler2.Babble()),
	}
}

func GenListResourceType() (list []types.ResourceType) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenResourceType())
	}

	return
}
