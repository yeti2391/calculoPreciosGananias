package main

import (
    "fmt"
)

type Objeto struct {
    Cantidad int
    Precio   float64
}

func main() {
    // Lista de objetos con propiedades de cantidad y precio
    objetos := []Objeto{
        {Cantidad: 5, Precio: 100},
        {Cantidad: 10, Precio: 50},
        // Agrega más objetos si es necesario
    }

    // Variable costoTodosObjetos
    costoTodosObjetos := 1176

    // Calcula el costo proporcional al objeto y el costo adicional por unidad para cada objeto
    var gananciaTotal float64
    listaObjetos := make([]Objeto, len(objetos))
    for i, obj := range objetos {
        costoProporcional := ((obj.Precio * float64(obj.Cantidad)) / float64(costoTodosObjetos)) * 761
        costoAdicionalPorUnidad := costoProporcional / float64(obj.Cantidad)
        nuevoCosto := obj.Precio + costoAdicionalPorUnidad
        precioVenta := nuevoCosto * (1 + 0.50)

        gananciaTotal += nuevoCosto

        // Guarda el objeto con su nuevo costo en la lista de objetos
        listaObjetos[i] = Objeto{Cantidad: obj.Cantidad, Precio: nuevoCosto}

        fmt.Printf("Objeto %d:\n", i+1)
        fmt.Printf(" - Nuevo costo: %.2f\n", nuevoCosto)
        fmt.Printf(" - Precio de venta: %.2f\n\n", precioVenta)
    }

    // Muestra la ganancia total
    fmt.Printf("Ganancia total: %.2f\n", gananciaTotal)
}
