package calculator

import (
	"fmt"
	"time"
)

type Add struct{}
type Subtraction struct{}

func (s Add) Calculator(d *Data, c chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("The add of", d.First, "plus", d.Second, "is :", d.First+d.Second)
	c <- "add"
}

func (s Subtraction) Calculator(d *Data, c chan string) {
	time.Sleep(5 * time.Second)
	fmt.Println("The subtraction of", d.First, "lest", d.Second, "is:", d.First-d.Second)
	c <- "subtraction"
}
