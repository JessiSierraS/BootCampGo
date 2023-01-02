package main

import (
	"fmt"
	"time"
)

type Person struct {
	ID int
	Name string
	DateOfBirth time.Time
}

type Employee struct {
	ID int
	Position string
	Person
}

func (metodo *Employee) PrintEmployee(){
	fmt.Printf("ID: %d\n", metodo.ID)
	fmt.Printf("Name: %s\n", metodo.Name)
	fmt.Printf("Date of birth: %s\n", metodo.DateOfBirth.Format("02-01-2006"))
	fmt.Printf("Position: %s\n", metodo.Position)
}

func main (){
	p := Person{ID: 1, Name: "John Smith", DateOfBirth: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)}
	e := Employee{ID: 2, Position: "Manager", Person: p}
	e.PrintEmployee()
}

/*
Ejercicio 2 - Employee
Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará a gestionar
correctamente dichos empleados. Los objetivos son:
Crear una estructura Person con los campos ID, Name, DateOfBirth.
Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.
Realizar el método a la estructura Employee que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos de un
empleado.
Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar el método
PrintEmployee().
Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.
*/