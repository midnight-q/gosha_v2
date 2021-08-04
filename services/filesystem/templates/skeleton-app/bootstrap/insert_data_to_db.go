package bootstrap

import (
	"fmt"
	"os"
	"skeleton-app/bootstrap/fixtures"
	"skeleton-app/core"
	"skeleton-app/dbmodels"
)

func FillDBTestData() {

	if core.DbErr != nil {
		fmt.Println("Error dabatabse connect", core.DbErr.Error())
		os.Exit(0)
	}

	isDropTables := false

	if (len(os.Args) > 1 && os.Args[1] == "drop") ||
		(len(os.Args) > 2 && os.Args[2] == "drop") {
		isDropTables = true
	}

	if isDropTables == true {

		core.Db.Migrator().DropTable(

			&dbmodels.Language{},
			&dbmodels.Region{},
			&dbmodels.TranslateError{},
			&dbmodels.CurrentUser{},
			&dbmodels.Auth{},
			&dbmodels.UserRole{},
			&dbmodels.ResourceType{},
			&dbmodels.Resource{},
			&dbmodels.RoleResource{},
			&dbmodels.Role{},
			&dbmodels.User{},
		)

		fmt.Println("All tables removed")
		os.Exit(1)
	}

	_ = core.Db.Migrator().AutoMigrate(

		&dbmodels.Language{},
		&dbmodels.Region{},
		&dbmodels.TranslateError{},
		&dbmodels.CurrentUser{},
		&dbmodels.Auth{},
		&dbmodels.UserRole{},
		&dbmodels.ResourceType{},
		&dbmodels.Resource{},
		&dbmodels.RoleResource{},
		&dbmodels.Role{},
		&dbmodels.User{},
	)

	// add fixtures
	fixtures.AddRouteType()
	fixtures.AddUser()
}
