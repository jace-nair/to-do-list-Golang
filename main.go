package main

import "fmt"

func main() {
	//invoke functions from basics.go
	showMyConstInt(myConstInt)
	showMyConstString(myCosntString)
	showMyVarInt(myVarInt)
	showMyVarString(myVarString)
	showMyFloat(myFloat)
	listMyStringArray(myStringArray)
	listMyIntArray(myIntArray)
	listMyStringSlice(myStringSlice)
	listMyIntSlice(myIntSlice)
	listMyStringMap(myStringMap)
	listMyIntMap(myIntMap)
	listMyIntStringMap(myIntStringMap)

	//invoke functions from tasks.go

	fmt.Println("Todo List App")
	fmt.Println() //print an empty line
	taskItems = addTask(taskItems, "Do the dishes")
	taskItems = addTask(taskItems, "Go for a run")
	taskItems = addTask(taskItems, "Lift weights")
	printTasks(taskItems)
}
