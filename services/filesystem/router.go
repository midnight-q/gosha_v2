package filesystem

import (
	"fmt"
	"go/token"
	"gosha_v2/common"
	"gosha_v2/services/utils"
	"gosha_v2/types"

	"github.com/dave/dst"
)

func CheckRoutesAvailability(path string, models []types.Model) (res []types.Model, err error) {
	routerFuncNames := []string{}

	filePath := path + "/router/router.go"

	file, err := readFile(filePath)
	if err != nil {
		return nil, err
	}

	for _, decl := range file.Decls {
		funcDecl, isOk := decl.(*dst.FuncDecl)
		if !isOk {
			continue
		}
		if funcDecl.Name.Name != "Router" {
			continue
		}

		for _, stmt := range funcDecl.Body.List {
			exprStmt, isOk := stmt.(*dst.ExprStmt)
			if !isOk {
				continue
			}
			callExpr, isOk := exprStmt.X.(*dst.CallExpr)
			if !isOk {
				continue
			}
			selectorExpr, isOk := callExpr.Fun.(*dst.SelectorExpr)
			if !isOk {
				continue
			}

			callExpr, isOk = selectorExpr.X.(*dst.CallExpr)
			if !isOk {
				continue
			}

			if len(callExpr.Args) < 2 {
				continue
			}

			selectorExpr, isOk = callExpr.Args[1].(*dst.SelectorExpr)
			if !isOk {
				continue
			}

			routerFuncNames = append(routerFuncNames, selectorExpr.Sel.Name)
		}
	}

	for _, model := range models {
		if model.IsServiceModel {
			res = append(res, model)
			continue
		}

		model.HttpMethods.Find = common.CheckInArray(model.Name+"Find", routerFuncNames)
		model.HttpMethods.Create = common.CheckInArray(model.Name+"Create", routerFuncNames)
		model.HttpMethods.MultiCreate = common.CheckInArray(model.Name+"MultiCreate", routerFuncNames)
		model.HttpMethods.Read = common.CheckInArray(model.Name+"Read", routerFuncNames)
		model.HttpMethods.Update = common.CheckInArray(model.Name+"Update", routerFuncNames)
		model.HttpMethods.MultiUpdate = common.CheckInArray(model.Name+"MultiUpdate", routerFuncNames)
		model.HttpMethods.Delete = common.CheckInArray(model.Name+"Delete", routerFuncNames)
		model.HttpMethods.MultiDelete = common.CheckInArray(model.Name+"MultiDelete", routerFuncNames)
		model.HttpMethods.FindOrCreate = common.CheckInArray(model.Name+"FindOrCreate", routerFuncNames)
		model.HttpMethods.UpdateOrCreate = common.CheckInArray(model.Name+"UpdateOrCreate", routerFuncNames)
		res = append(res, model)
	}

	return
}

func AddRouteInSettings(currentPath string, modelName string) (err error) {
	route := utils.GenerateRouteForModel(modelName)
	constName := modelName + "Route"
	filePath := currentPath + "/settings/routes.go"

	file, err := readFile(filePath)
	if err != nil {
		return err
	}
	index := 0
	for i, decl := range file.Decls {
		genDecl, isOk := decl.(*dst.GenDecl)
		if !isOk {
			continue
		}
		if genDecl.Tok != token.VAR {
			continue
		}
		if len(genDecl.Specs) < 1 {
			continue
		}

		valueSpec, isOk := genDecl.Specs[0].(*dst.ValueSpec)
		if !isOk {
			continue
		}
		if len(valueSpec.Names) < 1 {
			continue
		}
		if valueSpec.Names[0].Name != "RoutesArray" {
			continue
		}
		if len(valueSpec.Values) < 1 {
			continue
		}

		compositLit, isOk := valueSpec.Values[0].(*dst.CompositeLit)
		if !isOk {
			continue
		}
		newIdent := utils.GetName(constName)
		newIdent.Decs.After = dst.NewLine
		compositLit.Elts = append(compositLit.Elts, newIdent)
		index = i
	}
	newConst := utils.GetConst(constName, route)

	file.Decls = append(file.Decls, &dst.GenDecl{})
	copy(file.Decls[index+1:], file.Decls[index:])
	file.Decls[index] = newConst

	return saveFile(file, filePath)
}

func RegisterNewRoute(currentPath string, model types.Model) (err error) {
	filePath := currentPath + "/router/router.go"

	file, err := readFile(filePath)

	index := 0
	for _, decl := range file.Decls {
		funcDecl, isOk := decl.(*dst.FuncDecl)
		if !isOk {
			continue
		}
		if funcDecl.Name.Name != "Router" {
			continue
		}
		for i, stmt := range funcDecl.Body.List {
			_, isAssign := stmt.(*dst.AssignStmt)
			_, isReturn := stmt.(*dst.ReturnStmt)
			if isAssign || isReturn {
				index = i
				break
			}
			exprStmt, isOk := stmt.(*dst.ExprStmt)
			if !isOk {
				continue
			}
			callExpr, isOk := exprStmt.X.(*dst.CallExpr)
			if !isOk {
				continue
			}
			selectorExpr, isOk := callExpr.Fun.(*dst.SelectorExpr)
			if !isOk {
				continue
			}
			utils.DebugPrintf(selectorExpr.X)
			break
		}
	}
	fmt.Println(index)
	return nil
}
