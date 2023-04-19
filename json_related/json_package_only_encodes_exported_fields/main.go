package main

import (
	"encoding/json"
	"fmt"
)

func does_not_encode_non_exported_field() {
	type user struct {
		name     string
		password string
	}

	users := []user{
		{
			name:     "eshan",
			password: "eshan",
		},
		{
			name:     "eshan1",
			password: "eshan1",
		},
	}

	out, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}

func does_encode_exported_field() {
	type user struct {
		Name     string
		Password string
	}

	users := []user{
		{
			Name:     "eshan",
			Password: "eshan",
		},
		{
			Name:     "eshan1",
			Password: "eshan1",
		},
	}

	out, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}

func main() {
	does_not_encode_non_exported_field()
	does_encode_exported_field()
}
