package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"sklrsn.github.com/Robot/store"
)

var (
	HOST     = "172.17.0.1"
	PORT     = "5432"
	USERNAME = "postgres"
	PASSWORD = "love"
	DATABASE = "postgres"
)

type Config struct {
	Title    string `toml:"title"`
	Database struct {
		Username string `toml:"username"`
		Password string `toml:"password"`
		Database string `toml:"database"`
	} `toml:"database"`
}

func main() {
	// file, err := os.Open("server-config.toml")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	//
	// var config Config
	// if err := toml.NewDecoder(file).Decode(&config); err != nil {
	// 	panic(err)
	// }

	connection := store.Connection{}
	// connection.Initialize(string(config.Database.Username), string(config.Database.Password), string(config.Database.Database))
	connection.Initialize(HOST, PORT, USERNAME, PASSWORD, DATABASE)

	product := store.Product{Name: "Bat", Price: 27}
	fmt.Println(store.Persist(&product, &connection))

}
