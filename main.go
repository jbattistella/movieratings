package main

import db "github.com/jbattistella/movieratings/db/sqlc"

func main() {
	db.Connect()
}
