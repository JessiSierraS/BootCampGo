package main

import "fmt"

func forma2(){
	numeromes := map[int]string {
		1: "enero",
		2: "febrero",
		3: "marzo",
		4: "abril",
		5: "mayo",
		6: "junio",
		7: "julio",
		8: "agosto",
		9: "septiembre",
		10: "octubre",
		11: "noviembre",
		12: "diciembre",
	}
	var meselegido int
	fmt.Println("Ingrese el número del mes: ")
	fmt.Scanln(&meselegido)
	fmt.Println((numeromes[meselegido]))
}

func main(){
	meses := [12] string {"enero", "febrero","marzo","abril","mayo","junio","julio","agosto","septiembre","octubre","noviembre","diciembre"}
	var nummes int
	fmt.Println("Ingrese el número del mes: ")
	fmt.Scanln(&nummes)
	fmt.Println((meses[nummes-1]))

	forma2()
}

/*
Ejercicio 3 - A qué mes corresponde
Realizar una aplicación que reciba  una variable con el número del mes. 
Según el número, imprimir el mes que corresponda en texto. 
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.
Nota: Validar que el número del mes, sea correcto.
*/