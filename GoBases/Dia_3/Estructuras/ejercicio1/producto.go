package main

import "fmt"

type Product struct {
	ID int
	Name string
	Price float64
	Description string
	Category int
}

var Products []*Product

func (p *Product) Save(Product) {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Println(*product)
	}
}

func getByld(ID int) *Product {
	for _, prod := range Products {
		if prod.ID == ID {
			return prod
		}
	}
	panic("Articulo no econtrado")
}

func main() {

	p1 := Product{ID: 1, Name: "Lapiz", Price: 2.00, Description: "Color rojo", Category: 1}
	p1.Save(p1)
	p1.GetAll()
	p1.Name = "Aaaaaaaaaa" // p1 = Product{ID: 1, Name: "AAAAAAALapiz", Price: 2.00, Description: "Color rojo", Category: 1}
	p1.GetAll()
	fmt.Println("---------------------")

	p2 := Product{ID: 2, Name: "Lapiz", Price: 2.00, Description: "Color azul", Category: 2}
	p2.Save(p2)
	p2.GetAll()
	fmt.Println("---------------------")

	fmt.Println(*getByld(1))
	fmt.Println(*getByld(2))
	fmt.Println(*getByld(3))
}
/*
Ejercicio 1
Crear un programa que cumpla los siguiente puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.   
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto
desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado. 
Ejecutar al menos una vez cada método y función definido desde main().
*/