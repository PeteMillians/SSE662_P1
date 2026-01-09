package main

import (
	taskManager "SSE662_P1/TaskManager"
	"SSE662_P1/util"
)

func main() {
	manager := &taskManager.TaskManager{}

	util.CommandLine(manager)
}
