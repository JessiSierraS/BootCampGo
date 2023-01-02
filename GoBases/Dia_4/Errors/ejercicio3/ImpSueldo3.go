package main

import (
	"errors"
	"fmt"
)

var ErrorSalario = errors.New("error: el salario es menor a 10000")

func salario(salario int) (msg string, err error) {
	if salario <= 10000 {
		err = ErrorSalario
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
	if errors.Is(err, ErrorSalario){
		fmt.Println(err)
		return
	}
	fmt.Println(msg, err)
}
/*
Ejercicio 3 - Impuestos de salario #3
Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”, 
se implemente “errors.New()”.
*/