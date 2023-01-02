package main

import "fmt"

//Estructura Alumnos
type Alumnos struct {
	Nombre string
	Apellido string
	DNI int
	Fecha string
}

//Método detalle
func (metodo Alumnos) detalle () {
	fmt.Printf("Nombre: %s\n", metodo.Nombre)
	fmt.Printf("Apellido: %s\n", metodo.Apellido)
	fmt.Printf("DNI: %d\n", metodo.DNI)
	fmt.Printf("Fecha: %s\n", metodo.Fecha)
}

//Func main 
func main () {
	alumno1 := Alumnos {
		Nombre: "Jessica",
		Apellido: "Sierra Saenz",
		DNI: 45678990,
		Fecha: "19 de Diciembre del 2022",
	}
	alumno1.detalle()
}

/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada
uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/