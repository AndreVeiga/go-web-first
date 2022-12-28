package main

import (
	"net/http"
	"rotas"
)

func main() {
	rotas.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
