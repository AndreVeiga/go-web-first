package controllers

import (
	"html/template"
	"models"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()

	temp.ExecuteTemplate(w, "Index", produtos)
}
