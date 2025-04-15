package main

import (
	"fmt"
)

const (
	Create = "create"
	Read   = "read"
	Update = "update"
	Delete = "delete"
	Cancel = "!cancel"
)

type TaskData struct {
	taskName        string
	taskDescription string
}

var tasks = make([]TaskData, 0, 3)

func CreateTask(taskName string, taskDescription string) {
	if checkTaskExists(taskName) {
		fmt.Print("Error creating task, reason: task already exists")
		return
	}
	tasks = append(tasks, TaskData{taskName, taskDescription})
	fmt.Printf("Created task name is:  %s \n", taskName)
}

func checkTaskExists(taskName string) bool {
	for i, _ := range tasks {
		if tasks[i].taskName == taskName {
			return true
		}
	}
	return false
}

func UpdateTaskName(newTaskName, oldTaskName string) {
	for i, _ := range tasks {
		if tasks[i].taskName == oldTaskName {
			tasks[i].taskName = newTaskName
			fmt.Printf("New task name is:  %s \n", newTaskName)
			return
		}
	}
	fmt.Print("Error updating task name, reason: task does not exist")
}

func DeleteTask(taskName string) bool {
	for i, task := range tasks {
		if task.taskName == taskName {
			before := tasks[:i]
			after := tasks[i+1:]
			tasks = append(before, after...)
			fmt.Println("Task deleted")
			return true
		}
	}
	fmt.Println("Task does not exist")
	return false
}

func PrintTaskNamesMessage() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	fmt.Print("List of tasks: \n")
	for i, task := range tasks {
		fmt.Printf("[%v]| %s: |\n	{--%s--}\n", i+1, task.taskName, task.taskDescription)
	}
}

func checkCancel(command *string) bool {
	if *command == Cancel {
		*command = ""
		return true
	}
	return false
}

func printCancelMessage() {
	fmt.Println("You can cancel operation if type: !cancel")
}

func printErrorMessage() {
	fmt.Println("Error read input")
	printCancelMessage()
}

func checkInputLen(input string) bool {
	if len(input) < 3 {
		fmt.Println("Please enter name more than 3 symbols")
		return true
	}
	return false
}

func main() {
	var err error
	input := ""
	for {
		if len(input) == 0 {
			fmt.Printf("Enter your command (%s, %s, %s, %s): \n", Create, Read, Update, Delete)
			_, err = fmt.Scan(&input)
			if err != nil {
				printErrorMessage()
				continue
			}
		}
		switch input {
		case Create:
			fmt.Println("Enter task name")
			printCancelMessage()
			_, err = fmt.Scan(&input)
			if err != nil {
				printErrorMessage()
				return
			}
			if checkCancel(&input) {
				break
			}
			if checkInputLen(input) {
				input = Create
				break
			}
			taskName := input
			fmt.Println("Enter task description")
			printCancelMessage()
			_, err = fmt.Scan(&input)
			if err != nil {
				printErrorMessage()
				return
			}
			if checkCancel(&input) {
				break
			}
			CreateTask(taskName, input)
			input = ""
		case Read:
			PrintTaskNamesMessage()
			input = ""
		case Update:
			fmt.Println("Please enter task name which you want to update")
			printCancelMessage()
			_, err = fmt.Scan(&input)
			if err != nil {
				printErrorMessage()
				return
			}
			if checkCancel(&input) {
				break
			}
			if checkInputLen(input) {
				input = Update
				break
			}
			if !checkTaskExists(input) {
				fmt.Println("Task does not exist")
				input = Update
				break
			}
			oldTaskName := input
			fmt.Println("Enter new task name")
			printCancelMessage()
			_, err = fmt.Scan(&input)
			if err != nil {
				printErrorMessage()
				return
			}
			if checkCancel(&input) {
				break
			}
			if checkInputLen(input) {
				input = Update
				break
			}
			if len(tasks) > 1 && checkTaskExists(input) {
				fmt.Println("Can't update task name, new name already exists")
				input = Update
				break
			}
			UpdateTaskName(input, oldTaskName)
			input = ""
		case Delete:
			fmt.Println("Enter delete task name: ")
			printCancelMessage()
			_, err = fmt.Scan(&input)
			if err != nil {
				printErrorMessage()
				return
			}
			if checkCancel(&input) {
				break
			}
			DeleteTask(input)
			input = ""
		case Cancel:
			input = ""
			break
		default:
			printErrorMessage()
			input = ""
		}
	}
}
