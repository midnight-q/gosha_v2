package utils

import (
	"fmt"
	"go/token"
	"strconv"
	"strings"

	"github.com/dave/dst"
)

func GetNames(name string) []*dst.Ident {
	return []*dst.Ident{{Name: name}}
}

func GetName(name string) *dst.Ident {
	return &dst.Ident{Name: name}
}

func GetComment(text string) dst.NodeDecs {
	if len(text) < 1 {
		return dst.NodeDecs{}
	}
	commentList := []string{}
	for _, s := range strings.Split(text, "\n") {
		if len(s) > 0 {
			commentList = append(commentList, fmt.Sprintf("// %s", s))
		}
	}
	return dst.NodeDecs{
		Before: dst.NewLine,
		Start:  commentList,
	}
}

func WrapString(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}

func GetNewLineDecorations() dst.NodeDecs {
	return dst.NodeDecs{
		Before: dst.NewLine,
	}
}

func GetBlankIdentifier() *dst.Ident {
	return &dst.Ident{
		Name: "_",
	}
}

func GetIntValue(in int) *dst.BasicLit {
	return &dst.BasicLit{
		Kind:  token.INT,
		Value: strconv.Itoa(in),
	}
}

func GetStringValue(in string) *dst.BasicLit {
	return &dst.BasicLit{
		Kind:  token.STRING,
		Value: WrapString(in),
	}
}

func GetConst(name, value string) *dst.GenDecl {

	return &dst.GenDecl{
		Tok:    token.CONST,
		Lparen: false,
		Specs: []dst.Spec{
			&dst.ValueSpec{
				Names: GetNames(name),
				Values: []dst.Expr{
					GetStringValue(value),
				},
			},
		},
		Rparen: false,
		Decs: dst.GenDeclDecorations{
			NodeDecs: GetNewLineDecorations(),
		},
	}
}
