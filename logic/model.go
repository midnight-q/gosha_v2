package logic

import (
	"github.com/elliotchance/orderedmap"

	"gosha_v2/services/filesystem"
	"gosha_v2/services/utils"
	"gosha_v2/types"
	"strings"
)

func ModelFind(filter types.ModelFilter) (result []types.Model, totalRecords int, err error) {
	currentPath, err := filter.GetPwd()
	if err != nil {
		return nil, 0, err
	}

	// load dbmodels
	dbModels, err := filesystem.LoadDbModels(currentPath)
	if err != nil {
		return nil, 0, err
	}

	// load type models
	typeModels, err := filesystem.LoadTypeModels(currentPath)
	if err != nil {
		return nil, 0, err
	}

	// Merge models for result
	result, err = mergeModels(typeModels, dbModels)
	if err != nil {
		return nil, 0, err
	}

	result, err = utils.GenerateRoutesForModels(result)
	if err != nil {
		return nil, 0, err
	}

	result, err = filesystem.CheckRoutesAvailability(currentPath, result)
	if err != nil {
		return nil, 0, err
	}

	if !filter.IsShowServiceModels {
		result = utils.FilterServiceModels(result)
	}

	return result, len(result), nil
}

func mergeModels(typeModels, dbModels []filesystem.Model) (res []types.Model, err error) {
	for _, typeModel := range typeModels {
		if typeModel.IsFilter && !typeModel.IsServiceModel {
			continue
		}

		dbModel, isFindDbModel := filesystem.FindModel(typeModel.Name, dbModels)
		filter, isFindFilter := filesystem.FindModel(typeModel.Name+"Filter", typeModels)

		model := types.Model{
			Name:           typeModel.Name,
			Fields:         []types.Field{},
			IsTypeModel:    true,
			IsDbModel:      isFindDbModel,
			PkType:         getPkType(dbModel),
			CommentType:    typeModel.Comment,
			CommentDb:      dbModel.Comment,
			IsServiceModel: typeModel.IsServiceModel,
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
					SourceModel: field.SourceModel,
				})
			}
		}

		res = append(res, model)
	}
	return
}

func mergeFields(typeField []filesystem.Field, dbField []filesystem.Field) (res []types.Field) {
	fieldsMap := orderedmap.NewOrderedMap()
	for _, field := range typeField {
		fieldsMap.Set(strings.ToLower(field.Name), types.Field{
			Name:        field.Name,
			Type:        field.Type,
			CommentType: field.Comment,
			IsTypeField: true,
			SourceModel: field.SourceModel,
		})
	}

	for _, dbField := range dbField {
		rawField, isFind := fieldsMap.Get(strings.ToLower(dbField.Name))
		if !isFind {
			fieldsMap.Set(strings.ToLower(dbField.Name), types.Field{
				Name:        dbField.Name,
				Type:        dbField.Type,
				IsDbField:   true,
				CommentDb:   dbField.Comment,
				SourceModel: dbField.SourceModel,
			})
		} else {
			field := rawField.(types.Field)
			field.IsDbField = true
			field.CommentDb = dbField.Comment
			fieldsMap.Set(strings.ToLower(dbField.Name), field)
		}

	}

	for el := fieldsMap.Front(); el != nil; el = el.Next() {
		res = append(res, el.Value.(types.Field))
	}
	return
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
