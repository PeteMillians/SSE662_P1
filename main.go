package main

import (
	taskManager "SSE662_P1/TaskManager"
	"SSE662_P1/util"
)

func main() {
	/*
		Main executable to run the To-Do List Manager
	*/

	// Instantiate a TaskManager object
	manager := &taskManager.TaskManager{}

	// Run the Command Line utility
	util.CommandLine(manager)
}
