package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
}

// ejemplo 1
func GetTotalTickets(destination string) (pasajero int, err error) {
	archivo, err := os.Open("tickets.csv")
	if err != nil {
		return 
	}
	defer archivo.Close()

	lineas, err := csv.NewReader(archivo).ReadAll()
	if err != nil {
		return 
	}
	defer archivo.Close()	

	for _, r := range lineas{
		if r[3] == destination{
			pasajero += 1
		}
	}
	return
}

func GetCountByPeriod(time string) (cantidad int, err error){
	archivo, err := os.Open("tickets.csv")
	if err != nil {
		return 
	}
	defer archivo.Close()

	lineas, err := csv.NewReader(archivo).ReadAll()
	if err != nil {
		return 
	}
	defer archivo.Close()	

	switch time {
	case "Madrugada":
		cantidad, err = GetMadrugada(lineas)
	case "Mañana":
		cantidad, err = GetMañana(lineas)
	case "Tarde":
		cantidad, err = GetTarde(lineas)
	case "Noche":
		cantidad, err = GetNoche(lineas)
	}
	return
}

func GetMadrugada(lineas [][] string) (cant int, err error) {
		for _, r := range lineas {
			//Atoi Convierte de String a entero
			//itoA de entero a String
			k, err := strconv.Atoi(strings.Split(r[4], ":")[0])
			if err != nil {
				return 0, err
			}
			if k >= 0 && k <= 6 {
				cant += 1
			}
		}
		return
}

func GetMañana(lineas [][] string) (cant int, err error) {
	for _, r := range lineas {
		//Atoi Convierte de String a entero
		//itoA de entero a String
		k, err := strconv.Atoi(strings.Split(r[4], ":")[0])
		if err != nil {
			return 0, err
		}
		if k >= 7 && k <= 12 {
			cant += 1
		}
	}
	return
}

func GetTarde(lineas [][] string) (cant int, err error) {
	for _, r := range lineas {
		//Atoi Convierte de String a entero
		//itoA de entero a String
		//
		k, err := strconv.Atoi(strings.Split(r[4], ":")[0])
		if err != nil {
			return 0, err
		}
		//Aqui defino de que hora a que hora es mañana dentro de un if
		if k >= 13 && k <= 19 {
			cant += 1
		}
	}
	return
}

func GetNoche(lineas [][] string) (cant int, err error) {
	for _, r := range lineas {
		//Atoi Convierte de String a entero
		//itoA de entero a String
		k, err := strconv.Atoi(strings.Split(r[4], ":")[0])
		if err != nil {
			return 0, err
		}
		if k >= 20 && k <= 23 {
			cant += 1
		}
	}
	return
}

// ejemplo 3
func AverageDestination(destination string) (promedio float64, err error) {
	var cantpasajeros int
	var totalpasajeros int
	
	archivo, err := os.Open("tickets.csv")
	if err != nil {
		return 
	}
	defer archivo.Close()

	lineas, err := csv.NewReader(archivo).ReadAll()
	if err != nil {
		return 
	}
	defer archivo.Close()	

	totalpasajeros = len(lineas)
	cantpasajeros, err = GetTotalTickets(destination)
	promedio = float64(cantpasajeros) / float64(totalpasajeros) * 100
	roundPromedio := fmt.Sprintf("%.2f", promedio)
	promedio, err = strconv.ParseFloat(roundPromedio, 64)
	return	
}