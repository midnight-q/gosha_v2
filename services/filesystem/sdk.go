package filesystem

import (
	"fmt"
	"gosha_v2/errors"
	"gosha_v2/services/filesystem/sdk"
	"gosha_v2/types"
	"io/ioutil"
	"os"
)

func GenerateJsTypesSdk(currentPath, namespace string, models []types.Model) (err error) {

	folder := currentPath + "/jstypes"
	folderStore := folder + "/store"
	folderData := folder + "/data"
	fileModel := folder + "/apiModel.js"
	apiFile := folder + "/api.js"
	apiCSRFile := folder + "/apiCSR.js"
	apiSSRFile := folder + "/apiSSR.js"
	commonFile := folder + "/common.js"

	_, err = os.Stat(folder)
	if err != nil {
		if os.IsNotExist(err) {
			_ = os.RemoveAll(folder)
			err = nil
		} else {
			return err
		}
	}

	err = os.MkdirAll(folder, 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(folderStore, 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(folderData, 0755)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileModel, sdk.GetModelsFile(models), 0644)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(apiFile, sdk.GetSkeletonTemplateFile(templateFS, "/jstypes/api.js"), 0644)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(apiCSRFile, sdk.GetSkeletonTemplateFile(templateFS, "/jstypes/apiCSR.js"), 0644)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(apiSSRFile, sdk.GetSkeletonTemplateFile(templateFS, "/jstypes/apiSSR.js"), 0644)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(commonFile, sdk.GetSkeletonTemplateFile(templateFS, "/jstypes/common.js"), 0644)
	if err != nil {
		return err
	}

	for _, d := range sdk.GetStores(templateFS, models) {
		err = ioutil.WriteFile(fmt.Sprintf("%s/%s.js", folderStore, d.Name), []byte(d.Content), 0644)
		if err != nil {
			return err
		}
	}

	for _, d := range sdk.GetData(templateFS, models) {
		err = ioutil.WriteFile(fmt.Sprintf("%s/%sData.js", folderData, d.Name), []byte(d.Content), 0644)
		if err != nil {
			return err
		}
	}

	return
}

func GenerateJsSdk(currentPath, namespace string, models []types.Model) (err error) {
	return errors.New("GenerateJsSdk not implement")
}

func GenerateTsSdk(currentPath, namespace string, models []types.Model) (err error) {
	return errors.New("GenerateTsSdk not implement")
}

func GenerateGoSdk(currentPath, namespace string, models []types.Model) (err error) {
	return errors.New("GenerateGoSdk not implement")
}

func GenerateDartSdk(currentPath, namespace string, models []types.Model) (err error) {
	return errors.New("GenerateDartSdk not implement")
}
