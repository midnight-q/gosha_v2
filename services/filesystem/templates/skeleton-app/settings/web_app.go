package settings

const ServerPort = "7005"

const (
	FunctionTypeFind              = "00000000-0000-0000-0000-000000000001"
	FunctionTypeRead              = "00000000-0000-0000-0000-000000000002"
	FunctionTypeUpdate            = "00000000-0000-0000-0000-000000000003"
	FunctionTypeMultiUpdate       = "00000000-0000-0000-0000-000000000023"
	FunctionTypeDelete            = "00000000-0000-0000-0000-000000000004"
	FunctionTypeMultiDelete       = "00000000-0000-0000-0000-000000000024"
	FunctionTypeCreate            = "00000000-0000-0000-0000-000000000005"
	FunctionTypeMultiCreate       = "00000000-0000-0000-0000-000000000025"
	FunctionTypeFindOrCreate      = "00000000-0000-0000-0000-000000000006"
	FunctionTypeMultiFindOrCreate = "00000000-0000-0000-0000-000000000026"
	FunctionTypeUpdateOrCreate    = "00000000-0000-0000-0000-000000000007"
)

const (
	OrderDirectionDesc = "desc"
	OrderDirectionAsc  = "asc"
)
