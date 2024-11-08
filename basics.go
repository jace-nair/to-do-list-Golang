package main

import "fmt"

//------------------------------Declare Variables ------------------------------//

const myConstInt int = 2
const myCosntString string = "This is a constant string"

var myVarInt int = 3
var myVarString string = "This is a variable string"

var myFloat float64 = 2.33

var myStringArray [3]string = [3]string{"string1", "string2", "string3"}
var myIntArray [3]int = [3]int{1, 2, 3}

var myStringSlice []string = []string{"string1", "string2", "string3", "string4", "string5"}
var myIntSlice []int = []int{1, 2, 3, 4, 5}

var myStringMap map[string]string = map[string]string{"Name": "Henry", "Occupation": "Student", "Status": "Single"}
var myIntMap map[int]int = map[int]int{1: 100, 2: 200, 3: 300}
var myIntStringMap map[int]string = map[int]string{1: "Jack", 2: "Henry", 3: "David"}

// -----------------------------Declare Functions ------------------------------//
func showMyConstInt(myConstInt int) {
	fmt.Printf("The value of the const int is %v\n", myConstInt)
}

func showMyConstString(myConstString string) {
	fmt.Printf("The value of the const string is %v\n", myConstString)
}

func showMyVarInt(myVarInt int) {
	fmt.Printf("The value of the var int is %v\n", myVarInt)
}

func showMyVarString(myVarString string) {
	fmt.Printf("The value of the var string is %v\n", myVarString)
}

func showMyFloat(myFloat float64) {
	fmt.Printf("The value of the floate is %v\n", myFloat)
}

func listMyStringArray(myStringArray [3]string) {
	// For loop for Array
	for k, v := range myStringArray {
		fmt.Println(k, "-", v)
	}
}

func listMyIntArray(myIntArray [3]int) {
	// For loop for Array
	for k, v := range myIntArray {
		fmt.Println(k, "-", v)
	}
}

func listMyStringSlice(myStringSlice []string) {
	// For loop for Slice
	for k, v := range myStringSlice {
		fmt.Println(k, "-", v)
	}
}

func listMyIntSlice(myIntSlice []int) {
	// For loop for Slice
	for k, v := range myIntSlice {
		fmt.Println(k, "-", v)
	}
}

func listMyStringMap(map[string]string) {
	//For loop for Map
	for k, v := range myStringMap {
		fmt.Println(k, "-", v)
	}
}

func listMyIntMap(map[int]int) {
	//For loop for Map
	for k, v := range myIntMap {
		fmt.Println(k, "-", v)
	}
}

func listMyIntStringMap(map[int]string) {
	//For loop for Map
	for k, v := range myIntStringMap {
		fmt.Println(k, "-", v)
	}
}
