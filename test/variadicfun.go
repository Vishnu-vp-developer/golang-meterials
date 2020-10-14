package main

import ( 
    "fmt"
)

func variadicmain(){
	fmt.Println("<= Variadic Example =>")
	sum(1,2,3,4,5,6)
}
func sum(a ...int){
	fmt.Println("Values : ",a)
	sum := 0
	for _,a := range a {
		sum += a
	}
	fmt.Println("Total : ",sum)
}