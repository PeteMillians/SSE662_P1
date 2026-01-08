# Task Design

This document outlines the design of the Task struct, which contains the information stored in a user-defined task. 

## Requirements
- Contains the following instance variables:
    - Name
    - Description
    - Number
- Contains getter and setter methods to interface with an instance of a Task
- Instance variables are lowercase so that they are private

## Public Methods
```Go
type Task struct {
    name string
    description string
    number int
}
```

## Private Methods
*func (t Task) getName() string {}*

    Returns the name of the Task object

    Returns:
        the string name

*func (t Task) setName(name string) {}**

    Updates the name of a Task to a new string

    Arguments:
        name (string): the name we are setting as the new Task name

*func (t Task) getDescription() string {}*

    Returns the description of the Task object

    Returns:
        the string description

*func (t Task) setDescription(description string) {}**

    Updates the description of a Task to a new string

    Arguments:
        description (string): the description we are setting as the new Task description

*func (t Task) getNumber() int {}*

    Returns the number of the Task object

    Returns:
        the int number

*func (t Task) setNumber(number int) {}**

    Updates the number of a Task to a new int

    Arguments:
        number (int): the number we are setting as the new Task number