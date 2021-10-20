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

func NewModelWithSoftDeleteFind(filter types.NewModelWithSoftDeleteFilter) (result []types.NewModelWithSoftDelete, totalRecords int, err error) {

	var dbModelData []dbmodels.NewModelWithSoftDelete

	limit := filter.PerPage
	offset := filter.GetOffset()

	if limit < 1 {
		return nil, 0, nil
	}

	filterIds := filter.GetIds()
	filterExceptIds := filter.GetExceptIds()

	var count int64

	criteria := core.Db.Where(dbmodels.NewModelWithSoftDelete{})

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

	q := criteria.Model(dbmodels.NewModelWithSoftDelete{}).Count(&count)

	if q.Error != nil {
		log.Println("FindNewModelWithSoftDelete > Ошибка получения данных:", q.Error)
		return result, 0, nil
	}

	// order global criteria
	if len(filter.Order) > 0 {
		for index, Field := range filter.Order {
			if core.Db.Migrator().HasColumn(&dbmodels.NewModelWithSoftDelete{}, Field) {
				criteria = criteria.Order("\"" + strings.ToLower(Field) + "\"" + " " + filter.OrderDirection[index])
			} else {
				err = errors.NewErrorWithCode("Ordering by unknown Field", errors.ErrorCodeNotValid, Field)
				return
			}
		}
	}

	q = criteria.Limit(limit).Offset(offset).Find(&dbModelData)

	if q.Error != nil {
		log.Println("FindNewModelWithSoftDelete > Ошибка получения данных2:", q.Error)
		return []types.NewModelWithSoftDelete{}, 0, nil
	}

	//формирование результатов
	for _, item := range dbModelData {
		result = append(result, AssignNewModelWithSoftDeleteTypeFromDb(item))
	}

	return result, int(count), nil
}

