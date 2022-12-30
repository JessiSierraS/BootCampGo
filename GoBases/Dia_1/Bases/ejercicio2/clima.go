package main

import "fmt"

var temperatura string = "35°"
var humedad string = "65"
var presion int = 45

func main(){
	fmt.Printf("La temperatura es de: %s\n", temperatura)
	fmt.Printf("La humedad es de: %s\n", humedad)
	fmt.Printf("La presion es de: %d\n", presion)
}

/*
Ejercicio 2 - Clima
Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares. 
Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
Imprimí los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?
*/