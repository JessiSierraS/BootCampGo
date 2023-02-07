package hunt

// En el import de paquetes agregar -> "testing", "github.com/stretchr/testify/assert"
// En la terminal se debe ejecutar el siguiente comando para instalar el go mod -> go mod init Ejercicio1
// En la terminal se debe ejecutar el siguiente comando para instalar la librería testify -> go get github.com/stretchr/testify
// En la terminal se debe ejecutar el siguiente comando para instalar cualquier librería necesaria que no se haya incluído -> go mod tidy

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestSharkHuntsSuccessfully, El Tiburón cazó la Presa exitosamente
func TestSharkHuntsSuccessfully(t *testing.T) {
	//			INICIALIZACIÓN DE DATOS PARA TEST: DECLARACIÓN DE PARÁMETROS DE VALIDACIÓN
	// Parámetros a validar
	sh := Shark{
		hungry: false,
		tired: true,
		speed: 100,
	}

	pr := Prey{
		name: "Salmón",
		speed: 90,
	}
	// 			EJECUCIÓN DEL TEST: LLAMADO DE LA FUNCIÓN
	// Se declara la variable que guarda el resultado de la función
	errResultado := sh.Hunt(&pr)

	// 			VALIDACIÓN DE RESULTADOS ESPERADOS
	// Esta función compara dos errores y deben ser iguales
	assert.NoError(t, errResultado)
	assert.Equal(t, true, sh.tired)
	assert.Equal(t, false, sh.hungry)
}

// TestSharkCannotHuntBecauseIsTired, El Tiburón no cazó la Presa porque estaba cansado
func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	//			INICIALIZACIÓN DE DATOS PARA TEST: DECLARACIÓN DE PARÁMETROS DE VALIDACIÓN
	// Resultado esperado
	errEsperado := errors.New("cannot hunt, i am really tired")
	// Parámetros a validar
	sh := Shark{
		hungry: false,
		tired: true,
		speed: 90,
	}
	pr := Prey{
		name: "Cachama",
		speed: 100,
	}
	// 			EJECUCIÓN DEL TEST: LLAMADO DE LA FUNCIÓN
	// Se declara la variable que guarda el resultado de la función
	errResultado := sh.Hunt(&pr)

	// 			VALIDACIÓN DE RESULTADOS ESPERADOS
	// Esta función compara dos errores y deben ser iguales
	assert.EqualError(t, errEsperado, errResultado.Error())
}

// TestSharkCannotHuntBecaisIsNotHungry, El Tiburón no cazó la Presa porque no tenía hambre
func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	// 			INICIALIZACIÓN DE DATOS PARA TEST: DECLARACIÓN DE PARÁMETROS DE VALIDACIÓN
	// Resultado esperado
	errEsperado := errors.New("cannot hunt, i am not hungry")
	// Parámetros a validar
	sh := Shark{
		hungry: false,
		tired: false,
		speed: 100,
	}
	pr := Prey{
		name: "Trucha",
		speed: 90,
	}
	// 			EJECUCIÓN DEL TEST: LLAMADO DE LA FUNCIÓN
	// Se declara la variable que guarda el resultado de la función
	errResultado := sh.Hunt(&pr)

	// 			VALIDACIÓN DE RESULTADOS ESPERADOS
	// Esta función compara dos errores y deben ser iguales
	assert.EqualError(t, errEsperado, errResultado.Error())
}

// TestSharkCannotReachThePrey, El Tiburón no pudo alcanzar la Presa
func TestSharkCannotReachThePrey(t *testing.T) {
	// 			INICIALIZACIÓN DE DATOS PARA TEST: DECLARACIÓN DE PARÁMETROS DE VALIDACIÓN
	// Resultado esperado
	errEsperado := errors.New("could not catch it")
	// Parámetros a validar
	sh := Shark{
		hungry: true,
		tired: false,
		speed: 90,
	}
	pr := Prey{
		name: "Delfin",
		speed: 100,
	}
	// 			EJECUCIÓN DEL TEST: LLAMADO DE LA FUNCIÓN
	// Se declara la variable que guarda el resultado de la función
	errResultado := sh.Hunt(&pr)

	// 			VALIDACIÓN DE RESULTADOS ESPERADOS
	// Esta función compara dos errores y deben ser iguales
	assert.EqualError(t, errEsperado, errResultado.Error())
}

// Prey Nil
func TestSharkHuntNilPrey(t *testing.T) {
	// 			INICIALIZACIÓN DE DATOS PARA TEST: DECLARACIÓN DE PARÁMETROS DE VALIDACIÓN
	// Resultado esperado
	errEsperado := errors.New("cannot hunt, prey nil")
	// Parámetros a validar
	sh := Shark{
		hungry: true,
		tired: false,
		speed: 90,
	}

	// 			EJECUCIÓN DEL TEST: LLAMADO DE LA FUNCIÓN
	// Se declara la variable que guarda el resultado de la función
	errResultado := sh.Hunt(nil)

	// 			VALIDACIÓN DE RESULTADOS ESPERADOS
	// Esta función compara dos errores y deben ser iguales
	assert.EqualError(t, errEsperado, errResultado.Error())
}

/*
		Ejercicio 1
Testear los siguientes casos negativos:
	1. El tiburón está cansado.
	2. El tiburón ya comió y está lleno.
	3. La presa es más rápida que el tiburón.
Validar no solo que la función devuelve un error, sino también que el error sea el correspondiente.

		Ejercicio 2
Testear el caso feliz, donde el tiburón caza a la presa. Se debe validar que una vez que cazó a la presa, el tiburón se llenó y se cansó.

		Ejercicio 3
¿Qué sucede si la presa es nula? En un caso ideal se debería devolver un error indicando el problema. Sin embargo,
el programa posee un error de diseño. Realizar un test unitario para este caso, donde se debe validar que la función
devuelve un error.
¿Qué sucede cuando corremos el test? Corregir el test para que pase, realizando las validaciones necesarias.
Finalmente, volver al test anterior y corregir el código para que el test pase, devolviendo un error sin panic.
*/