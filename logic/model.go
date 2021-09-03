package logic

import (
	"gosha_v2/errors"

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
			TypePath:       typeModel.Path,
			DbPath:         dbModel.Path,
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
	newModel := filter.GetModelModel()
	currentDir, err := filter.GetPwd()
	if err != nil {
		return types.Model{}, err
	}

	appName, err := getAppName(currentDir)
	if err != nil {
		return types.Model{}, err
	}
	filter.IsShowServiceModels = true
	models, _, err := ModelFind(filter)
	if err != nil {
		return
	}
	for _, model := range models {
		if model.Name == newModel.Name {
			err = errors.New("Model already exist")
			return
		}
	}

	if !newModel.IsTypeModel && newModel.IsDbModel {
		err = errors.New("Creating only db models is not allow")
		return
	}

	if newModel.IsTypeModel {
		err = filesystem.CopyNewModelFile(currentDir, "/types/", newModel.Name, appName)
		if err != nil {
			return types.Model{}, err
		}
		err = filesystem.AddCommentToModel(currentDir, "/types/", newModel.Name, newModel.CommentType)
		if err != nil {
			return types.Model{}, err
		}
	}
	if newModel.IsDbModel {
		err = filesystem.CopyNewModelFile(currentDir, "/dbmodels/", newModel.Name, appName)
		if err != nil {
			return types.Model{}, err
		}
		err = filesystem.AddCommentToModel(currentDir, "/dbmodels/", newModel.Name, newModel.CommentDb)
		if err != nil {
			return types.Model{}, err
		}
	}

	// create logic
	err = filesystem.CopyNewModelFile(currentDir, "/logic/", newModel.Name, appName)
	if err != nil {
		return types.Model{}, err
	}

	if newModel.IsTypeModel && !newModel.IsDbModel {
		// If model is virtual - remove logic body and assigners
		err = filesystem.ClearLogic(currentDir, newModel.Name)
		if err != nil {
			return types.Model{}, err
		}
	}

	// add fields
	for _, field := range newModel.Fields {
		fField := types.FieldFilter{}
		fField.SetPwd(currentDir)
		field.ModelName = newModel.Name
		fField.SetFieldModel(field)
		_, err = FieldCreate(fField)
		if err != nil {
			return types.Model{}, err
		}
	}

	// create webapp
	err = filesystem.CopyNewModelFile(currentDir, "/webapp/", newModel.Name, appName)
	if err != nil {
		return types.Model{}, err
	}

	// create route in settings
	err = filesystem.AddRouteInSettings(currentDir, newModel.Name)
	if err != nil {
		return types.Model{}, err
	}

	// register routes in router
	err = filesystem.RegisterNewRoute(currentDir, newModel)

	return
}

func getAppName(currentDir string) (name string, err error) {
	fApplication := types.ApplicationFilter{}
	fApplication.CurrentPage = 1
	fApplication.PerPage = 1
	fApplication.SetPwd(currentDir)

	apps, _, err := ApplicationFind(fApplication)
	if err != nil {
		return "", err
	}

	if len(apps) < 1 {
		err = errors.NewErrorWithCode("Not found app", errors.ErrorCodeNotFound, "")
		return "", err
	}

	return apps[0].Name, nil
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
