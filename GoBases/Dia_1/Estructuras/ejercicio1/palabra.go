package main

import "fmt"

var palabra string

func main(){
	fmt.Print("Ingrese la palabra a validar: ")
	fmt.Scanln(&palabra)
	fmt.Println(len(palabra))

	for _, v := range palabra {
		fmt.Println(string(v))
	}
}

/*
Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado
para deletrearla. Para eso tendrán que:
Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
Luego, imprimí cada una de las letras.
*/