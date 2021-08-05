package types

import (
	"net/http"
	"skeleton-app/settings"
)

type Region struct {
	Id   int
	Name int
	Code string

	validator
}

func (region *Region) Validate() {
}

type RegionFilter struct {
	model Region
	list  []Region

	AbstractFilter
}

func GetRegionFilter(request *http.Request, functionType string) (filter RegionFilter, err error) {

	filter.request = request
	filter.rawRequestBody, err = GetRawBodyContent(request)
	if err != nil {
		return filter, err
	}

	switch functionType {
	case settings.FunctionTypeMultiCreate, settings.FunctionTypeMultiUpdate, settings.FunctionTypeMultiDelete, settings.FunctionTypeMultiFindOrCreate:
		err = ReadJSON(filter.rawRequestBody, &filter.list)
		if err != nil {
			return
		}
		break
	default:
		err = ReadJSON(filter.rawRequestBody, &filter.model)
		if err != nil {
			return
		}
		break
	}

	filter.AbstractFilter, err = GetAbstractFilter(request, filter.rawRequestBody, functionType)

	return filter, err
}

func (filter *RegionFilter) GetRegionModel() Region {

	filter.model.Validate()

	return filter.model
}

func (filter *RegionFilter) GetRegionModelList() (data []Region, err error) {

	for k, _ := range filter.list {
		filter.list[k].Validate()

		if !filter.list[k].IsValid() {
			err = filter.list[k].GetValidationError()
			break
		}
	}

	return filter.list, nil
}

func (filter *RegionFilter) SetRegionModel(typeModel Region) {

	filter.model = typeModel
}

func (filter *RegionFilter) SetRegionModelList(data []Region) {

	filter.list = data
}
