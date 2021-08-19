package filesystem

import (
	"fmt"
	"go/token"
	"gosha_v2/services/utils"
	"gosha_v2/types"
	"strings"

	"github.com/dave/dst"
)

func AddFieldInAssigner(model types.Field, currentPath string) (err error) {
	logicPath := currentPath + "/logic/" + utils.GetFileNameFromModelName(model.ModelName) + ".go"

	file, err := readFile(logicPath)

	for _, decl := range file.Decls {
		funcDecl, isOk := decl.(*dst.FuncDecl)
		if !isOk {
			continue
		}

		funcName := funcDecl.Name.Name

		if !strings.HasPrefix(funcName, "Assign") {
			continue
		}

		inputModelName, err := getInputModelName(funcDecl)
		if err != nil {
			return err
		}

		for _, stmt := range funcDecl.Body.List {
			switch typedStmt := stmt.(type) {
			case *dst.ReturnStmt:
				compositeLit, isOk := typedStmt.Results[0].(*dst.CompositeLit)
				if !isOk {
					continue
				}
				if addFieldInCompositeLit(compositeLit, model, inputModelName) != nil {
					continue
				}

			case *dst.AssignStmt:
				if typedStmt.Tok != token.DEFINE && typedStmt.Tok != token.ASSIGN {
					continue
				}
				for _, rh := range typedStmt.Rhs {
					compositeLit, isOk := rh.(*dst.CompositeLit)
					if !isOk {
						continue
					}
					if addFieldInCompositeLit(compositeLit, model, inputModelName) != nil {
						continue
					}
				}
			}
		}
	}
	return saveFile(file, logicPath)
}

func getInputModelName(funcDecl *dst.FuncDecl) (string, error) {
	if len(funcDecl.Type.Params.List) < 1 {
		return "", fmt.Errorf("not found input model name")
	}
	if len(funcDecl.Type.Params.List[0].Names) < 1 {
		return "", fmt.Errorf("not found input model name")
	}
	return funcDecl.Type.Params.List[0].Names[0].Name, nil
}

func addFieldInCompositeLit(compositeLit *dst.CompositeLit, model types.Field, inputModelName string) error {
	selectorExpr, isOk := compositeLit.Type.(*dst.SelectorExpr)
	if !isOk {
		return fmt.Errorf("pass")
	}
	if selectorExpr.Sel.Name != model.ModelName {
		return fmt.Errorf("pass")
	}
	compositeLit.Elts = append(compositeLit.Elts, &dst.KeyValueExpr{
		Key: utils.GetName(model.Name),
		Value: &dst.SelectorExpr{
			X:   utils.GetName(inputModelName),
			Sel: utils.GetName(model.Name),
		},
		Decs: utils.GetNewLineDecorations(),
	})
	return nil
}
