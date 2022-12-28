module github.com/AndreVeiga/go-web-first

go 1.19

replace db => ./db

replace models => ./models

replace rotas => ./rotas

replace controllers => ./controllers

require (
	controllers v0.0.0-00010101000000-000000000000 // indirect
	db v0.0.0-00010101000000-000000000000 // indirect
	github.com/lib/pq v1.10.7 // indirect
	models v0.0.0-00010101000000-000000000000 // indirect
	rotas v0.0.0-00010101000000-000000000000 // indirect
)
