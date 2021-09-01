package logic

import (
	"gosha_v2/errors"
	"gosha_v2/services/filesystem"
	"gosha_v2/services/utils"
	"gosha_v2/types"
)

func FieldFind(filter types.FieldFilter) (result []types.Field, totalRecords int, err error) {

	return
}

func FieldMultiCreate(filter types.FieldFilter) (data []types.Field, err error) {

	return
}

func FieldCreate(filter types.FieldFilter) (data types.Field, err error) {
	fieldModel := filter.GetFieldModel()

	if !fieldModel.IsValid() {
		err = fieldModel.GetValidationError()
		return types.Field{}, err
	}

	currentPath, err := filter.GetPwd()
	if err != nil {
		return types.Field{}, err
	}

	appName := filesystem.FindAppName(currentPath)

	fieldType, err := utils.GetType(fieldModel.Type, fieldModel.IsArray, fieldModel.IsPointer)
	if err != nil {
		return types.Field{}, err
	}

	fModel := types.ModelFilter{}
	models, _, err := ModelFind(fModel)
	if err != nil {
		return types.Field{}, err
	}

	model, isFilter, err := utils.GetModel(fieldModel.ModelName, models)
	if err != nil {
		return types.Field{}, err
	}

	if utils.IsFieldExistInModel(fieldModel.Name, model, isFilter) {
		return types.Field{}, errors.New("Field already exist in model: " + fieldModel.Name)
	}

	if isFilter {
		if fieldModel.IsDbField {
			err = errors.New("Cant create dbField in filter")
			return types.Field{}, err
		}

		if fieldModel.IsPointer && fieldModel.IsArray {
			err = errors.New("Filter field must be array OR pointer. Not both at the same time")
			return types.Field{}, err
		}

		err = filesystem.AddFieldInModel(fieldModel.Name, fieldModel.CommentType, fieldModel.ModelName, model.TypePath, fieldModel.Type, fieldType)
		if err != nil {
			return types.Field{}, err
		}

		err = filesystem.AddParserInFilter(fieldModel, model.TypePath, appName)
		if err != nil {
			return types.Field{}, err
		}

	} else {
		if !model.IsDbModel && fieldModel.IsDbField {
			err = errors.New("Cant create dbField because dbModel not exist")
			return types.Field{}, err
		}
		if !model.IsTypeModel && fieldModel.IsTypeField {
			err = errors.New("Cant create typeField because typeModel not exist")
			return types.Field{}, err
		}

		if fieldModel.IsDbField {
			err = filesystem.AddFieldInModel(fieldModel.Name, fieldModel.CommentDb, fieldModel.ModelName, model.DbPath, fieldModel.Type, fieldType)
			if err != nil {
				return types.Field{}, err
			}
		}
		if fieldModel.IsTypeField {
			err = filesystem.AddFieldInModel(fieldModel.Name, fieldModel.CommentType, fieldModel.ModelName, model.TypePath, fieldModel.Type, fieldType)
			if err != nil {
				return types.Field{}, err
			}
		}
		if fieldModel.IsDbField && fieldModel.IsTypeField {
			// TODO: Add field in Update
			err = filesystem.AddFieldInAssigner(fieldModel, currentPath)
			if err != nil {
				return types.Field{}, err
			}
		}
	}

	return
}

func FieldRead(filter types.FieldFilter) (data types.Field, err error) {
	return
}

func FieldMultiUpdate(filter types.FieldFilter) (data []types.Field, err error) {
	return
}

func FieldUpdate(filter types.FieldFilter) (data types.Field, err error) {
	return
}

func FieldMultiDelete(filter types.FieldFilter) (isOk bool, err error) {
	return
}

func FieldDelete(filter types.FieldFilter) (isOk bool, err error) {
	return
}

func FieldFindOrCreate(filter types.FieldFilter) (data types.Field, err error) {
	return
}

func FieldUpdateOrCreate(filter types.FieldFilter) (data types.Field, err error) {
	return
}

// add all assign functions
