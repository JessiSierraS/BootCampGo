package tickets

import (
    "fmt"
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
)

// Requerimiento 1 - TEST CON LISTA DE PRUEBA 
func TestGetTotalTickets_Caso1(t *testing.T) {
    // Creamos un archivo temporal para las pruebas
    f, err := os.Create("tickets.csv")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove("tickets.csv")
    defer f.Close()

    // Escribimos algunos datos de prueba en el archivo
    f.Write([]byte("id,nombre,correo,pais,hora,precio\n"))
    f.Write([]byte("1,Juan,juan@example.com,Inglaterra,10:00,100\n"))
    f.Write([]byte("2,Ana,ana@example.com,Francia,11:00,200\n"))
    f.Write([]byte("3,Pedro,pedro@example.com,Inglaterra,12:00,300\n"))
    f.Write([]byte("4,Sara,sara@example.com,Argentina,13:00,400\n"))

    var pais string = "Francia"
    Resultado, err := GetTotalTickets(pais)
    var ResultadoEsperado int = 1

    if err != nil {
        fmt.Println(err)
    } else {
        assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
    }
}

// Requerimiento 1 - TEST CON LISTA REAL
func TestGetTotalTickets_Caso2(t *testing.T) {
    var pais string = "Poland"
    Resultado, err := GetTotalTickets(pais)
    var ResultadoEsperado int = 36

    if err != nil {
        fmt.Println(err)
    } else {
        assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
    }
}

// Requerimiento 2 - TEST MADRUGADA
func TestGetCountByPeriod_Caso1(t *testing.T) {
    // Creamos un archivo temporal para las pruebas
    f, err := os.Create("tickets.csv")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove("tickets.csv")
    defer f.Close()

    // Escribimos algunos datos de prueba en el archivo
    f.Write([]byte("id,nombre,correo,pais,hora,precio\n"))
    f.Write([]byte("1,Juan,juan@example.com,Inglaterra,10:00,100\n"))
    f.Write([]byte("2,Ana,ana@example.com,Francia,11:00,200\n"))
    f.Write([]byte("3,Pedro,pedro@example.com,Inglaterra,12:00,300\n"))
    f.Write([]byte("4,Sara,sara@example.com,Argentina,13:00,400\n"))

    var Time string = "Madrugada"
    Resultado, err := GetCountByPeriod(Time)
    var ResultadoEsperado int = 0

    if err != nil {
        fmt.Println(err)
    } else {
        assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
    }

    if err != nil {
        fmt.Println(err)
    } else {
        assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
    }
}

// Requerimiento 2 - TEST MAÑANA
func TestGetCountByPeriod_Caso2(t *testing.T) {
    // Creamos un archivo temporal para las pruebas
    f, err := os.Create("tickets.csv")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove("tickets.csv")
    defer f.Close()

    // Escribimos algunos datos de prueba en el archivo
    f.Write([]byte("id,nombre,correo,pais,hora,precio\n"))
    f.Write([]byte("1,Juan,juan@example.com,Inglaterra,10:00,100\n"))
    f.Write([]byte("2,Ana,ana@example.com,Francia,11:00,200\n"))
    f.Write([]byte("3,Pedro,pedro@example.com,Inglaterra,12:00,300\n"))
    f.Write([]byte("4,Sara,sara@example.com,Argentina,13:00,400\n"))

    var Time string = "Mañana"
    Resultado, err := GetCountByPeriod(Time)
    var ResultadoEsperado int = 3

    if err != nil {
        fmt.Println(err)
    } else {
        assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
    }
}

