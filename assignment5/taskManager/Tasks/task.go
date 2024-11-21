package Tasks

import "time"

type Task struct {
	name        string
	description string
	dueDate     time.Time
	isComplete  bool
}

func CreateTask(name string, description string, dueDate time.Time, isComplete bool) *Task {
	return &Task{
		name:        name,
		description: description,
		dueDate:     dueDate,
		isComplete:  isComplete,
	}
}
func (t *Task) MarkAsCompleted() {
	t.isComplete = true
}
func (t *Task) UpdateDescription(newDescription string) {
	t.description = newDescription
}
func FilterTasks(tasks []Task, condition func(Task) bool) []Task {
	conditionTasks := []Task{}
	for _, task := range tasks {
		if IsCompleted(task) {
			conditionTasks = append(conditionTasks, task)
		}
	}
	return conditionTasks
}
func IsCompleted(task Task) bool {
	return task.isComplete
}
func CountCompletedTasks(tasks ...Task) int {
	count := 0
	for _, task := range tasks {
		if IsCompleted(task) {
			count++
		}
	}
	return count
}
