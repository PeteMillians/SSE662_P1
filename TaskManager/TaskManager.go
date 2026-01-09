package taskManager

import (
	"SSE662_P1/Task"
)

const HELP_MENU = "Welcome to the Task Manager! Available commands are listed below: \n\tA: Add a task\n\tL: List all current tasks\n\tM: Mark a task as complete\n\tD: Delete a task\n\tH: Display the help menu\n\n*Note: These commands are case-sensitive"

type TaskManager struct {
	tasks  []Task.Task
	number int
}

func (t TaskManager) GetSize() int {
	/*
		Gets the number of Tasks in the slice

		Returns:
			the int length of the tasks list
	*/

	return len(t.tasks)
}

func (t *TaskManager) Add(task Task.Task) {
	/*
		Adds a given Task to the slice

		Arguments:
		    task (Task): the Task we want to add
	*/

	t.tasks = append(t.tasks, task)
}

func (t *TaskManager) List() []string {
	/*
		Returns a string array of Tasks in the slice

		Returns:
		    a slice of strings containing the required Task information
	*/

	// Instantiate an empty list of string
	list := []string{}

	// Iterate through each task
	for _, task := range t.tasks {

		// Add the __str__ implementation of the Task to the list
		list = append(list, task.String())
	}

	// Return the list
	return list
}

func (t *TaskManager) Mark(number int) {
	/*
		Marks the Task at a given number complete

		Arguments:
		    number (int): the Task number we are marking as complete
	*/

	// Find the index of the requested number
	index := t.find(number)

	// if index == -1 {
	// 	err := errors.New("Could not find Task at given number")
	// }

	// Set the status of that index to true
	t.tasks[index].SetStatus(true)

}

func (t *TaskManager) Delete(number int) {
	/*
		Deletes a Task at a given number

		Arguments:
		    number (int): the Task number we are deleting
	*/

	// Find the index of the requested number
	index := t.find(number)

	// if index == -1 {
	// 	err := errors.New("Could not find Task at given number")
	// }

	// Append each side of the index
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *TaskManager) Help() string {
	/*
		Returns a help menu for the user with instructions for each command

		Returns:
		    a string help menu
	*/

	return HELP_MENU
}

func (t TaskManager) find(number int) int {
	/*
		Finds the index of a task number in the Tasks slice

		Arguments:
		    number (int): the task number we are searching for
		Returns:
		     the index number
	*/

	// Iterate through each task in the slice
	for i, _ := range t.tasks {

		// Check if the current task number is the one we are looking for
		if t.tasks[i].GetNumber() == number {

			// If it is, return the index
			return i // Returns the next index because that is what is displayed in the list
		}
	}

	// Could not find the task, so return an invalid index
	return -1
}
