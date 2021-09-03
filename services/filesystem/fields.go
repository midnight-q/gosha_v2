package filesystem

import (
	"fmt"
	"go/token"
	"gosha_v2/errors"
	"gosha_v2/services/utils"
	"gosha_v2/types"
	"strings"

	"github.com/dave/dst"
)

func AddFieldInModel(fieldName, fieldComment, modelName, fileName, typeName string, fieldType dst.Expr) error {
	file, err := readFile(fileName)
	if err != nil {
		return err
	}

	if typeName == "uuid" || typeName == "time" {
		utils.AddImportIfNeeded(typeName, file)
	}

	for _, decl := range file.Decls {
		modelDecl, isOk := decl.(*dst.GenDecl)
		if !isOk {
			continue
		}

		if modelDecl.Tok != token.TYPE {
			continue
		}
		typeSpec := modelDecl.Specs[0].(*dst.TypeSpec)

		if typeSpec.Name.Name != modelName {
			continue
		}

		structType := typeSpec.Type.(*dst.StructType)

		structType.Fields.List = addFieldInList(structType.Fields.List, fieldName, fieldType, fieldComment)
		if err != nil {
			return err
		}

		return saveFile(file, fileName)
	}

	return errors.New(fmt.Sprintf("Not found model %s in %s", modelName, fileName))
}

func addFieldInList(fields []*dst.Field, name string, fieldType dst.Expr, comment string) []*dst.Field {
	index := 0

	for i, field := range fields {
		if len(field.Names) < 1 {
			index = i
			break
		}
		if field.Names[0].Name == "CreatedAt" {
			index = i
			break
		}
	}
	fields = append(fields, &dst.Field{})
	copy(fields[index+1:], fields[index:])

	fields[index] = &dst.Field{
		Names: utils.GetNames(name),
		Type:  fieldType,
		Decs: dst.FieldDecorations{
			NodeDecs: utils.GetComment(comment),
		},
	}

	return fields
}

func AddParserInFilter(model types.Field, path string, appName string) error {
	file, err := readFile(path)
	if err != nil {
		return err
	}

	utils.AddImportIfNeeded(appName+"/common", file)

	for _, decl := range file.Decls {
		funcDecl, isOk := decl.(*dst.FuncDecl)
		if !isOk {
			continue
		}
		if !strings.HasPrefix(funcDecl.Name.Name, "Get") || !strings.HasSuffix(funcDecl.Name.Name, "Filter") {
			continue
		}

		err = addParserInFunc(funcDecl, model)
		if err != nil {
			return err
		}
	}

	return saveFile(file, path)
}

func addParserInFunc(funcDecl *dst.FuncDecl, model types.Field) error {
	parserExpr, err := utils.GetParserForType(model)
	if err != nil {
		return err
	}

	index := 0
	for i, stmt := range funcDecl.Body.List {
		_, isOk := stmt.(*dst.SwitchStmt)
		if isOk {
			index = i
			break
		}
		_, isOk = stmt.(*dst.ReturnStmt)
		if isOk {
			index = i
			break
		}
	}

	funcDecl.Body.List = append(funcDecl.Body.List, &dst.AssignStmt{})
	copy(funcDecl.Body.List[index+1:], funcDecl.Body.List[index:])

	funcDecl.Body.List[index] = parserExpr

	return nil
}
