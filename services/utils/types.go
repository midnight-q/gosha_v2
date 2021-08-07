package utils

import (
	"fmt"

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

func ParseType(in dst.Expr) string {
	switch t := in.(type) {
	case *dst.Ident:
		return t.Name
	case *dst.SelectorExpr:
		x, isOk := t.X.(*dst.Ident)
		if !isOk {
			return ""
		}
		return x.Name + "." + t.Sel.Name
	case *dst.StarExpr:
		if ident, ok := t.X.(*dst.Ident); ok {
			return "*" + ident.Name
		} else {
			x, ok := t.X.(*dst.SelectorExpr)
			if !ok {
				return ""
			}
			if ident, ok := x.X.(*dst.Ident); ok {
				return "*" + ident.Name + "." + x.Sel.Name
			}
		}
	case *dst.BasicLit:
		return t.Value

	case *dst.ArrayType:
		arrLen := ""
		if t.Len != nil {
			l := t.Len.(*dst.BasicLit)
			arrLen = l.Value
		}
		return "[" + arrLen + "]" + ParseType(t.Elt)

	default:
		fmt.Printf("Unsopported type: %T\n", in)
		return "unknown"
	}
	return ""
}
