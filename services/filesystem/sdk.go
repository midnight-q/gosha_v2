package filesystem

import (
	"gosha_v2/errors"
	"gosha_v2/types"
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

	_ = os.RemoveAll(folder)
	_ = os.MkdirAll(folder, 0755)
	_ = os.MkdirAll(folderStore, 0755)
	_ = os.MkdirAll(folderData, 0755)

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
