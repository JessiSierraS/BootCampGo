package main

import "fmt"

func promedio (Notas ...float64) float64 {
	var acumulado float64
	for _, nota := range Notas {
		acumulado += nota
	}
	return acumulado / float64(len(Notas))
}

func main() {
	notas := []float64{}
	var cantidad int
	var nota float64
	fmt.Println("Introduzca cantidad de notas: ")
	fmt.Scan(&cantidad)
	for i := 0; i < cantidad; i++ {
		for {
			fmt.Printf("Nota %d: ", i + 1)
			fmt.Scanln(&nota)
			if(nota > 0){
				break
			} else {
				fmt.Println("La nota no puede ser menor a 0")
				fmt.Println("Ingrese nuevamente la nota")
			}
		}
		notas = append(notas, nota)
	}
	resultado := promedio(notas...)
	fmt.Println("Las notas ingrsadas fueron: ", notas)
	fmt.Println("El promedio total es: ", resultado)
}

/*
Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda
pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.
*/