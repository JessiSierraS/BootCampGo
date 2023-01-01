package main

import (
	"fmt"
)

func horas (minutos int) int{
	horadia := minutos / 60
	return horadia
}

func salarios (minutos int, categoria string) (total float64) {
	var porcentajeb float64 = 1.20
	var porcentajea float64 = 1.50
	horas := horas(minutos)
	if categoria == "C" {
		total = float64(1000 * horas)
	} else if categoria == "B" {
		total = float64(1500 * horas) * porcentajeb
	} else if categoria == "A" {
		total = float64(3000 * horas) * porcentajea
	}
	return total
}

func main(){
	var minutos int
	var categoria string
	fmt.Println("Introduzca cantidad de minutos trabajados: ")
	fmt.Scan(&minutos) 
	fmt.Println("Introduzca la categoria: ")
	fmt.Scan(&categoria)
	resultado := salarios(minutos, categoria)
	fmt.Println("Su salario es de: ", resultado)
}

/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su
salario.
*/