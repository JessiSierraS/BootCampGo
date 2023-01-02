package main

import (
	"errors"
	"fmt"
)

func calcularSalarios (horas int, valor int) (salario int, err error) {
	var impuesto float64 = 0.1
	if horas < 80 || horas < 0 {
		err = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	salario = horas * valor
	if salario >= 150000 {
		salario = salario-(salario*int(impuesto))
	}
	return
}

func verificacion(salario int) (msg string, err error) {
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
	var horas int
	fmt.Print("Ingrese cantidad de horas: ")
	fmt.Scan(&horas)

	var valor int
	fmt.Print("Ingrese el valor por hora: ")
	fmt.Scan(&valor)

	costo, err := calcularSalarios(horas, valor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("El salario mensual es de: %d\n", costo)

	msg, err := verificacion(costo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg, err)
}

/*
Ejercicio 5 -  Impuestos de salario #5
Vamos a hacer que nuestro programa sea un poco más complejo y útil. 
Desarrollá las funciones necesarias para permitir a la empresa calcular:

- Salario mensual de un trabajador según la cantidad de horas trabajadas.
  + La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
  + Dicha función deberá retornar más de un valor (salario calculado y error).
  + En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de impuesto.
  + En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error.
    El mismo tendrá que indicar “Error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
*/