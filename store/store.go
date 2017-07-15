package store

import (
	"database/sql"
	"fmt"
	"log"
)

//SayHello - Just for Checking
func SayHello() string {
	return "Hello sir!!"
}

//Sum - Add t
func Sum(a, b int) int {
	return a + b
}

//Product - Holds Product Attributes
type Product struct {
	ID    int
	Name  string
	Price float64
}

//Connection - store db Connection
type Connection struct {
	DB *sql.DB
}

//Persist - Store Product attributes in Database
func Persist(product *Product, connection *Connection) bool {
	success := true
	var err error
	_, err = connection.DB.Exec("INSERT INTO products(name,price) VALUES($1,$2) RETURNING id", product.Name, product.Price)
	if err != nil {
		fmt.Println("Insert data failed")
		success = false
	}
	return success
}

//GetConnection - Get a Database Connection
func (connection *Connection) GetConnection() *sql.DB {
	return connection.DB
}

//Initialize - Create Database Connection
func (connection *Connection) Initialize(username, password, dbname string) (*sql.DB, error) {
	var err error
	dataSourceName := fmt.Sprintf("user=%s password=%s database=%s", username, password, dbname)
	fmt.Println(dataSourceName)
	connection.DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(connection.DB)
	return connection.DB, err
}
