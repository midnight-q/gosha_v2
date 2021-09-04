package sdk

import (
	"embed"
	"fmt"
	"gosha_v2/services/utils"
	"gosha_v2/settings"
	"gosha_v2/types"
	"strings"
)

func GetModelsFile(models []types.Model) []byte {
	content := `
	//some common actions
	`

	for _, model := range models {
		if !model.IsTypeModel {
			continue
		}
		content += fmt.Sprintf("\nexport function %s() {\n", model.Name)

		content += getFieldsBody(model.Fields, models)

		content += "\n\treturn this;\n}\n"

		if model.Filter != nil {
			content += fmt.Sprintf("\nexport function %s() {\n", model.Filter.Name)

			content += getFieldsBody(model.Filter.Fields, models)

			content += "\n\treturn this;\n}\n"
		}
	}

	return []byte(content)
}

func GetSkeletonTemplateFile(fs embed.FS, filePath string) []byte {
	res, err := fs.ReadFile(settings.SkeletonSdkPath + filePath)
	if err != nil {
		return nil
	}
	return res
}

type data struct {
	Name    string
	Content string
}

func GetStores(fs embed.FS, models []types.Model) (res []data) {
	for _, model := range models {
		if !model.IsTypeModel {
			continue
		}
		b := GetSkeletonTemplateFile(fs, "/jstypes/store.js")
		content := string(b)

		content = strings.Replace(content, "EntityName", model.Name, -1)

		content = strings.Replace(content, "entityName", utils.GetFirstLowerCase(model.Name), -1)

		res = append(res, data{
			Name:    model.Name,
			Content: content,
		})
	}
	return
}

func GetData(fs embed.FS, models []types.Model) (res []data) {
	for _, model := range models {
		if !model.IsTypeModel {
			continue
		}
		b := GetSkeletonTemplateFile(fs, "/jstypes/data.js")
		content := string(b)

		content = strings.Replace(content, "EntityName", model.Name, -1)

		content = strings.Replace(content, "entityName", utils.GetFirstLowerCase(model.Name), -1)

		res = append(res, data{
			Name:    model.Name,
			Content: content,
		})
	}
	return
}

func getFieldsBody(fields []types.Field, models []types.Model) string {
	arr := []string{}
	for _, field := range fields {
		if !field.IsTypeField {
			continue
		}
		arr = append(arr, fmt.Sprintf("\t%s = %s", field.Name, getFieldType(field, models)))
	}
	return strings.Join(arr, "\n")
}

func getFieldType(field types.Field, models []types.Model) string {
	if field.IsPointer {
		return "null"
	}
	if field.IsArray {
		return "[]"
	}

	switch field.Type {
	case "int":
		return "0"
	case "float64":
		return "0.0"
	case "string":
		return `""`
	case "byte":
		return `""`
	case "time":
		return "null"
	case "uuid":
		return "null"
	case "bool":
		return "false"
	default:
		for _, model := range models {
			if field.Type == model.Name {
				return fmt.Sprintf("new %s()", field.Type)
			}
		}
		return "{}"
	}
}
