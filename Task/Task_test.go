package Task

import (
	"testing"
)

func createTestTask() Task {
	task := Task{
		Name:        "Test Name",
		Description: "Test Description",
		Status:      false,
		Number:      1,
	}

	return task
}

func TestCreateTask(t *testing.T) {

	var task Task = createTestTask()

	if task.Name != "Test Name" {
		t.Errorf("expected Name 'Test Name', got %q", task.Name)
	}
	if task.Description != "Test Description" {
		t.Errorf("expected Description 'Test Description', got %q", task.Description)
	}
	if task.Status != false {
		t.Errorf("expected Status false, got %v", task.Status)
	}
	if task.Number != 1 {
		t.Errorf("expected Number 1, got %d", task.Number)
	}
}

func TestString(t *testing.T) {
	var task Task = createTestTask()

	if task.String() != "Task 1 - Test Name: Test Description -- Status: Pending..." {
		t.Errorf("Incorrect String() method")
	}

}

func TestGetName(t *testing.T) {
	var task Task = createTestTask()

	if task.GetName() != "Test Name" {
		t.Errorf("GetName expected Test Name")
	}

}

func TestSetName(t *testing.T) {
	var task Task = createTestTask()

	task.SetName("New Name")

	if task.GetName() != "New Name" {
		t.Errorf("SetName did not update Name")
	}
}

func TestGetDescription(t *testing.T) {
	var task Task = createTestTask()

	if task.GetDescription() != "Test Description" {
		t.Errorf("GetDescription expected Test Description")
	}

}

func TestSetDescription(t *testing.T) {
	var task Task = createTestTask()

	task.SetDescription("New Description")

	if task.GetDescription() != "New Description" {
		t.Errorf("SetDescription did not update Description")
	}
}

func TestGetNumber(t *testing.T) {
	var task Task = createTestTask()

	if task.GetNumber() != 1 {
		t.Errorf("GetNumber expected Test Number")
	}

}

func TestSetNumber(t *testing.T) {
	var task Task = createTestTask()

	task.SetNumber(5)

	if task.GetNumber() != 5 {
		t.Errorf("SetNumber did not update Number")
	}
}

func TestGetStatus(t *testing.T) {
	var task Task = createTestTask()

	if task.GetStatus() != false {
		t.Errorf("GetStatus expected Test Status")
	}

}

func TestSetStatus(t *testing.T) {
	var task Task = createTestTask()

	task.SetStatus(true)

	if task.GetStatus() != true {
		t.Errorf("SetStatus did not update Status")
	}
}
