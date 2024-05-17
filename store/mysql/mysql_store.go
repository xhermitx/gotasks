package mysql

import (
	"fmt"
	"log"

	models "github.com/xhermitx/gotasks/models"
	"gorm.io/gorm"
)

type MySQLStore struct {
	db *gorm.DB
}

func NewMySQLStore(db *gorm.DB) *MySQLStore {
	return &MySQLStore{db: db}
}

func (m *MySQLStore) CreateTask(task *models.Task) error {

	result := m.db.Create(&task)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return result.Error // Return the error to the caller
	}

	log.Print("Task ID: ", task.Tid)
	log.Print("Rows Affected: ", result.RowsAffected)

	return nil
}

func (m *MySQLStore) DeleteTask(taskID int) error {

	result := m.db.Delete(&models.Task{}, taskID)
	if result.Error != nil {
		log.Println("An Error occured while deleting the Task")
		return result.Error
	}

	log.Println("Number of Rows Affected : ", result.RowsAffected)

	return nil
}

func (m *MySQLStore) UpdateTask(task *models.Task) error {

	result := m.db.Save(&task)
	if result.Error != nil {
		log.Println("Error updating the task")
		return result.Error
	}

	log.Println("Rows Affected : ", result.RowsAffected)
	return nil
}

func (m *MySQLStore) ViewTasks() ([]models.Task, error) {

	var tasks []models.Task

	result := m.db.Find(&tasks)
	if result.Error != nil {
		log.Print("Error retrieving Data")
		return []models.Task{}, result.Error
	}

	log.Println("Rows retrieved: ")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Name: %s, Date: %s, Status: %s\n", task.Tid, task.TaskName, task.Date, task.Status)
	}

	return tasks, nil
}
