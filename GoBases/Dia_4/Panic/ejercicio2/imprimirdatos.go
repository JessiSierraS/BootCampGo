package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//defer 
	defer func() {
		if err := recover(); err != nil {
				fmt.Println(err)
				fmt.Println("La ejecución finalizó")
		}
	}()
	
	//Abrir y cerrar el archivo
	archivo, err := os.Open("customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	defer archivo.Close()

	fmt.Println(archivo)
	content, err := io.ReadAll(archivo)
	fmt.Print(string(content))
	fmt.Println(err)
}
/*
Ejercicio 2 - Imprimir datos
A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio. 
Ahora que el archivo sí existe, el panic no debe ser lanzado. 
Creamos el archivo “customers.txt” y le agregamos la información de los clientes. 
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga. En el caso de no poder leerlo,
se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer” que nos indique que la
ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.
*/