// Requerimiento 2 - TEST TARDE
func TestGetCountByPeriod_Caso3(t *testing.T) {
    // Creamos un archivo temporal para las pruebas
    f, err := os.Create("tickets.csv")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove("tickets.csv")
    defer f.Close()

    // Escribimos algunos datos de prueba en el archivo
    f.Write([]byte("id,nombre,correo,pais,hora,precio\n"))
    f.Write([]byte("1,Juan,juan@example.com,Inglaterra,10:00,100\n"))
    f.Write([]byte("2,Ana,ana@example.com,Francia,11:00,200\n"))
    f.Write([]byte("3,Pedro,pedro@example.com,Inglaterra,12:00,300\n"))
    f.Write([]byte("4,Sara,sara@example.com,Argentina,13:00,400\n"))

    var Time string = "Tarde"
    Resultado, err := GetCountByPeriod(Time)
    var ResultadoEsperado int = 1

    if err != nil {
        fmt.Println(err)
    } else {
        assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
    }
}

// Requerimiento 2 - TEST NOCHE
func TestGetCountByPeriod_Caso4(t *testing.T) {
    // Creamos un archivo temporal para las pruebas
    f, err := os.Create("tickets.csv")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove("tickets.csv")
    defer f.Close()

    // Escribimos algunos datos de prueba en el archivo
    f.Write([]byte("id,nombre,correo,pais,hora,precio\n"))
    f.Write([]byte("1,Juan,juan@example.com,Inglaterra,10:00,100\n"))
    f.Write([]byte("2,Ana,ana@example.com,Francia,11:00,200\n"))
    f.Write([]byte("3,Pedro,pedro@example.com,Inglaterra,12:00,300\n"))
    f.Write([]byte("4,Sara,sara@example.com,Argentina,13:00,400\n"))
    f.Write([]byte("5,Pedro,pedro@example.com,Inglaterra,21:00,300\n"))
    f.Write([]byte("6,Sara,sara@example.com,Argentina,23:00,400\n"))

    var Time string = "Noche"
    Resultado, err := GetCountByPeriod(Time)
    var ResultadoEsperado int = 2

    if err != nil {
        fmt.Println(err)
    } else {
        assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
    }
}

// Requerimiento 3 - TEST PROMEDIO CASO 1
func TestAverageDestination_Caso1(t *testing.T) {
    // Creamos un archivo temporal para las pruebas
    f, err := os.Create("tickets.csv")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove("tickets.csv")
    defer f.Close()

    // Escribimos algunos datos de prueba en el archivo
    f.Write([]byte("1,Juan,juan@example.com,Inglaterra,10:00,100\n"))
    f.Write([]byte("2,Ana,ana@example.com,Francia,11:00,200\n"))
    f.Write([]byte("3,Pedro,pedro@example.com,Inglaterra,12:00,300\n"))
    f.Write([]byte("4,Sara,sara@example.com,Argentina,13:00,400\n"))
    f.Write([]byte("5,Pedro,pedro@example.com,Inglaterra,21:00,300\n"))
    f.Write([]byte("6,Sara,sara@example.com,Argentina,23:00,400\n"))

    var destination string = "Argentina"
    Resultado, err := AverageDestination(destination)
    var ResultadoEsperado float64 = 33.33 

    if err != nil {
        fmt.Println(err)
    } else {
        assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
    }
}

func TestAverageDestination_Caso2(t *testing.T) {
    // Creamos un archivo temporal para las pruebas
    f, err := os.Create("tickets.csv")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove("tickets.csv")
    defer f.Close()

    // Escribimos algunos datos de prueba en el archivo 
    f.Write([]byte("1,Juan,juan@example.com,Inglaterra,10:00,100\n"))
    f.Write([]byte("2,Ana,ana@example.com,Francia,11:00,200\n"))
    f.Write([]byte("3,Pedro,pedro@example.com,Inglaterra,12:00,300\n"))
    f.Write([]byte("4,Sara,sara@example.com,Argentina,13:00,400\n"))
    f.Write([]byte("5,Pedro,pedro@example.com,Inglaterra,21:00,300\n"))
    f.Write([]byte("6,Sara,sara@example.com,Argentina,23:00,400\n"))
    
    var destination string = "Poland"
    Resultado, err := AverageDestination(destination)
    var ResultadoEsperado float64 = 0

    if err != nil {
        fmt.Println(err)
    } else {
        assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
    }
}