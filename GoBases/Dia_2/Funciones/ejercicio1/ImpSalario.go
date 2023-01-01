package main

import "fmt"
var salario float64

func impuesto(salario float64) float64 {
	impuesto1 := 0.17
	impuesto2 := 0.1

	if salario > 50000 && salario < 150000 {
		resultado := salario - (salario*impuesto1)
		fmt.Println("Te generamos un impuesto por valor de: ",resultado)
		return resultado
	} else if salario > 150000 {
		resultado := salario - (salario*(impuesto1+impuesto2))
		fmt.Println("Te generamos un impuesto por valor de: ", resultado)
		return resultado
	} else {
		fmt.Println("No te generamos impuestos")
		return salario
	}
}

func main() {
	fmt.Println("Ingrese su sueldo: ")
	fmt.Scanln(&salario)
	impuesto(salario)
}

/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto
de sus empleados al momento de depositar el sueldo, para
cumplir el objetivo es necesario crear una función que
devuelva el impuesto de un salario. 
Teniendo en cuenta que si la persona gana más de $50.000
se le descontará un 17 % del sueldo y si gana más de
$150.000 se le descontará además un 10 % (27% en total).
*/