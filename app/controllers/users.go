package controllers

import (
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/qor/render"
	"github.com/raven-chen/qor_doc_demo/app/models"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	// var ctx *admin.Context
	DB, _ := gorm.Open("sqlite3", "demo.db")

	var user models.User
	DB.First(&user, 1)

	funcMap := template.FuncMap{
		"Greet": func(name string) string { return "hello " + name },
	}

	ctx := make(map[string]interface{})
	// ctx["CurrentUser"] = user
	// ctx.Result = user

	Render := render.New()
	Render.Layout("new").Funcs(funcMap).Execute("users/profile", ctx, r, w)
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Greeting(name string) string {
	return "hello " + name
}
