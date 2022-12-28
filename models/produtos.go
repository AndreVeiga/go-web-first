package models

import (
	"db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaBanco()
	dados, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}

	for dados.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := dados.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}

	defer db.Close()
	return produtos
}
