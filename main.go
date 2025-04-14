package main

import (
	"fmt"
)

const (
	E_Create = "create"
	E_Read   = "read"
	E_Update = "update"
	E_Delete = "delete"
	E_Cancel = "!cancel"
)

type TaskData struct {
	TaskName        string
	TaskDescription string
}

var Tasks = make([]TaskData, 0, 3)

func CreateTask(TaskName string, TaskDescription string) {
	if CheckTaskExists(TaskName) {
		fmt.Print("Error creating task, reason: ")
		PrintTaskExists()
		return
	}
	Tasks = append(Tasks, TaskData{TaskName, TaskDescription})
	fmt.Printf("Created task name is:  %s \n", TaskName)
}

func CheckTaskExists(TaskName string) bool {
	for i, _ := range Tasks {
		if Tasks[i].TaskName == TaskName {
			return true
		}
	}
	return false
}

func UpdateTaskName(NewTaskName, OldTaskName string) {
	for i, _ := range Tasks {
		if Tasks[i].TaskName == OldTaskName {
			Tasks[i].TaskName = NewTaskName
			fmt.Printf("New task name is:  %s \n", NewTaskName)
			return
		}
	}
	fmt.Print("Error updating task name, reason: ")
	PrintTaskExists()
}

func DeleteTask(TaskName string) {
	for i, task := range Tasks {
		if task.TaskName == TaskName {
			Tasks[i] = Tasks[len(Tasks)-1]
			Tasks = Tasks[:len(Tasks)-1]
		}
	}
}

func PrintCancelMessage() {
	fmt.Println("You can cancel operation if type: !cancel")
}

func PrintTaskNoExists() {
	fmt.Println("Task does not exist")
}

func PrintTaskExists() {
	fmt.Println("Task exists")
}

func PrintTaskNames() {
	fmt.Print("List of tasks: \n")
	for i, task := range Tasks {
		fmt.Printf("[%v]| %s: |\n	{--%s--}\n", i+1, task.TaskName, task.TaskDescription)
	}
}

func PrintBaseMessage() {
	fmt.Printf("Enter your command (%s, %s, %s, %s): \n", E_Create, E_Read, E_Update, E_Delete)
}

func PrintEnterTaskName() {
	fmt.Println("Enter task name")
	PrintCancelMessage()
}

func PrintEnterNewTaskName() {
	fmt.Println("Enter new task name")
	PrintCancelMessage()
}

func PrintEnterTaskDescription() {
	fmt.Println("Enter task description")
	PrintCancelMessage()
}

func PrintLittleTaskName() {
	fmt.Println("Please enter name more than 3 symbols")
}

func PrintErrorMessage() {
	fmt.Println("Error read input")
	PrintCancelMessage()
}

func PrintSelectTaskToUpdate() {
	fmt.Println("Please enter task name which you want to update")
	PrintCancelMessage()
}

func main() {
	input := ""

	for {

		if len(input) == 0 {
			PrintBaseMessage()
			_, err := fmt.Scan(&input)
			if err != nil {
				PrintErrorMessage()
				continue
			}

		}
		switch input {
		case E_Create:
			//BackToStartOfCreate:
			PrintEnterTaskName()
			_, err1 := fmt.Scan(&input)
			if err1 != nil {
				PrintErrorMessage()
				return
			}
			if input == E_Cancel {
				input = ""
				break
			}
			if len(input) < 3 {
				PrintLittleTaskName()
				input = E_Create
				break
				//goto BackToStartOfCreate
			}
			TaskName := input
			PrintEnterTaskDescription()
			_, err2 := fmt.Scan(&input)
			if err2 != nil {
				PrintErrorMessage()
				return
			}
			if input == E_Cancel {
				input = ""
				break
			}
			CreateTask(TaskName, input)
			input = ""

		case E_Read:
			PrintTaskNames()
			input = ""
		case E_Update:
			PrintSelectTaskToUpdate()
			_, err1 := fmt.Scan(&input)
			if err1 != nil {
				PrintErrorMessage()
				return
			}
			if input == E_Cancel {
				input = ""
				break
			}
			if len(input) < 3 {
				PrintLittleTaskName()
				input = E_Update
				break
			}
			if !CheckTaskExists(input) {
				PrintTaskNoExists()
				input = E_Update
				break
			}
			OldTaskName := input
			PrintEnterNewTaskName()
			_, err2 := fmt.Scan(&input)
			if err2 != nil {
				PrintErrorMessage()
				return
			}
			if input == E_Cancel {
				input = ""
				break
			}
			if len(input) < 3 {
				PrintLittleTaskName()
				input = E_Update
				break
			}
			if len(Tasks) > 1 && CheckTaskExists(input) {
				PrintTaskExists()
				input = E_Update
				break
			}
			UpdateTaskName(input, OldTaskName)
			input = ""
		case E_Delete:

		case E_Cancel:
			input = ""
			break
		default:
			PrintErrorMessage()
			input = ""
		}

	}

}
