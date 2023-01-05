package main

import (
	"Dia_1/metodo_get/handlers"
	"Dia_1/metodo_get/servicios"
	"github.com/gin-gonic/gin"
)

func main(){
	//Creo un router con gin
	router := gin.Default()
	//2. Creo una variable para la ruta /products
	productos := router.Group("/products")
	//Leer JSON
	servicios.LeerProductJSON()

	//1. Creo un handler para la ruta /ping
	router.GET("/ping", handlers.Ping)
	//2. Creo un handler para la ruta /products usando la variable productos
	productos.GET("/", handlers.ObtenerTodosProducts)
	//3. Creo un handler para la ruta /products/:id 
	productos.GET("/:id", handlers.ObtenerProductsById)
	//4. Creo un handler para la ruta /products/search
	productos.GET("/search", handlers.ObtenerProductsMayor)
	
	router.Run(":8080")
}

/*
Ejercicio 2 : Creando un servidor
Vamos a levantar un servidor utilizando el paquete gin en el puerto 8080. Para probar nuestros endpoints haremos uso de postman.  
1. Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
2. Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
3. Crear una ruta /products/:id que nos devuelva un producto por su id.
4. Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.
*/