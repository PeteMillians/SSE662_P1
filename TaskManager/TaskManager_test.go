package taskManager

import (
	"SSE662_P1/Task"
	"fmt"
	"math/rand"
	"testing"
)

func createTasks(number int) []Task.Task {

	var tasks []Task.Task

	for i := 0; i < number; i++ {
		tasks = append(tasks,
			Task.Task{
				Name:        "Test Name",
				Description: "Test Description",
				Status:      false,
				Number:      i + 1,
			},
		)
	}

	return tasks
}

func createTaskManager() TaskManager {
	return TaskManager{
		tasks:  createTasks(5),
		number: 5,
	}
}

func TestGetSize(t *testing.T) {
	var manager TaskManager = createTaskManager()

	if manager.GetSize() != 5 {
		t.Errorf("Expected GetSize - 5")
	}
}

func TestAdd(t *testing.T) {
	var manager TaskManager = createTaskManager()

	manager.Add(createTasks(1)[0])

	if manager.GetSize() != 6 {
		t.Errorf("Expected Size to be 6")
	}

}

func TestList(t *testing.T) {
	var manager TaskManager = createTaskManager()

	if len(manager.List()) != 5 {
		t.Errorf("Expect 5 items in List")
	}

	for i, str := range manager.List() {
		if str != fmt.Sprintf("Task %d: Test Description -- Status: Pending...", i+1) {
			t.Errorf("Incorrect List for task %d", i)
		}
	}
}

func TestMark(t *testing.T) {
	var manager TaskManager = createTaskManager()

	index := rand.Intn(5)

	manager.Mark(index)

	for i, task := range manager.tasks {
		if i == index-1 && !task.Status {
			t.Errorf("Epected Status to be Complete")
		}

		if i != index-1 && task.Status {
			t.Errorf("Expected Status to be Pending")
		}
	}
}

func TestDelete(t *testing.T) {
	var manager TaskManager = createTaskManager()

	index := rand.Intn(5)

	manager.Delete(index)

	if manager.GetSize() != 4 {
		t.Errorf("Expected size to be 4")
	}

	for _, task := range manager.tasks {
		if task.Number == index {
			t.Errorf("Expected this to be deleted")
		}
	}

}

func TestFind(t *testing.T) {
	var manager TaskManager = createTaskManager()

	number := rand.Intn(5) + 1

	index := manager.Find(number)

	for i, task := range manager.tasks {
		if i != index && number == task.Number {
			t.Errorf("Did not expect a number match here")
		}

		if i == index && number != task.Number {
			t.Errorf("Expected match here")
		}
	}
}

func TestUpdate(t *testing.T) {
	var manager TaskManager = createTaskManager()

	number := rand.Intn(4) + 1 // Delete from somewhere that is not the last index

	index := manager.Find(number)

	manager.Delete(index)

	if manager.Find(number) == -1 {
		t.Errorf("Expected to find a task here still")
	}
}
