package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/raven-chen/qor_doc_demo/app/controllers"
	"github.com/raven-chen/qor_doc_demo/config/admin"
	"github.com/raven-chen/qor_doc_demo/db"
	"github.com/theplant/appkit/contexts"
	"github.com/theplant/appkit/log"
	"github.com/theplant/appkit/server"
	"gopkg.in/authboss.v0"
)

func Mux(logger log.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/profile", controllers.Profile)

	// I18n := i18n.New(
	//  database.New(DB),                                       // load translations from the database
	//  yaml.New(filepath.Join(config.Root, "config/locales")), // load translations from the YAML files in directory `config/locales`
	// )
	// fmt.Println(config.Root)
	// // I18n := i18n.New(database.New(DB), yaml.New(filepath.Join(config.Root, "config/locales"),)
	// Admin.AddResource(I18n)
	// fmt.Println(I18n.T("en-US", "demo.greeting"))
	// fmt.Println(I18n.T("en-US", "demo.hello", "Marry!"))

	// I18n.SaveTranslation(&i18n.Translation{Key: "hello-world", Locale: "en-US", Value: "hello world"})

	// var a interface{}
	// a["Name"] = "test"
	// I18n.AddTranslation(&i18n.Translation{Key: "hello", Locale: "en-US", Value: "Hello {{.Name}}"})
	// fmt.Println(I18n.T("en-US", "hello", {Name: "Jinzhu"})) //=> Hello Jinzhu

	// amount to /admin, so visit `/admin` to view the admin interface
	admin.Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")

	ab := authboss.New() // Usually store this globally
	ab.MountPath = "/authboss"
	ab.LogWriter = os.Stdout

	// Make sure to put authboss's router somewhere
	http.Handle("/authboss", ab.NewRouter())

	middleware := server.Compose(
		contexts.WithGorm(db.DB),
		server.DefaultMiddleware(logger),
	)

	return middleware(mux)
}
