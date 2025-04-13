package main

import "fmt"

const (
	E_Create = "create"
	E_Read   = "read"
	E_Update = "update"
	E_Delete = "delete"
)

type TaskData struct {
	TaskName        string
	TaskDescription string
}

var Tasks = make([]TaskData, 0, 3)

func PrintBaseMessage() {
	fmt.Printf("Enter your command (%s, %s, %s, %s): \n", E_Create, E_Read, E_Update, E_Delete)
}

func CreateTask(TaskName string, TaskDescription string) {
	Tasks = append(Tasks, TaskData{TaskName, TaskDescription})
	fmt.Printf("Created task name is:  %s \n", TaskName)
}

func PrintEnterTaskName() {
	fmt.Println("Enter task name")
}

func PrintEnterTaskDescription() {
	fmt.Println("Enter task description")
}

func PrintTaskNames() {
	fmt.Print("List of tasks: \n")
	for _, task := range Tasks {
		fmt.Printf("| %s: |\n--%s", task.TaskName, task.TaskDescription)
	}
}

func main() {
	input := ""
	PrintBaseMessage()
	for {
		fmt.Scan(&input)
		switch input {
		case E_Create:
			PrintEnterTaskName()
			fmt.Scan(&input)
			TaskName := input
			PrintEnterTaskDescription()
			fmt.Scan(&input)
			CreateTask(TaskName, input)
		case E_Read:
			PrintTaskNames()
		}
		input = ""
	}

}
