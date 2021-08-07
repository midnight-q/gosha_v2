package logic

import (
	"gosha_v2/services/filesystem"
	"gosha_v2/services/utils"
	"gosha_v2/types"
	"strings"
)

func ModelFind(filter types.ModelFilter) (result []types.Model, totalRecords int, err error) {
	path, err := filter.GetPwd()
	if err != nil {
		return nil, 0, err
	}
	// load dbmodels
	dbModels, err := filesystem.LoadDbModels(path)
	if err != nil {
		return nil, 0, err
	}

	// load type models
	typeModels, err := filesystem.LoadTypeModels(path)
	if err != nil {
		return nil, 0, err
	}

	// Merge models for result
	result, err = mergeModels(typeModels, dbModels)
	if err != nil {
		return nil, 0, err
	}

	//result, err = utils.GenerateRoutesForModels(result)
	//if err != nil {
	//	return nil, 0, err
	//}
	//result, err = filesystem.CheckRoutesAvailability(result)
	//if err != nil {
	//	return nil, 0, err
	//}

	return result, len(result), nil
}

func mergeModels(typeModels, dbModels []filesystem.Model) (res []types.Model, err error) {
	for _, typeModel := range typeModels {
		if typeModel.IsFilter {
			continue
		}

		dbModel, isFindDbModel := findModel(typeModel.Name, dbModels)
		filter, isFindFilter := findModel(typeModel.Name+"Filter", typeModels)

		model := types.Model{
			Name:        typeModel.Name,
			Fields:      []types.Field{},
			IsTypeModel: true,
			IsDbModel:   isFindDbModel,
			PkType:      getPkType(dbModel),
			CommentType: typeModel.Comment,
			CommentDb:   dbModel.Comment,
		}

		model.Fields = mergeFields(typeModel.Fields, dbModel.Fields)

		if isFindFilter {
			model.Filter = &types.Model{
				Name:        filter.Name,
				Fields:      []types.Field{},
				IsTypeModel: true,
				CommentType: filter.Comment,
			}
			for _, field := range filter.Fields {
				model.Filter.Fields = append(model.Filter.Fields, types.Field{
					Name:        field.Name,
					Type:        field.Type,
					CommentType: field.Comment,
					IsTypeField: true,
				})
			}
		}

		res = append(res, model)
	}
	return
}

func mergeFields(typeField []filesystem.Field, dbField []filesystem.Field) (res []types.Field) {
	fieldsMap := map[string]types.Field{}
	for _, field := range typeField {
		fieldsMap[strings.ToLower(field.Name)] = types.Field{
			Name:        field.Name,
			Type:        field.Type,
			CommentType: field.Comment,
			IsTypeField: true,
		}
	}

	for _, dbField := range dbField {
		field, isFind := fieldsMap[strings.ToLower(dbField.Name)]
		if !isFind {
			field = types.Field{
				Name: dbField.Name,
				Type: dbField.Type,
			}
		}

		field.IsDbField = true
		field.CommentDb = dbField.Comment
		fieldsMap[strings.ToLower(dbField.Name)] = field
	}

	for _, field := range fieldsMap {
		res = append(res, field)
	}
	return
}

func findModel(name string, models []filesystem.Model) (filesystem.Model, bool) {
	for _, model := range models {
		if model.Name == name {
			return model, true
		}
	}
	return filesystem.Model{}, false
}

func getPkType(model filesystem.Model) string {
	for _, field := range model.Fields {
		if utils.IsIdField(field.Name) {
			return field.Type
		}
	}

	return ""
}

func ModelMultiCreate(filter types.ModelFilter) (data []types.Model, err error) {

	return
}

func ModelCreate(filter types.ModelFilter) (data types.Model, err error) {
	return
}

func ModelRead(filter types.ModelFilter) (data types.Model, err error) {
	return
}

func ModelMultiUpdate(filter types.ModelFilter) (data []types.Model, err error) {
	return
}

func ModelUpdate(filter types.ModelFilter) (data types.Model, err error) {
	return
}

func ModelMultiDelete(filter types.ModelFilter) (isOk bool, err error) {
	return
}

func ModelDelete(filter types.ModelFilter) (isOk bool, err error) {
	return
}

func ModelFindOrCreate(filter types.ModelFilter) (data types.Model, err error) {
	return
}

func ModelUpdateOrCreate(filter types.ModelFilter) (data types.Model, err error) {
	return
}

// add all assign functions
