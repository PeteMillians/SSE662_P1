package main

import (
	"SSE662_P1/Task"
	taskManager "SSE662_P1/TaskManager"
	"fmt"
	"strings"
)

func main() {
	manager := &taskManager.TaskManager{}

	commandLine(manager)
}

func commandLine(manager *taskManager.TaskManager) {

	// Starts the loop to run the command-line interface

	// Arguments:
	//     manager (TaskManager): the TaskManager instance which controls the list of Tasks

	// Iterate until a break
	for {

		// Clear the terminal screen
		clearScreen()

		var input string

		fmt.Print("Enter a command: ")
		fmt.Scanln(&input)

		processCommand(strings.TrimSpace(input), manager)
	}

}

func processCommand(input string, manager *taskManager.TaskManager) {
	/*
		Processes a user command to send to the Task Manager

		Arguments:
			input (string): the valid user input
			manager (TaskManager): the current TaskManager object
	*/

	clearScreen()

	// Switch statement to process commands
	switch input {

	// Add Task
	case "A":

		// Instantiate a new Task object
		var task Task.Task = createTask(manager)

		// Add the Task to the TaskManager
		manager.Add(task)

	case "L":

		fmt.Println("Tasks: ")
		fmt.Println("------")

		// Prints the list of the available Tasks
		for _, str := range manager.List() {

			// Print the task
			fmt.Println(str)
		}

	case "M":

		// Print the available tasks
		fmt.Println("Available Tasks to mark: ")
		for _, str := range manager.List() {
			fmt.Println(str)
		}

		fmt.Println()

		// Take input for the Task we are marking
		var number int
		fmt.Print("Task Number: ")
		fmt.Scanln(&number)
		manager.Mark(number)

	case "D":

		// Print the available tasks
		fmt.Println("Available Tasks to delete: ")
		for _, str := range manager.List() {
			fmt.Println(str)
		}

		fmt.Println()

		// Take input for the Task we are deleting
		var number int
		fmt.Print("Task Number: ")
		fmt.Scanln(&number)
		manager.Delete(number)

	case "H":

		// Print the help screen
		fmt.Println(manager.Help())

	default:

		fmt.Println("Unrecognized entry. Use 'H' to display a help menu.")
	}

	fmt.Println()
	fmt.Print("Press Enter to continue...")
	fmt.Scanln()

}

func createTask(manager *taskManager.TaskManager) Task.Task {
	/*
		Creates a custom Task from user input

		Arguments:
			manager (TaskManager): the TaskManager we are creating a task for

		Returns:
			an instantied Task object
	*/
	fmt.Println("Task Creation Screen")

	// Get the Task name
	var Name string
	fmt.Print("Name: ")
	fmt.Scanln(&Name)

	// Get the Task name
	var Description string
	fmt.Print("Description: ")
	fmt.Scanln(&Description)

	// Instantiate a Task and return it
	return Task.Task{
		Name:        Name,
		Description: Description,
		Status:      false,
		Number:      manager.GetSize() + 1,
	}

}

func clearScreen() {
	/*
		Clears the terminal screen
	*/

	fmt.Print("\033[H\033[2J")
}
