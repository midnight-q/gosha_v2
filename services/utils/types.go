package utils

import (
	"fmt"
	"go/token"
	"gosha_v2/errors"
	"gosha_v2/types"
	"strings"

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

func GetByteType() *dst.Ident {
	return &dst.Ident{Name: "byte"}
}

func GetBoolType() *dst.Ident {
	return &dst.Ident{Name: "bool"}
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

func GetTimeType() *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X: &dst.Ident{
			Name: "time",
		},
		Sel: &dst.Ident{
			Name: "Time",
		},
	}
}

func GetArrayType(t dst.Expr) *dst.ArrayType {
	return &dst.ArrayType{
		Elt: t,
	}
}

func GetPointerType(t dst.Expr) *dst.StarExpr {
	return &dst.StarExpr{
		X: t,
	}
}

func GetType(name string, isArray bool, isPointer bool) (dst.Expr, error) {
	var t dst.Expr
	switch name {
	case "int":
		t = GetIntType()
	case "float64":
		t = GetFloat64Type()
	case "string":
		t = GetStringType()
	case "byte":
		t = GetByteType()
	case "uuid":
		t = GetUuidType()
	case "time":
		t = GetTimeType()
	case "bool":
		t = GetBoolType()
	default:
		//TODO: implement model types
		return nil, errors.New("Unknown type: " + name)
	}
	if isArray {
		t = GetArrayType(t)
	}
	if isPointer {
		t = GetPointerType(t)
	}
	return t, nil
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

func IsTypeArray(typeStr string) bool {
	return strings.Contains(typeStr, "[]")
}
func IsTypePointer(typeStr string) bool {
	return strings.Contains(typeStr, "*")
}
func ClearType(typeStr string) string {
	typeStr = strings.Replace(typeStr, "[]", "", -1)
	typeStr = strings.Replace(typeStr, "*", "", -1)
	if typeStr == "time.Time" {
		return "time"
	}
	if typeStr == "uuid.Uuid" {
		return "uuid"
	}

	return typeStr
}

func GetParserForType(model types.Field) (*dst.AssignStmt, error) {
	switch model.Type {
	case "int":
		if model.IsArray {
			return GetParser(model.Name, "ParseIntArrayFromRequest"), nil
		}
		if model.IsPointer {
			return GetParser(model.Name, "ParseIntPointerFromRequest"), nil
		}
		return GetParser(model.Name, "ParseIntFromRequest"), nil

	case "float64":
		if model.IsArray {
			return GetParser(model.Name, "ParseFloatArrayFromRequest"), nil
		}
		if model.IsPointer {
			return GetParser(model.Name, "ParseFloatPointerFromRequest"), nil
		}
		return GetParser(model.Name, "ParseFloatFromRequest"), nil

	case "bool":
		if model.IsArray {
			return GetParser(model.Name, "ParseBoolArrayFromRequest"), nil
		}
		if model.IsPointer {
			return GetParser(model.Name, "ParseBoolPointerFromRequest"), nil
		}
		return GetParser(model.Name, "ParseBoolFromRequest"), nil

	case "string":
		if model.IsArray {
			return GetParser(model.Name, "ParseStringArrayFromRequest"), nil
		}
		if model.IsPointer {
			return GetParser(model.Name, "ParseStringPointerFromRequest"), nil
		}
		return GetParser(model.Name, "ParseStringFromRequest"), nil

	case "uuid":
		if model.IsArray {
			return GetParser(model.Name, "ParseUuidArrayFromRequest"), nil
		}
		if model.IsPointer {
			return GetParser(model.Name, "ParseUuidPointerFromRequest"), nil
		}
		return GetParser(model.Name, "ParseUuidFromRequest"), nil

	case "time":
		if model.IsArray {
			return GetParser(model.Name, "ParseTimeArrayFromRequest"), nil
		}
		if model.IsPointer {
			return GetParser(model.Name, "ParseTimePointerFromRequest"), nil
		}
		return GetParser(model.Name, "ParseTimeFromRequest"), nil

	default:
		// TODO: implement model types
		return nil, errors.New("Unknown type: " + model.Type)
	}
}

func GetParser(fieldName, parserName string) *dst.AssignStmt {
	return &dst.AssignStmt{
		Lhs: []dst.Expr{&dst.SelectorExpr{X: GetName("filter"), Sel: GetName(fieldName)}},
		Tok: token.ASSIGN,
		Rhs: []dst.Expr{&dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   GetName("common"),
				Sel: GetName(parserName),
			},
			Args: []dst.Expr{
				GetName("request"),
				GetName(WrapString(fieldName)),
			},
		}},
	}
}
