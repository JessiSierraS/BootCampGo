package main

import "fmt"

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible"
}

func main () {
	e := myError{"Error"}
	var salary int = 170000
	if salary < 150000 {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("Debe pagar impuesto")
}

/*
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario ingresado no alcanza el mínimo
imponible" y lanzalo en caso de que “salary” sea menor a 150.000. De lo contrario, tendrás que imprimir por consola el mensaje
“Debe pagar impuesto”.
*/