package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAnimalDog(t *testing.T) {
	var kg float64 = 10
	var cantidad float64 = 10

	valor_esperado := cantidad * kg

	respuesta := animalDog(cantidad)
	assert.Equal(t, valor_esperado, respuesta, "Cantidad Dog incorrecta")
}

func TestAnimalCat(t *testing.T) {
	var kg float64 = 5
	var cantidad float64 = 10

	valor_esperado := cantidad * kg

	respuesta := animalCat(cantidad)
	assert.Equal(t, valor_esperado, respuesta, "Cantidad Gato incorrecta")
}

func TestAnimalHamster(t *testing.T) {
	var kg float64 = 250
	var cantidad float64 = 10

	valor_esperado := (cantidad * kg) / 1000

	respuesta := animalHamster(cantidad)
	assert.Equal(t, valor_esperado, respuesta, "Cantidad Hamster incorrecta")
}

func TestAnimalTarantula(t *testing.T) {
	var kg float64 = 150
	var cantidad float64 = 10

	valor_esperado := (cantidad * kg) / 1000

	respuesta := animalTarantula(cantidad)
	assert.Equal(t, valor_esperado, respuesta, "Cantidad Tarantula incorrecta")
}