package main

import "fmt"

const (
   minimum = "minimum"
   average = "average"
   maximum = "maximum"
)

func minFunc(valores ... float64) float64 {
	min := valores[0]
	for _, v := range valores {
		if v < min {
			min = v
		}
	}
	return min
}

func averageFunc(valores ... float64) (average float64) {
	for _, v := range valores {
		average += float64(v)
	}
	return average / float64(len(valores))
}

func maxFunc(valores ...float64) float64 {
	max := valores[0]
	for _, v := range valores {
		if v > max {
			max = v
		}
	}
	return max
}

func  operation(operacion string) (func(valores ...float64) float64, error) {
   switch operacion {
   case minimum:
      return minFunc, nil
   case average:
      return averageFunc, nil
   case maximum:
      return maxFunc, nil
   }
   return nil, nil
}

func main (){
   minFunc, err1 := operation(minimum)
   averageFunc, err2 := operation(average)
   maxFunc, err3 := operation(maximum)
 
   fmt.Println(err1, err2, err3)

   minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
   averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
   maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

   fmt.Printf("Max: %.2f,\nMin: %.2f, \nAverage: %.2f\n", maxValue, minValue, averageValue)
}

/*
Ejercicio 4 - Calcular estadísticas
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un
curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva
otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el
cálculo que se indicó en la función anterior.
Ejemplo:

const (
   minimum = "minimum"
   average = "average"
   maximum = "maximum"
)

...
 
minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...

minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/