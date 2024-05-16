package store

import (
	models "github.com/xhermitx/gotasks/models"
)

type Store interface {
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	DeleteTask(taskID int) error
	ViewTasks() error
}