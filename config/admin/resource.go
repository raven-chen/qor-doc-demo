package admin

import (
	"github.com/qor/admin"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	"github.com/raven-chen/qor_doc_demo/app/models"
	"github.com/raven-chen/qor_doc_demo/db"
)

var (
	Admin = admin.New(&qor.Config{DB: db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff)})
)

func init() {
	Admin.SetSiteName("Qor DEMO")

	Admin.AddResource(&models.Product{})

	// Create resources from GORM-backend model
	user := Admin.AddResource(&models.User{})
	user.UseTheme("users")

	user.NewAttrs(
		&admin.Section{
			Title: "Account",
			Rows: [][]string{
				{"Name", "Gender"},
			},
		},
		&admin.Section{
			Title: "Detail",
			Rows: [][]string{
				{"Active"},
			},
		},
	)

	user.Action(&admin.Action{
		Name: "enable",
		Handle: func(actionArgument *admin.ActionArgument) error {
			// `FindSelectedRecords` => return selected record in bulk action mode, return current record in other mode
			for _, record := range actionArgument.FindSelectedRecords() {
				actionArgument.Context.DB.Model(record.(*models.User)).Update("Active", true)
			}
			return nil
		},
		Modes: []string{"index", "edit", "show", "menu_item"},
	})

	// order := Admin.AddResource(&Order{})

	// order.Action(&admin.Action{
	//  Name: "Ship",
	//  Handle: func(argument *admin.ActionArgument) error {
	//    trackingNumberArgument := argument.Argument.(*trackingNumberArgument)
	//    for _, record := range argument.FindSelectedRecords() {
	//      argument.Context.GetDB().Model(record).UpdateColumn("tracking_number", trackingNumberArgument.TrackingNumber)
	//    }
	//    return nil
	//  },
	//  Resource: Admin.NewResource(&trackingNumberArgument{}),
	//  Modes:    []string{"show", "menu_item"},
	// })
}
