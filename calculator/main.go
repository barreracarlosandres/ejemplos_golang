package calculator

import (
	"fmt"
	"time"
)

type Suma struct{}
type Resta struct{}


func (s Suma) Calcular(d *Datos, c chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("La suma de", d.Dato1, "más", d.Dato2, "es:", (d.Dato1 + d.Dato2))
	c <- "suma"
}

func (s Resta) Calcular(d *Datos, c chan string) {
	time.Sleep(5 * time.Second)
	fmt.Println("La resta de", d.Dato1, "más", d.Dato2, "es:", (d.Dato1 - d.Dato2))
	c <- "resta"
}

// hacer una func suma
// hacer una func resta
// crear una struct para los datos
// crear una struct como receiver y llamar ls func igual, calcular
// crear una interface Operaciones y llamar los métodos
// crear una slice del tipo interface
// hacer un ciclo para recorrer
// adicionar threads
// dicionar channel