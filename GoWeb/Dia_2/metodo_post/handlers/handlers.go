package handlers

import (
	"errors"
	"metodo_post/articulos"
	"metodo_post/servicios"
	"net/http"
	"strconv"
	"strings"
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

//5. Función Post para crear un producto
func PostProduct(ctx *gin.Context) {
		
	var product articulos.Articulos

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "producto invalido"})
		return
	}

	valid, err := validateEmptys(&product)
	if !valid {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	valid, err = validateExpiration(&product)
	if !valid {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	valid = validateCodeValue(product.Codigo_Valor)
	if !valid {
		ctx.JSON(400, gin.H{"error": "el valor del codigo ya existe"})
		return
	}
	
	product.ID = len(servicios.CatalogoProductos.Productos) + 1
	servicios.CatalogoProductos.Productos = append(servicios.CatalogoProductos.Productos, product)
	ctx.JSON(201, product)
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(product *articulos.Articulos) (bool, error) {
	switch {
	case product.Nombre == "" || product.Codigo_Valor == "" || product.Vencimiento == "":
		return false, errors.New("los campos no pueden estar vacios")
	case product.Cantidad <= 0 || product.Precio <= 0:
		if product.Cantidad <= 0 {
			return false, errors.New("la cantidad debe ser mayor que 0")
		}
		if product.Precio <= 0 {
			return false, errors.New("la cantidad debe ser mayor que 0")
		}
	}
	return true, nil
}

// validateExpiration valida que la fecha de expiracion sea valida
func validateExpiration(product *articulos.Articulos) (bool, error) {
	dates := strings.Split(product.Vencimiento, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("fecha de caducidad no válida, debe estar en formato: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("fecha de caducidad no válida, deben ser números")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("fecha de caducidad no válida, la fecha debe estar entre el 1 y el 31/12/9999")
	}
	return true, nil
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func validateCodeValue(Codigo_Valor string) bool {
	for _, product := range servicios.CatalogoProductos.Productos {
		if product.Codigo_Valor == Codigo_Valor {
			return false
		}
	}
	return true
}