package main

import "fmt"

var taskItems []string = []string{}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for k, v := range taskItems {
		fmt.Printf("%v: %v\n", k+1, v)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
