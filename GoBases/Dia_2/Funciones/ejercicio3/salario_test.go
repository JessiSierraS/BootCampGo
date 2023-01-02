package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSalario(t *testing.T){
	categoria1 := "A"
	categoria2 := "B"
	categoria3 := "C"
	var valor1 int = 180
	var valor2 int = 360
	var valor3 int = 540

	var porcentajeb float64 = 1.20
	var porcentajea float64 = 1.50
	hora1 := valor1 / 60
	hora2 := valor2 / 60
	hora3 := valor3 / 60
	var salarioesp1 float64
	var salarioesp2 float64
	var salarioesp3 float64

	salarioesp1 = float64(1000 * hora1)
	salarioesp2 = float64(1500 * hora2) * porcentajeb
	salarioesp3 = float64(3000 * hora3) * porcentajea

	total1 := salarios(valor1, categoria1)
	total2 := salarios(valor2, categoria2)
	total3 := salarios(valor3, categoria3)

	assert.Equal(t, salarioesp1, total1, "Salario incorrecto para categoria C")
	assert.Equal(t, salarioesp2, total2, "Salario incorrecto para categoria B")
	assert.Equal(t, salarioesp3, total3, "Salario incorrecto para categoria A")
}

/*
Ejercicio 3 - Test del salario
La empresa marinera no está de acuerdo con los resultados obtenidos en los cálculos de los salarios, por ello nos piden realizar una
serie de tests sobre nuestro programa. Necesitaremos realizar las siguientes pruebas en nuestro código:
Calcular el salario de la categoría “A”.
Calcular el salario de la categoría “B”.
Calcular el salario de la categoría “C”.
*/