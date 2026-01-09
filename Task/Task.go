package Task

import "fmt"

type Task struct {
	Name        string
	Description string
	Status      bool
	Number      int
}

func (t Task) String() string {

	// Initialize the status string
	var status string

	if t.Status {
		status = "Complete"
	} else {
		status = "Pending..."
	}

	// Return the formatted Task string
	return fmt.Sprintf("Task %d: %s -- Status: %s", t.Number, t.Description, status)
}

func (t Task) GetName() string {

	// Returns the Name of the Task object

	// Returns:
	//     the string Name

	return t.Name
}

func (t *Task) SetName(Name string) {

	// Updates the Name of a Task to a new string

	// Arguments:
	//     Name (string): the Name we are Setting as the new Task Name

	t.Name = Name
}

func (t Task) GetDescription() string {

	// Returns the Description of the Task object

	// Returns:
	//     the string Description

	return t.Description
}

func (t *Task) SetDescription(Description string) {

	// Updates the Description of a Task to a new string

	// Arguments:
	//     Description (string): the Description we are Setting as the new Task Description

	t.Description = Description
}

func (t Task) GetNumber() int {

	// Returns the Number of the Task object

	// Returns:
	//     the int Number

	return t.Number
}

func (t *Task) SetNumber(Number int) {

	// Updates the Number of a Task to a new int

	// Arguments:
	//     Number (int): the Number we are Setting as the new Task Number

	t.Number = Number
}

func (t Task) GetStatus() bool {

	// Returns the Status of the Task object

	// Returns:
	//     the bool Status

	return t.Status

}

func (t *Task) SetStatus(Status bool) {

	// Updates the Status of a Task to True

	// Arguments:
	// 	Status (bool): the status we are Setting it to

	t.Status = Status
}
