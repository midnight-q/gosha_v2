package logic

import (
	"skeleton-app/types"
	"strconv"

	"fmt"
	"log"
	"skeleton-app/core"
	"skeleton-app/dbmodels"
	"skeleton-app/errors"
	"strings"

	"gorm.io/gorm"
)

func NewModelFind(filter types.NewModelFilter) (result []types.NewModel, totalRecords int, err error) {

	var dbModelData []dbmodels.NewModel

	limit := filter.PerPage
	offset := filter.GetOffset()

	if limit < 1 {
		return nil, 0, nil
	}

	filterIds := filter.GetIds()
	filterExceptIds := filter.GetExceptIds()

	var count int64

	criteria := core.Db.Where(dbmodels.NewModel{})

	if len(filterIds) > 0 {
		criteria = criteria.Where("id in (?)", filterIds)
	}

	if len(filterExceptIds) > 0 {
		criteria = criteria.Where("id not in (?)", filterExceptIds)
	}

	//if len(filter.Search) > 0 {
	//
	//    s := ("%" + filter.Search + "%")
	//
	//    if len(filter.SearchBy) > 0 {
	//    subCriteria := core.Db
	//
	//        for _, field := range filter.SearchBy {
	//
	//            if core.Db.Migrator().HasColumn(&dbmodels.BaseComposition{}, Field) {
	//                subCriteria = subCriteria.Or("\""+field+"\""+" ilike ?", s)
	//            } else {
	//                err = errors.New("Search by unknown field " + field)
	//                return
	//            }
	//        }
	//    criteria = criteria.Where(subCriteria)
	//    } else {
	//      criteria = criteria.Where("name ilike ? or code ilike ?", ("%" + filter.Search + "%"), ("%" + filter.Search + "%"))
	//    }
	//}

	q := criteria.Model(dbmodels.NewModel{}).Count(&count)

	if q.Error != nil {
		log.Println("FindNewModel > Ошибка получения данных:", q.Error)
		return result, 0, nil
	}

	// order global criteria
	if len(filter.Order) > 0 {
		for index, Field := range filter.Order {
			if core.Db.Migrator().HasColumn(&dbmodels.NewModel{}, Field) {
				criteria = criteria.Order("\"" + strings.ToLower(Field) + "\"" + " " + filter.OrderDirection[index])
			} else {
				err = errors.NewErrorWithCode("Ordering by unknown Field", errors.ErrorCodeNotValid, Field)
				return
			}
		}
	}

	q = criteria.Limit(limit).Offset(offset).Find(&dbModelData)

	if q.Error != nil {
		log.Println("FindNewModel > Ошибка получения данных2:", q.Error)
		return []types.NewModel{}, 0, nil
	}

	//формирование результатов
	for _, item := range dbModelData {
		result = append(result, AssignNewModelTypeFromDb(item))
	}

	return result, int(count), nil
}

func NewModelMultiCreate(filter types.NewModelFilter) (data []types.NewModel, err error) {

	typeModelList, err := filter.GetNewModelModelList()

	if err != nil {
		return
	}

	tx := core.Db.Begin()

	for _, typeModel := range typeModelList {

		filter.SetNewModelModel(typeModel)
		item, e := NewModelCreate(filter, tx)

		if e != nil {
			err = e
			data = nil
			break
		}

		data = append(data, item)
	}

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return
}

func NewModelCreate(filter types.NewModelFilter, query *gorm.DB) (data types.NewModel, err error) {

	typeModel := filter.GetNewModelModel()
	dbModel := AssignNewModelDbFromType(typeModel)
	dbModel.ID = 0

	dbModel.Validate()

	if !dbModel.IsValid() {
		fmt.Println("NewModelCreate > Create NewModel error:", dbModel)
		return types.NewModel{}, dbModel.GetValidationError()
	}

	err = query.Create(&dbModel).Error

	if err != nil {
		fmt.Println("NewModelCreate > Create NewModel error:", err)
		return types.NewModel{}, errors.NewErrorWithCode(fmt.Sprintf("Can't create NewModel: %s", err.Error()), errors.ErrorCodeSqlError, "")
	}

	return AssignNewModelTypeFromDb(dbModel), nil
}

func NewModelRead(filter types.NewModelFilter) (data types.NewModel, err error) {

	filter.CurrentPage = 1
	filter.PerPage = 1
	filter.ClearIds()
	filter.AddId(filter.GetCurrentId())

	findData, _, err := NewModelFind(filter)

	if len(findData) > 0 {
		return findData[0], nil
	}

	return types.NewModel{}, errors.NewErrorWithCode("Not found", errors.ErrorCodeNotFound, "")
}

func NewModelMultiUpdate(filter types.NewModelFilter) (data []types.NewModel, err error) {

	typeModelList, err := filter.GetNewModelModelList()
	if err != nil {
		return nil, err
	}

	tx := core.Db.Begin()

	for _, typeModel := range typeModelList {

		filter.SetNewModelModel(typeModel)
		filter.ClearIds()
		filter.SetCurrentId(typeModel.Id)

		item, e := NewModelUpdate(filter, tx)

		if e != nil {
			err = e
			data = nil
			break
		}

		data = append(data, item)
	}

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return data, nil
}

