package logic

import (
	"gosha_v2/errors"
	"gosha_v2/services/filesystem"
	"gosha_v2/types"
)

func ApplicationFind(_ types.ApplicationFilter) (result []types.Application, totalRecords int, err error) {

	return
}

func ApplicationMultiCreate(_ types.ApplicationFilter) (data []types.Application, err error) {

	return
}

func ApplicationCreate(filter types.ApplicationFilter) (data types.Application, err error) {
	existApps, _, err := ApplicationFind(filter)
	if err != nil {
		return types.Application{}, err
	}
	if len(existApps) > 0 {
		err = errors.NewErrorWithCode("Application already exist", errors.ErrorCodeAlreadyExist, "")
		return types.Application{}, err
	}
	model := filter.GetApplicationModel()
	if !model.IsValid() {
		err = model.GetValidationError()
		return types.Application{}, err
	}

	currentPath, err := filter.GetPwd()
	if err != nil {
		err = errors.NewErrorWithCode(err.Error(), errors.ErrorCodeNotFound, "")
		return types.Application{}, err
	}

	currentPath = "/home/alex/projects/gosha_test/"

	// Copy all skeleton to new dir
	err = filesystem.CopySkeletonApp(currentPath)
	if err != nil {
		return types.Application{}, err
	}
	// Update package name and imports
	err = filesystem.FindAndReplaceImports(currentPath)
	if err != nil {
		return types.Application{}, err
	}
	err = filesystem.CreateGoMod(currentPath)
	if err != nil {
		return types.Application{}, err
	}

	// Update db connection and script

	// Update seed for password
	// Update user fixtures

	// Set PK type



	return
}

func ApplicationRead(_ types.ApplicationFilter) (data types.Application, err error) {
	return
}

func ApplicationMultiUpdate(_ types.ApplicationFilter) (data []types.Application, err error) {
	return
}

func ApplicationUpdate(_ types.ApplicationFilter) (data types.Application, err error) {
	return
}

func ApplicationMultiDelete(_ types.ApplicationFilter) (isOk bool, err error) {
	return
}

func ApplicationDelete(_ types.ApplicationFilter) (isOk bool, err error) {
	return
}

func ApplicationFindOrCreate(_ types.ApplicationFilter) (data types.Application, err error) {
	return
}

func ApplicationUpdateOrCreate(_ types.ApplicationFilter) (data types.Application, err error) {
	return
}

// add all assign functions
