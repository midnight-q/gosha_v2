package filesystem

import (
	"fmt"
	"go/token"
	"gosha_v2/errors"
	"gosha_v2/services/utils"

	"github.com/dave/dst"
)

func AddFieldInModel(fieldName, fieldComment, modelName, fileName, typeName string, fieldType dst.Expr) error {
	file, err := readFile(fileName)
	if err != nil {
		return err
	}

	utils.AddImportForTypeIfNeeded(typeName, file)

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
		Names: utils.GetName(name),
		Type:  fieldType,
		Decs:  utils.GetComment(comment),
	}

	return fields
}
