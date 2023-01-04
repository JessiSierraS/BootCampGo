package main

import (
	"github.com/gin-gonic/gin"
)

	type Estructura struct {
		Nombre string
		Apellido string
	}

func main (){
	//Router creado
	router := gin.Default()

	//Capturar solicitud
	/*
	router.GET("/ping", func(c *gin.Context){c.String(200, "/pong")})
	router.Run(":8080")
	*/
	//

	router.POST("/saludo", func(c *gin.Context){

		var info Estructura

		if err := c.ShouldBindJSON(&info); err != nil {
			c.String(400, "Solicitud invalida")
			return
		}

		//imprimir := fmt.Sprintf(info.nombre, info.apellido)
		c.String(200, "Hola "+info.Nombre+" "+info.Apellido)
		//c.String(200, imprimir)
	})
	router.Run(":8080")
}

/*
Ejercicio 1 - Prueba de Ping
Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
El endpoint deberá ser de método GET
La respuesta de “pong” deberá ser enviada como texto, NO como JSON

------
BindJSON siempre intentará vincular el cuerpo de la solicitud a un struct, independientemente de si el cuerpo de la solicitud es válido o no
Si hay un error al vincular el cuerpo de la solicitud, BindJSON devolverá el error sin proporcionar una respuesta HTTP al cliente.
ShouldBindJSON del paquete gin también se utiliza para vincular el cuerpo de una solicitud HTTP a un struct de Go. Sin embargo, si hay un error al vincular el cuerpo de la solicitud, ShouldBindJSON devolverá el error y enviará una respuesta HTTP al cliente con el código de estado HTTP correspondiente.

Ejercicio 2 - Manipulando el body

Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en
texto “Hola + nombre + apellido”

El endpoint deberá ser de método POST
Se deberá usar el package JSON para resolver el ejercicio
La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
La estructura deberá ser como esta:
{
		“nombre”: “Andrea”,
		“apellido”: “Rivas”
}

------
*/


