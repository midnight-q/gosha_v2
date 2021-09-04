package logic

import (
	"gosha_v2/services/filesystem"
	"gosha_v2/types"
)

func SdkFind(filter types.SdkFilter) (result []types.Sdk, totalRecords int, err error) {

	return
}

func SdkCreate(filter types.SdkFilter) (data types.Sdk, err error) {
	currentPath, err := filter.GetPwd()
	if err != nil {
		return types.Sdk{}, err
	}

	sdkModel := filter.GetSdkModel()

	fApplication := types.ApplicationFilter{}
	fApplication.SetPwd(currentPath)
	fApplication.CurrentPage = 1
	fApplication.PerPage = 1000
	_, _, err = ApplicationFind(fApplication)
	if err != nil {
		return types.Sdk{}, err
	}
	fModel := types.ModelFilter{}
	fModel.SetPwd(currentPath)
	fModel.CurrentPage = 1
	fModel.PerPage = 1000
	models, _, err := ModelFind(fModel)
	if err != nil {
		return types.Sdk{}, err
	}

	if sdkModel.JsTypes {
		err = filesystem.GenerateJsTypesSdk(currentPath, sdkModel.Namespace, models)
		if err != nil {
			return types.Sdk{}, err
		}
	}
	if sdkModel.Js {
		err = filesystem.GenerateJsSdk(currentPath, sdkModel.Namespace, models)
		if err != nil {
			return types.Sdk{}, err
		}
	}
	if sdkModel.Ts {
		err = filesystem.GenerateTsSdk(currentPath, sdkModel.Namespace, models)
		if err != nil {
			return types.Sdk{}, err
		}
	}
	if sdkModel.Go {
		err = filesystem.GenerateGoSdk(currentPath, sdkModel.Namespace, models)
		if err != nil {
			return types.Sdk{}, err
		}
	}
	if sdkModel.Dart {
		err = filesystem.GenerateDartSdk(currentPath, sdkModel.Namespace, models)
		if err != nil {
			return types.Sdk{}, err
		}
	}

	return
}
