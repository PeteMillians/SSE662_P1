package main

import (
	taskManager "SSE662_P1/TaskManager"
	"fmt"
	"strings"
)

// Constant for valid user inputs
var VALID_INPUTS = []string{"A", "L", "M", "D", "H"}

func main() {
	var manager taskManager.TaskManager

	commandLine(manager)
}

func commandLine(manager taskManager.TaskManager) {

	// Starts the loop to run the command-line interface

	// Arguments:
	//     manager (TaskManager): the TaskManager instance which controls the list of Tasks

	fmt.Println(manager.Help())

	// Iterate until a break
	// for {

	// }

}

func validate(entry string) bool {

	// Validates a user input string against known options

	// Arguments:
	//     entry (string): the entry entered by the user
	// Returns:
	//     True if it is valid, False otherwise

	// Check the entry against every valid input
	for _, option := range VALID_INPUTS {

		// If the option is valid, return True
		if option == strings.TrimSpace(entry) {
			return true
		}
	}

	// If the entry was not in the valid inputs list, return False
	return false
}
