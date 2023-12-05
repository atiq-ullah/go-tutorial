package main

import "fmt"

func main() {
	tasks := []Task{}

	tasks = append(tasks, Task{ title: "Rule the world", done: false })

	for i := range tasks {
		tasks[i].complete()
	}

	newTask := createTask("New Task", "", false);
	tasks = append(tasks, newTask)

	printAllTasks(tasks)
	incomplete := getTasksNotDone(tasks)
	fmt.Println("Here are the incomplete tasks: ")
	printAllTasks((incomplete))
}
