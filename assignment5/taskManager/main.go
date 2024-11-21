package main

import (
	"fmt"
	"taskManager.com/Tasks"
	"time"
)

func main() {
	t := Tasks.CreateTask("Dinner", "Foodadas", time.Now(), false)
	fmt.Println(t)
	t.MarkAsCompleted()
	t.UpdateDescription("Nice food")
	fmt.Println(t)
	t2 := Tasks.CreateTask("Lunch", "Food", time.Now(), false)
	tasks := []Tasks.Task{*t, *t2}
	filteredTasks := Tasks.FilterTasks(tasks, Tasks.IsCompleted)
	fmt.Println(filteredTasks)
	isCompletedAnonymous := func(t Tasks.Task) bool {
		return Tasks.IsCompleted(t)
	}
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i], isCompletedAnonymous(tasks[i]))
	}
}
