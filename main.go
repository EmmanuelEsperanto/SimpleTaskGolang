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
	for i, task := range Tasks {
		fmt.Printf("[%v]| %s: |\n	{--%s--}\n", i+1, task.TaskName, task.TaskDescription)
	}
}

func PrintLittleTaskName() {
	fmt.Println("Please enter name more than 3 symbols")
}

func main() {
	input := ""

	for {

		if len(input) == 0 {
			PrintBaseMessage()
			fmt.Scan(&input)

		}
		switch input {
		case E_Create:
			//BackToStartOfCreate:
			PrintEnterTaskName()
			fmt.Scan(&input)
			if len(input) <= 3 {
				PrintLittleTaskName()
				input = E_Create
				break
				//goto BackToStartOfCreate
			}
			TaskName := input
			PrintEnterTaskDescription()
			fmt.Scan(&input)
			CreateTask(TaskName, input)
			input = ""

		case E_Read:
			PrintTaskNames()
			input = ""
		}

	}

}
