package generator

import (
    "skeleton-app/types"
    "math/rand"
    "strings"
)

func GenResourceType() types.ResourceType {

	return types.ResourceType{
		Id:   rand.Intn(100500),
		Name: strings.Title(Babbler2.Babble()),
		//ResourceType remove this line for disable generator functionality
	}
}

func GenListResourceType() (list []types.ResourceType) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenResourceType())
	}

	return
}
