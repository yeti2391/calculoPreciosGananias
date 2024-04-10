package main

import "fmt"

type Objeto struct {
    Cantidad int
    Precio   float64
}

func main() {
    // Lista de objetos con propiedades de cantidad y precio
    objetos := []Objeto{
            {Cantidad: 20, Precio: 3.80},
        {Cantidad: 10, Precio: 3.80},
        {Cantidad: 20, Precio: 3.11},
        {Cantidad: 5, Precio: 6.53},
        {Cantidad: 20, Precio: 5.80},
        {Cantidad: 20, Precio: 4.10},
        {Cantidad: 3, Precio: 13.50},
        {Cantidad: 2, Precio: 52},
        {Cantidad: 10, Precio: 3},
        {Cantidad: 5, Precio: 3},
        {Cantidad: 3, Precio: 2.2},
        {Cantidad: 20, Precio: 25},
        {Cantidad: 20, Precio: 5},
        // Agrega más objetos si es necesario
    }

    // Calcula el nuevo precio para cada objeto y muestra el valor correspondiente
    var total float64
    for i, obj := range objetos {
                nuevoPrecio := obj.Precio * 4.5
        total += nuevoPrecio * float64(obj.Cantidad) 

        fmt.Printf("Objeto %d - Nuevo precio: %.2f\n", i+1, nuevoPrecio)
    }

    // Muestra el total de los nuevos precios por sus respectivas cantidades
    fmt.Printf("\nTotal: %.2f\n", total)
}
