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
	if !isFilter {
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
			//create assigner
		}
	} else {
		if fieldModel.IsDbField {
			err = errors.New("Cant create dbField in filter")
			return types.Field{}, err
		}

		err = filesystem.AddFieldInModel(fieldModel.Name, fieldModel.CommentType, fieldModel.ModelName, model.TypePath, fieldModel.Type, fieldType)
		if err != nil {
			return types.Field{}, err
		}

		// Create parser in GetFilter depends on type
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
