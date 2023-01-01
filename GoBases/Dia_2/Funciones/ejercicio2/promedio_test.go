package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPromedio (t *testing.T){
	notas := []float64{5,4,3,2} 
	//valor que se asigna para validar
	var valoresperado float64

	for _, v := range notas {
		valoresperado += v
	}
	valoresperado = valoresperado/float64(len(notas))
	//variable que guarda el resultado de la función
	Resultado := promedio(notas...)
	assert.Equal(t, valoresperado, Resultado, "El valor esperado y el resultado deberian ser iguales")
}

/*
Ejercicio 2 - Calcular promedio
El colegio informó que las operaciones para calcular el promedio no se están realizando correctamente, por lo que ahora nos
corresponde realizar los test correspondientes:
Calcular el promedio de las notas de los alumnos.
*/