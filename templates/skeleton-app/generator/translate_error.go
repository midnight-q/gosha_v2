package generator

import (
    "skeleton-app/types"
    "math/rand"
    "strings"
)

func GenTranslateError() types.TranslateError {

	return types.TranslateError{
		Id:   rand.Intn(100500),
		Code: rand.Intn(100500),
		LanguageCode: strings.Title(Babbler2.Babble()),
		Translate: strings.Title(Babbler2.Babble()),
	}
}

func GenListTranslateError() (list []types.TranslateError) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenTranslateError())
	}

	return
}