func NewModelWithSoftDeleteMultiCreate(filter types.NewModelWithSoftDeleteFilter) (data []types.NewModelWithSoftDelete, err error) {

	typeModelList, err := filter.GetNewModelWithSoftDeleteModelList()

	if err != nil {
		return
	}

	tx := core.Db.Begin()

	for _, typeModel := range typeModelList {

		filter.SetNewModelWithSoftDeleteModel(typeModel)
		item, e := NewModelWithSoftDeleteCreate(filter, tx)

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

func NewModelWithSoftDeleteCreate(filter types.NewModelWithSoftDeleteFilter, query *gorm.DB) (data types.NewModelWithSoftDelete, err error) {

	typeModel := filter.GetNewModelWithSoftDeleteModel()
	dbModel := AssignNewModelWithSoftDeleteDbFromType(typeModel)
	dbModel.ID = 0

	dbModel.Validate()

	if !dbModel.IsValid() {
		fmt.Println("NewModelWithSoftDeleteCreate > Create NewModelWithSoftDelete error:", dbModel)
		return types.NewModelWithSoftDelete{}, dbModel.GetValidationError()
	}

	err = query.Create(&dbModel).Error

	if err != nil {
		fmt.Println("NewModelWithSoftDeleteCreate > Create NewModelWithSoftDelete error:", err)
		return types.NewModelWithSoftDelete{}, errors.NewErrorWithCode(fmt.Sprintf("Can't create NewModelWithSoftDelete: %s", err.Error()), errors.ErrorCodeSqlError, "")
	}

	return AssignNewModelWithSoftDeleteTypeFromDb(dbModel), nil
}

func NewModelWithSoftDeleteRead(filter types.NewModelWithSoftDeleteFilter) (data types.NewModelWithSoftDelete, err error) {

	filter.CurrentPage = 1
	filter.PerPage = 1
	filter.ClearIds()
	filter.AddId(filter.GetCurrentId())

	findData, _, err := NewModelWithSoftDeleteFind(filter)

	if len(findData) > 0 {
		return findData[0], nil
	}

	return types.NewModelWithSoftDelete{}, errors.NewErrorWithCode("Not found", errors.ErrorCodeNotFound, "")
}

func NewModelWithSoftDeleteMultiUpdate(filter types.NewModelWithSoftDeleteFilter) (data []types.NewModelWithSoftDelete, err error) {

	typeModelList, err := filter.GetNewModelWithSoftDeleteModelList()
	if err != nil {
		return nil, err
	}

	tx := core.Db.Begin()

	for _, typeModel := range typeModelList {

		filter.SetNewModelWithSoftDeleteModel(typeModel)
		filter.ClearIds()
		filter.SetCurrentId(typeModel.Id)

		item, e := NewModelWithSoftDeleteUpdate(filter, tx)

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

func NewModelWithSoftDeleteUpdate(filter types.NewModelWithSoftDeleteFilter, query *gorm.DB) (data types.NewModelWithSoftDelete, err error) {

	filter.CurrentPage = 1
	filter.PerPage = 1

	existsModel, err := NewModelWithSoftDeleteRead(filter)

	if existsModel.Id < 1 || err != nil {
		err = errors.NewErrorWithCode("NewModelWithSoftDelete not found in db with id: "+strconv.Itoa(filter.GetCurrentId()), errors.ErrorCodeNotFound, "Id")
		return
	}

	newModel := filter.GetNewModelWithSoftDeleteModel()

	updateModel := AssignNewModelWithSoftDeleteDbFromType(newModel)
	updateModel.ID = existsModel.Id

	updateModel.Validate()
	if !updateModel.IsValid() {
		err = updateModel.GetValidationError()
		return
	}

	err = query.Model(dbmodels.NewModelWithSoftDelete{}).Save(&updateModel).Error

	if err != nil {
		return
	}

	return AssignNewModelWithSoftDeleteTypeFromDb(updateModel), nil
}

func NewModelWithSoftDeleteMultiDelete(filter types.NewModelWithSoftDeleteFilter) (isOk bool, err error) {

	typeModelList, err := filter.GetNewModelWithSoftDeleteModelList()

	if err != nil {
		return
	}

	isOk = true

	tx := core.Db.Begin()

	for _, typeModel := range typeModelList {

		filter.SetNewModelWithSoftDeleteModel(typeModel)
		filter.ClearIds()
		filter.SetCurrentId(typeModel.Id)

		_, e := NewModelWithSoftDeleteDelete(filter, tx)

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

func NewModelWithSoftDeleteDelete(filter types.NewModelWithSoftDeleteFilter, query *gorm.DB) (isOk bool, err error) {

	filter.CurrentPage = 1
	filter.PerPage = 1

	existsModel, err := NewModelWithSoftDeleteRead(filter)

	if err != nil {
		err = errors.NewErrorWithCode(fmt.Sprintf("NewModelWithSoftDeleteRead error: %s", err.Error()), errors.ErrorCodeNotFound, "")
		return
	}

	if existsModel.Id < 1 {
		err = errors.NewErrorWithCode("Region not found in db with id: "+strconv.Itoa(filter.GetCurrentId()), errors.ErrorCodeNotFound, "")
		return
	}

	dbModel := AssignNewModelWithSoftDeleteDbFromType(existsModel)
	err = query.Model(dbmodels.NewModelWithSoftDelete{}).
		Where(dbmodels.NewModelWithSoftDelete{ID: dbModel.ID}).
		Delete(&dbModel).Error

	if err != nil {
		return
	}

	return true, nil
}

func NewModelWithSoftDeleteFindOrCreate(filter types.NewModelWithSoftDeleteFilter) (data types.NewModelWithSoftDelete, err error) {

	filter.CurrentPage = 1
	filter.PerPage = 1

	findOrCreateModel := AssignNewModelWithSoftDeleteDbFromType(filter.GetNewModelWithSoftDeleteModel())

	findOrCreateModel.Validate()
	if !findOrCreateModel.IsValid() {
		err = findOrCreateModel.GetValidationError()
		return
	}

	err = core.Db.Model(dbmodels.NewModelWithSoftDelete{}).
		Where(dbmodels.NewModelWithSoftDelete{ID: findOrCreateModel.ID}).
		FirstOrCreate(&findOrCreateModel).Error

	if err != nil {
		return
	}

	return AssignNewModelWithSoftDeleteTypeFromDb(findOrCreateModel), nil
}

func NewModelWithSoftDeleteUpdateOrCreate(filter types.NewModelWithSoftDeleteFilter) (data types.NewModelWithSoftDelete, err error) {

	filter.CurrentPage = 1
	filter.PerPage = 1

	updateOrCreateModel := AssignNewModelWithSoftDeleteDbFromType(filter.GetNewModelWithSoftDeleteModel())

	updateOrCreateModel.Validate()

	if !updateOrCreateModel.IsValid() {
		err = updateOrCreateModel.GetValidationError()
		return
	}

	err = core.Db.Model(dbmodels.NewModelWithSoftDelete{}).
		Where(dbmodels.NewModelWithSoftDelete{ID: updateOrCreateModel.ID}).
		Assign(updateOrCreateModel).
		FirstOrCreate(&updateOrCreateModel).Error

	if err != nil {
		return
	}

	return AssignNewModelWithSoftDeleteTypeFromDb(updateOrCreateModel), nil
}

func AssignNewModelWithSoftDeleteTypeFromDb(dbNewModelWithSoftDelete dbmodels.NewModelWithSoftDelete) types.NewModelWithSoftDelete {

	return types.NewModelWithSoftDelete{
		Id: dbNewModelWithSoftDelete.ID,
	}
}

func AssignNewModelWithSoftDeleteDbFromType(typesNewModelWithSoftDelete types.NewModelWithSoftDelete) dbmodels.NewModelWithSoftDelete {

	return dbmodels.NewModelWithSoftDelete{
		ID: typesNewModelWithSoftDelete.Id,
	}
}

func AssignNewModelWithSoftDeleteTypeFromDbList(dbNewModelWithSoftDeletes []dbmodels.NewModelWithSoftDelete) (res []types.NewModelWithSoftDelete) {
	for _, model := range dbNewModelWithSoftDeletes {
		res = append(res, AssignNewModelWithSoftDeleteTypeFromDb(model))
	}
	return res
}

func AssignNewModelWithSoftDeleteDbFromTypeList(typesNewModelWithSoftDeletes []types.NewModelWithSoftDelete) (res []dbmodels.NewModelWithSoftDelete) {
	for _, model := range typesNewModelWithSoftDeletes {
		res = append(res, AssignNewModelWithSoftDeleteDbFromType(model))
	}
	return res
}
