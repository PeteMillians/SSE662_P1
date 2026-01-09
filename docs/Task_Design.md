# Task Design

This document outlines the design of the Task struct, which contains the information stored in a user-defined task. 

## Requirements
- Contains the following instance variables:
    - Name
    - Description
    - Status
    - Number
- Contains getter and setter methods to interface with an instance of a Task

## Public Methods
```Go
type Task struct {
    Name string
    Description string
    Status bool
    Number int
}
```

## Private Methods
*func (t Task) getName() string {}*

    Returns the Name of the Task object

    Returns:
        the string Name

*func (t *Task) setName(Name string) {}**

    Updates the Name of a Task to a new string

    Arguments:
        Name (string): the Name we are setting as the new Task Name

*func (t Task) getDescription() string {}*

    Returns the Description of the Task object

    Returns:
        the string Description

*func (t *Task) setDescription(Description string) {}**

    Updates the Description of a Task to a new string

    Arguments:
        Description (string): the Description we are setting as the new Task Description

*func (t Task) getNumber() int {}*

    Returns the Number of the Task object

    Returns:
        the int Number

*func (t *Task) setNumber(Number int) {}**

    Updates the Number of a Task to a new int

    Arguments:
        Number (int): the Number we are setting as the new Task Number

*func (t Task) getStatus() bool {}*

    Returns the Status of the Task object

    Returns:
        the bool Status

*func (t *Task) setStatus(Status bool) {}**

    Updates the Status of a Task to True

    Arguments:
        Status (bool): the status we are setting to