package main

import (
	"fmt"
	"log"

	"example.com/eshan/db"
	"example.com/eshan/repo"
)

func main() {
	db := &db.Database{Host: "localhost", Port: 5432}
	repo := repo.NewRepository(db)
	val, err := repo.Get("123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}
