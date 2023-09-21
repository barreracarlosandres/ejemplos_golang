package pointers

import "fmt"

func Example(){
	fmt.Println("------------ Pointer Examples ------------")

	value := 1
	v := &value //v get the pointer (position for memory) of input value

	fmt.Println("The value of '*v' is:", *v)
	fmt.Println("The pointer of '&v' is:", &v)
	
	*v = *v + 2 //modify the v value
	fmt.Println("The modify '*v' plus 2 is:", *v)
	fmt.Println("The pointer of modify '&v' plus 2 is:", &v)

	fmt.Println()
}