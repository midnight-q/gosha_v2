package generator

import (
	"math/rand"
	"skeleton-app/types"
	"strings"
)

func GenLanguage() types.Language {

	return types.Language{
		Id:   rand.Intn(100500),
		Name: rand.Intn(100500),
		Code: strings.Title(Babbler2.Babble()),
	}
}

func GenListLanguage() (list []types.Language) {

	for i := 0; i < rand.Intn(5)+2; i++ {
		list = append(list, GenLanguage())
	}

	return
}
