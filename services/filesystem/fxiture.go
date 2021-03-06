package filesystem

import (
	"fmt"
	"gosha_v2/services/utils"
	"gosha_v2/types"

	"github.com/dave/dst"
)

func UpdateUserFixtures(currentPath, newSalt string, app types.Application) error {

	fileName := currentPath + "/bootstrap/fixtures/user.go"

	file, err := readFile(fileName)
	if err != nil {
		return err
	}

	IsPasswordFind := false
	IsEmailFind := false

MainLoop:
	for _, decl := range file.Decls {
		funcDecl, isOk := decl.(*dst.FuncDecl)
		if !isOk {
			continue
		}

		if funcDecl.Name.Name != "AddUser" {
			continue
		}

		for _, stmt := range funcDecl.Body.List {
			ifStmt, isOk := stmt.(*dst.IfStmt)
			if !isOk {
				continue
			}

			for _, stmt := range ifStmt.Body.List {
				assignStmt, isOk := stmt.(*dst.AssignStmt)
				if !isOk {
					continue
				}
				if len(assignStmt.Rhs) != 1 {
					continue
				}

				rhs, isOk := assignStmt.Rhs[0].(*dst.CompositeLit)
				if !isOk {
					continue
				}

				for _, elt := range rhs.Elts {
					keyValueExpr, isOk := elt.(*dst.KeyValueExpr)
					if !isOk {
						continue
					}

					key, isOk := keyValueExpr.Key.(*dst.Ident)
					if !isOk {
						continue
					}

					value, isOk := keyValueExpr.Value.(*dst.BasicLit)
					if !isOk {
						continue
					}

					switch key.Name {
					case "Password":
						value.Value = utils.WrapString(utils.HashPassword(newSalt, app.Password))
						IsPasswordFind = true
					case "Email":
						value.Value = utils.WrapString(app.Email)
						IsEmailFind = true
					}
				}

				break MainLoop
			}
		}
	}

	if !IsEmailFind {
		return fmt.Errorf("not found email field")
	}
	if !IsPasswordFind {
		return fmt.Errorf("not found password field")
	}

	err = saveFile(file, fileName)
	if err != nil {
		return err
	}

	return nil
}
