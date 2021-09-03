package fixtures

import (
	"skeleton-app/core"
	"skeleton-app/dbmodels"
	"skeleton-app/settings"
)

func AddRouteType() {

	resourceTypes := []dbmodels.ResourceType{
		{
			ID:   settings.HttpRouteResourceType.Int(),
			Name: "Route",
		},
		{
			ID:   settings.WsResourceType.Int(),
			Name: "Ws handler",
		},
		{
			ID:   settings.HtmlResourceType.Int(),
			Name: "HTML template resource",
		},
	}
	for _, resourceType := range resourceTypes {
		core.Db.Model(dbmodels.ResourceType{}).FirstOrCreate(&resourceType)
	}
}
