package main

import "fmt"

//Interfaz
type Producto interface {
	precio()float64
}

//Estructura producto pequeño
type Pequeño struct {
	PrecioUnidad float64
}

//Metodo para la estructura de producto pequeño
func (metodoP Pequeño) precio () float64 {
	return metodoP.PrecioUnidad
}

//Estructura producto mediano
type Mediano struct {
	PrecioUnidad float64 
}

//Metodo para la estructura de producto mediano
func (metodoM Mediano) precio () float64 {
	return metodoM.PrecioUnidad * 1.03
}

//Estructura producto grande
type Grande struct {
	PrecioUnidad float64
}

//Metodo para la estructura de producto grande
func (metodoG Grande) precio () float64 {
	return 2500 + metodoG.PrecioUnidad * 1.06
}

//Función Factory
func factory (tipo string, precio float64) (metodo Producto) {
	switch tipo {
	case "Pequeño":
		metodo = Pequeño{precio}
	case "Mediano":
		metodo = Mediano{precio}
	case "Grande": 
		metodo = Grande{precio}
	}
	return metodo
}

//Func Main lleno.
func main() {
	articulo1 := factory("Grande", 500)
	precio1 := articulo1.precio()
	fmt.Println("El precio es: ", precio1)
}

/*
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:
Pequeño: solo tiene el costo del producto
Mediano: el precio del producto + un 3% de mantenerlo en la tienda
Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.
*/