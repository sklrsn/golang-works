package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/naoina/toml"
	"sklrsn.github.com/Robot/store"
)

var (
	path = "/usr/local/go/src/sklrsn.github.com/Robot/"
)

//ApplicationConfig - Store Postgres Connection
type ApplicationConfig struct {
	Title    string `toml:"title"`
	Database struct {
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Database string `toml:"database"`
	} `toml:"database"`
}

func main() {
	file, err := os.Open(path + "server-config.toml")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	defer file.Close()

	var config ApplicationConfig
	if err := toml.NewDecoder(file).Decode(&config); err != nil {
		panic(err)
	}

	connection := store.Connection{}
	connection.Initialize(string(config.Database.Host), string(config.Database.Port), string(config.Database.Username), string(config.Database.Password), string(config.Database.Database))

	product := store.Product{Name: "Bat", Price: 270}
	fmt.Println(store.Persist(&product, &connection))

}
