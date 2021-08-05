package fixtures

import (
	"skeleton-app/core"
	"skeleton-app/dbmodels"
	"skeleton-app/logic"
	"skeleton-app/settings"
	"skeleton-app/types"
)

func AddRouteType() {

	resourceType := logic.AssignResourceTypeDbFromType(types.ResourceType{
		Id:   settings.HttpRouteResourceType.Int(),
		Name: "Route",
	})
	core.Db.Model(dbmodels.ResourceType{}).FirstOrCreate(&resourceType)
}
