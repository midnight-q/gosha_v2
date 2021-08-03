package utils

import (
	"github.com/dave/dst"
)

func GetIntType() *dst.Ident {
	return &dst.Ident{Name: "int"}
}

func GetStringType() *dst.Ident {
	return &dst.Ident{Name: "string"}
}

func GetFloat64Type() *dst.Ident {
	return &dst.Ident{Name: "float64"}
}

func GetFloat64ArrayType() *dst.ArrayType {
	return &dst.ArrayType{
		Elt: GetFloat64Type(),
	}
}

func GetIntArrayType() *dst.ArrayType {
	return &dst.ArrayType{
		Elt: GetIntType(),
	}
}

func GetStringArrayType() *dst.ArrayType {
	return &dst.ArrayType{
		Elt: GetStringType(),
	}
}
