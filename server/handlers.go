package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jesperkha/w4d/database"
)

func indexHandler(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseGlob("app/index.html")
		if err != nil {
			log.Println(err)
		}

		tmpl, err = tmpl.ParseGlob("app/templates/*.html")
		if err != nil {
			log.Println(err)
		}

		recipes := db.GetAllRecipes()
		if err := tmpl.ExecuteTemplate(w, "main", recipes); err != nil {
			log.Println(err)
		}
	}
}
