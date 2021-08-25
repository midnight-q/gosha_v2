package logic

import (
	"gosha_v2/types"
)

func FieldTypeFind(_ types.FieldTypeFilter) (result []types.FieldType, totalRecords int, err error) {
	result = []types.FieldType{
		{
			Name:                 "Int",
			Code:                 "int",
			IsAvailableForFilter: true,
		},
		{
			Name:                 "Float",
			Code:                 "float64",
			IsAvailableForFilter: true,
		},
		{
			Name:                 "String",
			Code:                 "string",
			IsAvailableForFilter: true,
		},
		{
			Name:                 "Byte",
			Code:                 "byte",
			IsAvailableForFilter: false,
		},
		{
			Name:                 "Int",
			Code:                 "uuid",
			IsAvailableForFilter: true,
		},
		{
			Name:                 "Time",
			Code:                 "time",
			IsAvailableForFilter: true,
		},
		{
			Name:                 "Bool",
			Code:                 "bool",
			IsAvailableForFilter: true,
		},
	}

	return result, len(result), nil
}
