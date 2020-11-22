package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	//var array [5]int
	//array := [5]int{10,20,30,40,50}
	//array := [...]int{10,20,30,40,50}
	array := [5]int{1: 10, 2: 20}
	array[2] = 35
	array2 := [5]*int{0: new(int), 1: new(int)}
	*array2[0] = 10
	*array2[1] = 20

	var array3 [5]string
	array3[2] = "coucou"
	array4 := [5]string{"Red","Blue","Green","Yellow","Pink"}
	array3 = array4

	array5 := [2][2]int{{0,1}, {2,3}}
	//here, you can mix all kind of declarations


}
