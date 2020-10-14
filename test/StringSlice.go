package main

import "fmt"

func sliceing(){

	fmt.Println("<= Slice Start =>")
	slice1()
	slice2()
	slice3()
	fmt.Println("<= Slice End =>")

}
func slice1(){
	fmt.Println("<= Slice1 =>")
	s := make([]int,0)
	s = append(s,12,11,10,9,8,7,6,5,4,3,2,1,0)
	fmt.Println(s)
}
func slice2(){
	fmt.Println("<= Slice2 =>")
	s := make([]int,3,10)
	s[1] = 3
	s = append(s,12,11,10,9,8,7,6,5,4,3,2,1,0)
	fmt.Println(s)
	fmt.Println(s[1:7])
	fmt.Println(s[3:4])
	fmt.Println(s[10:12])
}
func slice3(){
	fmt.Println("<= Slice3 =>")
	s := make([]string ,3)
	s[0] = "a"
	s[1] = "b"

	s = append(s,"c","d","e")

	fmt.Println(s)
	fmt.Println(s[4:4])

}