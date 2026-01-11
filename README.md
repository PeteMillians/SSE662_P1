# To-Do List Manager

This project, module 1 of SSE 662, contains a simple command-line to-do list manager implemented in Go. The To-Do List Manager allows the user to create and delete Tasks, which are defined by the user with a Name and Description.


## Getting Started
To get started with using the To-Do List Manager, make sure you have Go installed and in the PATH. A link to install Go can be found [here](https://go.dev/doc/install).

Clone the repository containing the To-Do List Manager by *cloning by URL* using the following [link](https://github.com/PeteMillians/SSE662_P1.git).

Once Go is installed and set in the PATH, open whatever file editor you prefer, and open a terminal. Ensure the terminal location is set to the top-level directory: 'SSE662_P1'. 

```
C:/Users/myname/.../SSE662_P1>
```

Now that the terminal is at the correct location, simply run the following command to initialize the Go module:

```
go mod init SSE662_P1
```

A file named **go.mod** should have been created, which gives the Go interpreter the module information. Now, to run the To-Do List Manager, run the following command:

```
go run .
```

## Usage Instructions
The To-Do List Manager operates as a command-line interface (CLI), completely controlled by the user. Available commands can be displayed by inputting typing **H** and pressing **enter**. All available commands are case-sensitive.

The following commands are available to the user:
- **A** -  Add a task
- **L** -  List all current tasks
- **M** -  Mark a task as complete
- **D** -  Delete a task
- **H** -  Display the help menu

#### Adding a Task
If the user inputs **A**, the terminal will clear and display a Task creation screen prompting the **Name** and **Description** of the created task. The script will automatically assign the task number and assume it is incomplete. 

#### Listing Tasks
If the user inputs **L**, the terminal will clear and display a list of all current Tasks, pending or completed. *Note: tasks will not be cleared from the list until the user deletes it.* The lst will contain each task's **Name**, **Number**, **Description**, and **Status**. 

#### Marking a Task
If the user inputs **M**, the terminal will clear before displaying the current list of tasks. By displaying the available tasks, the user can immediately see the task number of a desired task to mark. The user will be prompted to input a task number representing which task's status will be set to **Complete**.

#### Deleting a Task
If the user inputs **D**, the terminal will clear before displaying the current list of tasks. By displaying the available tasks, the user can immediately see the task number of a desired task to delete. The user will be prompted to input a task number representing which task to delete. After a task is deleted, all task numbers are automatically updated to sequential 1-N without skipping numbers.