func NewModelUpdate(filter types.NewModelFilter, query *gorm.DB) (data types.NewModel, err error) {

	filter.CurrentPage = 1
	filter.PerPage = 1

	existsModel, err := NewModelRead(filter)

	if existsModel.Id < 1 || err != nil {
		err = errors.NewErrorWithCode("NewModel not found in db with id: "+strconv.Itoa(filter.GetCurrentId()), errors.ErrorCodeNotFound, "Id")
		return
	}

	newModel := filter.GetNewModelModel()

	updateModel := AssignNewModelDbFromType(newModel)
	updateModel.ID = existsModel.Id

	updateModel.Validate()
	if !updateModel.IsValid() {
		err = updateModel.GetValidationError()
		return
	}

	err = query.Model(dbmodels.NewModel{}).Save(&updateModel).Error

	if err != nil {
		return
	}

	return AssignNewModelTypeFromDb(updateModel), nil
}

func NewModelMultiDelete(filter types.NewModelFilter) (isOk bool, err error) {

	typeModelList, err := filter.GetNewModelModelList()

	if err != nil {
		return
	}

	isOk = true

	tx := core.Db.Begin()

	for _, typeModel := range typeModelList {

		filter.SetNewModelModel(typeModel)
		filter.ClearIds()
		filter.SetCurrentId(typeModel.Id)

		_, e := NewModelDelete(filter, tx)

		if e != nil {
			err = e
			isOk = false
			break
		}
	}

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return isOk, err
}

func NewModelDelete(filter types.NewModelFilter, query *gorm.DB) (isOk bool, err error) {

	filter.CurrentPage = 1
	filter.PerPage = 1

	existsModel, err := NewModelRead(filter)

	if err != nil {
		err = errors.NewErrorWithCode(fmt.Sprintf("NewModelRead error: %s", err.Error()), errors.ErrorCodeNotFound, "")
		return
	}

	if existsModel.Id < 1 {
		err = errors.NewErrorWithCode("Region not found in db with id: "+strconv.Itoa(filter.GetCurrentId()), errors.ErrorCodeNotFound, "")
		return
	}

	dbModel := AssignNewModelDbFromType(existsModel)
	err = query.Model(dbmodels.NewModel{}).
		Where(dbmodels.NewModel{ID: dbModel.ID}).
		Delete(&dbModel).Error

	if err != nil {
		return
	}

	return true, nil
}

func NewModelFindOrCreate(filter types.NewModelFilter) (data types.NewModel, err error) {

	filter.CurrentPage = 1
	filter.PerPage = 1

	findOrCreateModel := AssignNewModelDbFromType(filter.GetNewModelModel())

	findOrCreateModel.Validate()
	if !findOrCreateModel.IsValid() {
		err = findOrCreateModel.GetValidationError()
		return
	}

	err = core.Db.Model(dbmodels.NewModel{}).
		Where(dbmodels.NewModel{ID: findOrCreateModel.ID}).
		FirstOrCreate(&findOrCreateModel).Error

	if err != nil {
		return
	}

	return AssignNewModelTypeFromDb(findOrCreateModel), nil
}

func NewModelUpdateOrCreate(filter types.NewModelFilter) (data types.NewModel, err error) {

	filter.CurrentPage = 1
	filter.PerPage = 1

	updateOrCreateModel := AssignNewModelDbFromType(filter.GetNewModelModel())

	updateOrCreateModel.Validate()

	if !updateOrCreateModel.IsValid() {
		err = updateOrCreateModel.GetValidationError()
		return
	}

	err = core.Db.Model(dbmodels.NewModel{}).
		Where(dbmodels.NewModel{ID: updateOrCreateModel.ID}).
		Assign(updateOrCreateModel).
		FirstOrCreate(&updateOrCreateModel).Error

	if err != nil {
		return
	}

	return AssignNewModelTypeFromDb(updateOrCreateModel), nil
}

func AssignNewModelTypeFromDb(dbNewModel dbmodels.NewModel) types.NewModel {

	return types.NewModel{
		Id: dbNewModel.ID,
	}
}

func AssignNewModelDbFromType(typesNewModel types.NewModel) dbmodels.NewModel {

	return dbmodels.NewModel{
		ID: typesNewModel.Id,
	}
}

func AssignNewModelSliceTypeFromDb(dbNewModels []dbmodels.NewModel) (res []types.NewModel) {
	for _, model := range dbNewModels {
		res = append(res, AssignNewModelTypeFromDb(model))
	}
	return res
}

func AssignNewModelSliceDbFromType(typesNewModels []types.NewModel) (res []dbmodels.NewModel) {
	for _, model := range typesNewModels {
		res = append(res, AssignNewModelDbFromType(model))
	}
	return res
}
