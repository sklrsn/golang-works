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

//Person - Holds Person Attributes
type Person struct {
	Name string
	Age  int
}

//Connection - store db Connection
type Connection struct {
	DB *sql.DB
}

//Persist - Store Person attributes in Database
func Persist(person *Person, connection *Connection) bool {
	log.Fatal(*person)
	log.Fatal(connection)
	success := true
	var err error
	_, err = connection.DB.Exec("INSERT TO person(name,age) VALUES($1,$2) RETURNING id", person.Name, person.Age)
	if err != nil {
		log.Fatal("Insert data failed")
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
	connection.DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return connection.DB, err
}
