package utils

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"
)

func GetNames(name string) []*dst.Ident {
	return []*dst.Ident{{Name: name}}
}

func GetName(name string) *dst.Ident {
	return &dst.Ident{Name: name}
}

func GetComment(text string) dst.FieldDecorations {
	if len(text) < 1 {
		return dst.FieldDecorations{}
	}
	commentList := []string{}
	for _, s := range strings.Split(text, "\n") {
		if len(s) > 0 {
			commentList = append(commentList, fmt.Sprintf("// %s", s))
		}
	}
	return dst.FieldDecorations{
		NodeDecs: dst.NodeDecs{
			Before: dst.NewLine,
			Start:  commentList,
			After:  dst.NewLine,
		},
	}
}

func WrapString(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}

func GetNewLineDecorations() dst.KeyValueExprDecorations {

	return dst.KeyValueExprDecorations{
		NodeDecs: dst.NodeDecs{
			Before: dst.NewLine,
			After:  dst.NewLine,
		},
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
		Value: fmt.Sprintf("%d", in),
	}
}
