package main

import (
	"errors"
	"fmt"
)

const (
   dog       = "dog"
   cat       = "cat"
   hamster   = "hamster"
   tarantula = "taramtula"
)

func animal(animal string) (func(cantidad float64) float64, error ) {
   switch animal {
   case dog: 
      return animalDog, nil
   case cat:
      return animalCat, nil
   case hamster:
      return animalHamster, nil
   case tarantula:
      return animalTarantula, nil
   default:
      return nil, errors.New("animal no existe")
   }
}

func animalDog(cantidad float64) float64 {
   return cantidad * 10
}

func animalCat(cantidad float64) float64 {
   return cantidad * 5
}

func animalHamster(cantidad float64) float64 {
   return (cantidad * 250) / 1000
}

func animalTarantula(cantidad float64) float64 {
   return (cantidad * 150) / 1000
}

func main(){
   animalDog, msg := animal(dog)
	if msg != nil {
		fmt.Println(msg.Error())
	}

   animalCat, msg := animal(cat)
	if msg != nil {
		fmt.Println(msg.Error())
	}

   animalHamster, msg := animal(hamster)
	if msg != nil {
		fmt.Println(msg.Error())
	}

   animalTarantula, msg := animal(tarantula)
	if msg != nil {
		fmt.Println(msg.Error())
	}
 
   var amount float64
   amount += animalDog(10)
   amount += animalCat(10)
   amount += animalHamster(5)
   amount += animalTarantula(8)

   fmt.Println("La cantidad de comida que debe comprar es de:", amount, "Kg")
}

/*
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas,
hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.

Tienen los siguientes datos:

Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.

Se solicita:

Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una
función y un mensaje (en caso que no exista el animal).

Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

Ejemplo:
const (
   dog    = "dog"
   cat    = "cat"
)
 
...
 
animalDog, msg := animal(dog)
animalCat, msg := animal(cat)
 
...
 
var amount float64
amount += animalDog(10)
amount += animalCat(10)
*/