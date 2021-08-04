package main

import (
	"fmt"
	"go/token"
	"gosha_v2/utils"
	"io/ioutil"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

func ososososososo() {
	filePath := "./skeleton_app/dbmodels/user_role.go"
	b, err := ioutil.ReadFile(filePath) // just pass the file name

	if err != nil {
		fmt.Print(err)
	}
	src := string(b) // convert content to a 'string'

	file, err := decorator.Parse(src)
	if err != nil {
		fmt.Println("Error in decorator.Parse:", err.Error())
	}

	for _, decl := range file.Decls {
		modelDecl, isOk := decl.(*dst.GenDecl)
		if !isOk {
			continue
		}

		if modelDecl.Tok == token.TYPE {
			typeSpec := modelDecl.Specs[0].(*dst.TypeSpec)

			structType := typeSpec.Type.(*dst.StructType)

			for _, field := range structType.Fields.List {
				//fmt.Printf("%T %+v\n", field.Type, field.Type)
				fmt.Printf("%+v\n", field.Tag)
			}

			structType.Fields.List = append(structType.Fields.List, &dst.Field{
				Names: utils.GetName("TestField"),
				Type:  utils.GetStringType(),
				Tag: &dst.BasicLit{
					Kind:  token.STRING,
					Value: "`gorm:\"primary_key\"`",
				},
				Decs: utils.GetComment("It's comment\n"),
			})
		}

		//fmt.Printf("Type: %T\n", decl)
	}

	err = decorator.Print(file)
	if err != nil {
		fmt.Println("Error in decorator.Print:", err)
		return
	}
}
