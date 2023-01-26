package main

import (
	"CRUD/internal/products"
	"database/sql"
	"fmt"
	"time"
	"github.com/go-sql-driver/mysql"
)

func main(){
	//Open Database connection
	databaseConfig := mysql.Config{
		User: "root",
		Passwd: "",
		Addr: "localhost:3306",
		DBName: "my_db",
		ParseTime: true,
	}
	println(databaseConfig.FormatDSN())
	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	} 

	//Ping database connection
	if err = database.Ping(); err != nil {
		panic (err)
	}
	
	println("Database connection established")

	//Create products repository.
	var repository products.Repository = &products.MySQLRepository{
		Database: database,
	}
	// Method GET
	product, err := repository.Get(1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Producto solicitado", product)

	//Method STORE
	//formato , fecha a ingresar
	tm,_ := time.Parse("2006-01-02","2023-01-26")
	s := products.Product{
		Name : "Shirt",
		Quantity : 211,
		Code_value : "4444-4444",
		Is_published : "1",
		Expiration : tm,
		Price : 23,
	}

	err = repository.Store(&s)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Insertado con Exito!")
	
	//Method UPDATE
	tm1,_ := time.Parse("2006-01-02","2023-01-26")
	p := products.Product{
		ID: 201,
		Name : "Blusa",
		Quantity : 6,
		Code_value : "05",
		Is_published: "1",
		Expiration: tm1,
		Price : 16,
	}

	err = repository.Update(&p)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Actualizado con Exito!")

	//Method DELETE
	err = repository.Delete(201)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Eliminado con Exito!")
}