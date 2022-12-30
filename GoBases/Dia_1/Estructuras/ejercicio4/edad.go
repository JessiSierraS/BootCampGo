package main

import "fmt"

var employees = map[string]int {
	"Benjamin": 20,
	"Nahuel": 26,
	"Brenda": 19,
	"Darío": 44,
	"Pedro": 30,
}

func main(){
	var nombre string
	fmt.Println("Ingrese el empleado: ")
	fmt.Scanln(&nombre)
	fmt.Println((employees[nombre]))

	contador := 0
	//for k: clave(indice), v: valor de cada llave
	for _, v := range employees {
		if v > 40 {
			contador++
		}
	}
	fmt.Println(contador)

	employees ["Federico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
}

/*
Ejercicio 4 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa,
ayudá a imprimir la edad de Benjamin. 

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario: 
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.
*/