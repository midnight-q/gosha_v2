package generator

import (
	"math/rand"
	"skeleton-app/types"
	"strings"
)

func GenRegion() types.Region {

	return types.Region{
		Id:   rand.Intn(100500),
		Name: rand.Intn(100500),
		Code: strings.Title(Babbler2.Babble()),
	}
}

func GenListRegion() (list []types.Region) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenRegion())
	}

	return
}
