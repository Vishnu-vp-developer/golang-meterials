package main

import ( 
    "fmt"
    "strings"
) 
func pointerMain(){
	fmt.Println("<= Pointer Start =>")
	pointer1()
	pointer2()
	fmt.Println("<= Pointer End =>")
}
func pointer1(){
	fmt.Println("<= Pointer 1 =>")
	var a *int
	b := 10
	a = &b
	fmt.Println("Address : ",a)
	fmt.Println("Value : ",*a)

}

func pointer2(){
	fmt.Println("<= Pointer 2 =>")
	s := "Hellow World .Welcome To World"
	fmt.Println("Initial Value Of S : ",s)
	ChangeString(s,"Golang")
	ChangeString(s,"Programing")
	fmt.Println("Final Result Of S : ",s)

	x := 1
	IncrementVariable(x)
	IncrementVariable(x)
	fmt.Println("Final Result Of IncrementVariable S : ",x)

	IncrementVariableValue(&x)
	IncrementVariableValue(&x)
	fmt.Println("Final Result Of IncrementVariable S : ",x)
}

func ChangeString(s string,r string){
	s = strings.Replace(s, "World", r,1)
	fmt.Println("Inside the ChangeString : ",s)
}

func IncrementVariable(x int){
	x++
	fmt.Println("Inside the IncrementVariable : ",x)
}
func IncrementVariableValue(x *int){
	*x++
	fmt.Println("Inside the IncrementVariable : ",*x)
}