package main

import "fmt"

type Task struct {
	title string;
	description string;
	done bool;
}

func (t *Task) complete() {
	t.done = true
}

func (t Task) isComplete() bool {
	return t.done;
}

func createTask(title string, description string, done bool) Task {
	return Task{title, description, done}
}

func (t Task) print() {
	fmt.Println("Title:", t.title)
	fmt.Println("Description: ", t.description)
	fmt.Println("Done:", t.done)
}

func printAllTasks(tasks []Task) {
	for _, task := range tasks {
		task.print()
		fmt.Println()
	}
}

func getTasksNotDone(tasks []Task) []Task {
	incomplete := []Task{}
	for _, t := range tasks {
		if !t.isComplete() {
			incomplete = append(incomplete, t)
		}
	}
	return incomplete
}
