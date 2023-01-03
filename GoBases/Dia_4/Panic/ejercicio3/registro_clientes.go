package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

//Estructura clientes
type Clientes struct {
	Legajo int
	Nombre string
	DNI string
	Telefono string
	Domicilio string
}

//Tarea1
//array creado
var clientesarray []*Clientes

//Registrar clientes
func RegistrarClientes (Legajo int, Nombre string, DNI string, Telefono string, Domicilio string) *Clientes {
	usuario := &Clientes{
		Legajo: Legajo,
		Nombre: Nombre,
		DNI: DNI,
		Telefono: Telefono,
		Domicilio: Domicilio,
	}
	return usuario
}

//Confirmar si el cliente ya existe
func ClienteYaExiste(cliente *Clientes) bool {
	for _, v := range clientesarray {
		if v.Legajo == cliente.Legajo {
			return true
		}
	}
	return false
}

//Tarea2
//Validar clientes distintos a 0
func ValidarClientes(cliente *Clientes) (bool, error) {
	if cliente.Legajo == 0 {
		return false, errors.New("error: el legajo no debe ser 0")
	}

	if cliente.Nombre == "" {
		return false, errors.New("error: el nombre no debe estar vacio")
	}

	if cliente.DNI == "" {
		return false, errors.New("error: el DNI no debe ser 0")
	}

	if cliente.Telefono == "" {
		return false, errors.New("error: el telefono no debe ser 0")
	}

	if cliente.Domicilio == "" {
		return false, errors.New("error: el domicilio no debe estar vacio")
	}
	return true, nil
}

//Abrir archivo de clientes
func AbrirClientes(archivos string) (err error) {
	archivo, err := os.Open(archivos)
	if err != nil {
		fmt.Println("Error abriendo el archivo", err)
		return
	}

	defer archivo.Close()

	reader := csv.NewReader(archivo)
	registro, err := reader.ReadAll()
	if err !=  nil {
		fmt.Println("Error leyendo registro", err)
		return
	}

	for _, v := range registro {
		legajo, _ := strconv.Atoi(v[0])
		usuario := &Clientes{
			Legajo:     legajo,
			Nombre:     v[1],
			DNI:        v[2],
			Telefono:   v[3],
			Domicilio:  v[4],
		}
		clientesarray = append(clientesarray, usuario)
	}
	return
}

//Guardar archivo de clientes
func GuardarClientes(clientesarray []*Clientes) (err error) {
	archivo, err := os.Create("customer.csv")
	if err != nil {
		fmt.Println("Error creando el archivo", err)
		return
	}

	defer archivo.Close()

	for _, v := range clientesarray {
		_, err := fmt.Fprintf(archivo, "%d,%s,%s,%s,%s\n", v.Legajo, v.Nombre, v.DNI, v.Telefono, v.Domicilio)
		if err != nil {
			return err
		}
	}
	return nil
}

//Imprimir Clientes
func ImprimirClientes(clientesarray []*Clientes) {
	fmt.Println("Clientes:")
	for _, v := range clientesarray {
		fmt.Printf("Legajo: %d, Nombre: %s, DNI: %s, Telefono: %s, Domicilio: %s\n", v.Legajo, v.Nombre, v.DNI, v.Telefono, v.Domicilio)
	}
	fmt.Println()
}

//Tarea3
func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Se detectaron varios errores en tiempo de ejecución.")
		}
		fmt.Println("Ejecución finalizada.")
	}()

	archivos := "customer.csv"

	// Abrir clientes
	AbrirClientes(archivos)

	// Crear nuevos clientes
	nuevoCliente := RegistrarClientes(1, "Jessica Sierra Sáenz", "1022384358", "3213564199", "DG 58 # 45-44")

	// Verificar si el cliente ya existe
	if ClienteYaExiste(nuevoCliente) {
		panic("Error: el cliente ya existe.")
	}

	// Validar Cliente
	validar, err := ValidarClientes(nuevoCliente)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Los datos del cliente son validos. %v\n", validar)
		fmt.Println()
	}

	// Agregar Cliente
	clientesarray = append(clientesarray, nuevoCliente)

	// Imprimir Clientes
	ImprimirClientes(clientesarray)

	// Guardar Clientes
	GuardarClientes(clientesarray)
	fmt.Println("Cliente guardado correctamente")
}

/*
Ejercicio 3 - Registrando clientes
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos
son:

Legajo
Nombre 
DNI
Número de teléfono
Domicilio

Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás leer los datos de un array. En
caso de que esté repetido, debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “Error: el cliente ya existe”, y continuar con la ejecución del programa normalmente.

Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar que todos los datos a
registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de ellos tendrá
que ser del tipo error para el caso de que se ingrese por parámetro algún valor cero (recordá los valores cero de cada tipo de dato,
	ej: 0, “”, nil).

Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes: “Fin de
la ejecución” y “Se detectaron varios errores en tiempo de ejecución”. Utilizá defer para cumplir con este requerimiento.

Requerimientos generales:
Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error. 
Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente
para el caso de error retornado).
*/