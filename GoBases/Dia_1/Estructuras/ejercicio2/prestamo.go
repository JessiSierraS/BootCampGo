package main

import "fmt"

var edad int
var empleado bool
var antiguedad bool
var salario int

func main(){
	edad = 23
	empleado = true
	antiguedad = false
	salario = 100000

	if edad > 22 && empleado && antiguedad {
		if salario > 100000 {
			fmt.Println("Usted puede optar a un prestamo sin interes")
		}else {
			fmt.Println("Usted puede optar a un prestamo con interes")
		}
	}else{
		fmt.Println("Paila prestamo")
	}
}

/*
Ejercicio 2 - Préstamo
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas
reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de
antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior
a $100.000. 
Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/