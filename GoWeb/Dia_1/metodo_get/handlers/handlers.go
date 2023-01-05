package handlers

import (
	"Dia_1/metodo_get/articulos"
	"Dia_1/metodo_get/servicios"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

//Estructura Respuesta
type Respuesta struct {
	Mensaje string `json:"mensaje"`
	Datos any `json:"data"`
}

//1. Función Ping para respuesta pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}

//2. Función ObtenerTodosProductos
func ObtenerTodosProducts(c *gin.Context) {
	products := servicios.CatalogoProductos
	c.JSON(http.StatusOK, Respuesta{Mensaje: "Productos Obtenidos", Datos: products})
}

//3. Función ObtenerProductosporID
func ObtenerProductsById(c *gin.Context) {
	//convertir en entero al recibir en string
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Respuesta{Mensaje: "Error de solicitud"})
		return
	}

	var productoRetornar articulos.Articulos

	//Recorrer la lista de productos en el catalogo
	for _, v := range servicios.CatalogoProductos.Productos {
		if v.ID == Id {
			productoRetornar = v
			break
		}
	}

	//Retornar producto o error si encuentra producto.
	if productoRetornar.ID != 0 {
		c.JSON(http.StatusOK, Respuesta{Mensaje: "Producto obtenido por ID", Datos: productoRetornar})
		return
	} else {
		c.JSON(http.StatusNotFound, Respuesta{Mensaje: "Error al obtener producto por ID"})
		return
	}
}

//4. Función ObtenerProductsMayor
func ObtenerProductsMayor(c *gin.Context) {
	////convertir en flotante al recibir en string
	precio, err := strconv.ParseFloat(c.Query("priceGt"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Respuesta{Mensaje: "Parametro Invalido"})
		return
	}

	//Recorrer la lista de productos en el catalogo
	precioRetornar := make([]articulos.Articulos, 0)
	for _, v := range servicios.CatalogoProductos.Productos {
		if precio != 0 && v.Precio >= precio {
			precioRetornar = append(precioRetornar, v)
		}
	}
	c.JSON(http.StatusOK, Respuesta{Mensaje: "Productos obtenidos", Datos: precioRetornar})
}