package filesystem

import (
	"gosha_v2/common"
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
