package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Bem bonita", 29., 10},
		{"Notebook", "Bem rápido", 1999., 2},
		{"Cafeteira", "Das boas", 99., 3},
		{"Phone", "Ótimo som", 75, 7},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
