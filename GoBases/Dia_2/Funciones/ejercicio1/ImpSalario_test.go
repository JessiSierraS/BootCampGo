package  main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestImpuestomenor (t *testing.T){
	//valor que se asigna para validar
	var salario float64 = 20000
	//resultado esperado
	var resultadoEsperado float64 = salario

	//variable que guarda el resultado de la función
	Resultado := impuesto(salario)
	assert.Equal(t, resultadoEsperado, Resultado, "El valor del resultado y el resultado esperado deberian ser iguales")
}

func TestImpuestomayor (t *testing.T){
	var salario float64 = 70000
	var resultadoEsperado float64 = 58100

	Resultado := impuesto(salario)
	assert.Equal(t, resultadoEsperado, Resultado, "El valor del resultado y el resultado esperado deberian ser iguales")
}

func TestImpuesto (t *testing.T){
	var salario float64 = 160000
	var resultadoEsperado float64 = 116800

	Resultado := impuesto(salario)
	assert.Equal(t, resultadoEsperado, Resultado, "El valor del resultado y el resultado esperado deberian ser iguales")
}

/*
La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados, al momento de
depositar el sueldo de los mismos ahora nos solicitó validar que los cálculos de estos impuestos están correctos. Para esto nos 
encargaron el trabajo de realizar los test correspondientes para:

Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $150.000.
*/