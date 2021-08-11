package utils

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

func GetName(name string) []*dst.Ident {
	return []*dst.Ident{{Name: name}}
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
		},
	}
}

func WrapString(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}
