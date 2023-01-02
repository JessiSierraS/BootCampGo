package main

import (
	"fmt"
)

func salario(salario int) (msg string, err error) {
	minimoImponible := 150000
	if salario < 100000 {
		err = fmt.Errorf("error: el minimo imponible es de %d y el salario ingresado es de: %d", minimoImponible, salario)
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
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg, err)
}
/*
Ejercicio 4 - Impuestos de salario #4
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de
“salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “Error: el mínimo imponible es
de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/