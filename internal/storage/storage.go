package storage

import "task_manager_backend/internal/models"

var tasks []models.Task

func GetTasks() []models.Task {
	return tasks
}

func AddTask(task models.Task) {
	tasks = append(tasks, task)
}

func UpdateTask(id int, updatedTask models.Task) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			return
		}
	}
}

func DeleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
}
