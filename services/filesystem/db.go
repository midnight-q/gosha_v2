package filesystem

import (
	"fmt"
	"gosha_v2/errors"
	"gosha_v2/services/utils"
	"gosha_v2/settings"
	"io/fs"
	"io/ioutil"

	"github.com/dave/dst"
)

func UpdateDbConnection(currentPath string, dbTypeId int) error {
	switch dbTypeId {
	case settings.PostgresqlDbTypeId:

	default:
		return errors.NewErrorWithCode("Unsupported DatabaseType", errors.ErrorCodeNotFound, "DatabaseType")
	}

	appName := utils.GetNameForNewApp(currentPath)
	dbPass := utils.GeneratePassword()
	dbPort := "35432"
	fileName := currentPath + "/settings/db.go"

	file, err := readFile(fileName)
	if err != nil {
		return err
	}

	isDbUserFind := false
	isDbPassFind := false
	isDbNameFind := false
	isDbPortFind := false

	for _, decl := range file.Decls {
		constDecl, isOk := decl.(*dst.GenDecl)
		if !isOk {
			continue
		}
		if len(constDecl.Specs) != 1 {
			continue
		}

		spec, isOk := constDecl.Specs[0].(*dst.ValueSpec)
		if !isOk {
			continue
		}
		if len(spec.Names) != 1 {
			continue
		}
		switch spec.Names[0].Name {
		case "DbUser":
			if len(spec.Values) != 1 {
				continue
			}
			val, isOk := spec.Values[0].(*dst.BasicLit)
			if !isOk {
				continue
			}
			val.Value = utils.WrapString(appName)
			isDbUserFind = true
		case "DbPass":
			if len(spec.Values) != 1 {
				continue
			}
			val, isOk := spec.Values[0].(*dst.BasicLit)
			if !isOk {
				continue
			}
			val.Value = utils.WrapString(dbPass)
			isDbPassFind = true
		case "DbName":
			if len(spec.Values) != 1 {
				continue
			}
			val, isOk := spec.Values[0].(*dst.BasicLit)
			if !isOk {
				continue
			}
			val.Value = utils.WrapString(appName)
			isDbNameFind = true
		case "DbPort":
			if len(spec.Values) != 1 {
				continue
			}
			val, isOk := spec.Values[0].(*dst.BasicLit)
			if !isOk {
				continue
			}
			val.Value = utils.WrapString(dbPort)
			isDbPortFind = true
		default:
			continue
		}
	}

	if !isDbUserFind {
		err = fmt.Errorf("Not found DbUser const")
		return err
	}
	if !isDbPassFind {
		err = fmt.Errorf("Not found DbPass const")
		return err
	}
	if !isDbNameFind {
		err = fmt.Errorf("Not found DbName const")
		return err
	}
	if !isDbPortFind {
		err = fmt.Errorf("Not found DbPort const")
		return err
	}

	err = saveFile(file, fileName)
	if err != nil {
		return err
	}

	return updateDbScript(currentPath, dbTypeId, appName, dbPass, dbPort)
}

const PgScript = `
#!/bin/bash
docker run --rm --name pg-%s -e POSTGRES_DB=%s -e POSTGRES_USER=%s -e POSTGRES_PASSWORD=%s -d -p %s:5432 -v "$(pwd)/.postgres:/var/lib/postgresql/data" postgres
`

func updateDbScript(currentPath string, dbTypeId int, appName, dbPass, dbPort string) error {
	scriptString := ""
	switch dbTypeId {
	case settings.PostgresqlDbTypeId:
		scriptString = fmt.Sprintf(PgScript, appName, appName, appName, dbPass, dbPort)
	default:
		return errors.NewErrorWithCode("Unsupported DatabaseType", errors.ErrorCodeNotFound, "DatabaseType")
	}

	err := ioutil.WriteFile(currentPath+"/db-docker.sh", []byte(scriptString), fs.ModePerm)
	return err
}
