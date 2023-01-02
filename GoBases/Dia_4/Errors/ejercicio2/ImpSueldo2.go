package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return "Error: el salario es menor a 10000"
}

func salario(salario int) (msg string, err error) {
	if salario <= 10000 {
		err = &myError{"Error: el salario es menor a 10000"}
		return
	} else {
		msg = "OK"
		return
	}
}

func main () {
	var salary int
	fmt.Print("Ingrese el salario: ")
	fmt.Scan(&salary)

	msg, err := salario(salary)
	if errors.Is(err, &myError{}){
		fmt.Println(err)
		return
	}
	fmt.Println(msg, err)
}

/*
Ejercicio 2 - Impuestos de salario #2
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario es menor a 10.000" y lanzalo
en caso de que “salary” sea menor o igual a  10.000. La validación debe ser hecha con la función “Is()” dentro del “main”.
*/