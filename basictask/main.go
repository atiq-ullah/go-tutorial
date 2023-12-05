package main

import "fmt"

type Task struct {
	title string;
	done bool;
}

func (t Task) print() {
	fmt.Println("Title:", t.title)
	fmt.Println("Done:", t.done)
}

func (t *Task) complete() {
	fmt.Println("Task", t.title ,"done!")
	t.done = true
}

func main() {
	tasks := []Task{}

	tasks = append(tasks, Task{ title: "Rule the world", done: false })
	fmt.Println(tasks)

	for _, t := range tasks {
		t.complete()
	}

	fmt.Println(tasks)
}
