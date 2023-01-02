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

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()

	query := "insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)"
	prepareDB, err := db.Prepare(query)

	if err != nil {
		panic(err.Error())
	}

	prepareDB.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id int) {
	db := db.ConectaBanco()

	query := "DELETE FROM produtos where id = $1"
	prepareDB, err := db.Prepare(query)

	if err != nil {
		panic(err.Error())
	}

	prepareDB.Exec(id)

	defer db.Close()
}
