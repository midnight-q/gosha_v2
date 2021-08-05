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

func GetUuidType() *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X: &dst.Ident{
			Name: "uuid",
		},
		Sel: &dst.Ident{
			Name: "UUID",
		},
	}
}

func GetArrayType(t dst.Expr) *dst.ArrayType {
	return &dst.ArrayType{
		Elt: t,
	}
}
