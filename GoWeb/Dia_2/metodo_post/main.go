package main

import (
	"metodo_post/handlers"
	"metodo_post/servicios"
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
	//5. Creo un handler para la ruta /products/POST
	productos.POST("/", handlers.PostProduct)
	
	router.Run(":8080")
}
/*
Ejercicio 1: Añadir un producto

En esta ocasión vamos a añadir un producto al slice cargado en memoria. Dentro de la ruta /products añadimos el método POST, al cual vamos a enviar en el cuerpo de
la request el nuevo producto. El mismo tiene ciertas restricciones, conozcámoslas:

1. No es necesario pasar el Id, al momento de añadirlo se debe inferir del estado de la lista de productos, verificando que no se repitan ya que debe ser un campo único.
2. Ningún dato puede estar vacío, exceptuando is_published (vacío indica un valor false).
3. El campo code_value debe ser único para cada producto.
4. Los tipos de datos deben coincidir con los definidos en el planteo del problema.
5. La fecha de vencimiento debe tener el formato: XX/XX/XXXX, además debemos verificar que día, mes y año sean valores válidos.

Recordá: si una consulta está mal formulada por parte del cliente, el status code cae en los 4XX.

Ejercicio 2: Traer el producto

Realiza una consulta a un método GET con el id del producto recién añadido, tené en cuenta que la lista de producto se encuentra cargada en la memoria, si terminás
la ejecución del programa este producto no estará en la próxima ejecución.
*/