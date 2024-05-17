package mysql

import (
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

	result := m.db.Create(task) // task is already a pointer
	if result.Error != nil {
		log.Printf("Error creating task: %v", result.Error)
		return result.Error // Return the error to the caller
	}

	// Log the task ID and the number of rows affected by the Create operation
	log.Printf("Task created successfully with ID: %d, Rows Affected: %d", task.ID, result.RowsAffected)

	return nil
}

func (m *MySQLStore) DeleteTask(taskID int) error {

	// Use a where clause to specify the ID of the task to be deleted
	result := m.db.Where("id = ?", taskID).Delete(&models.Task{})
	if result.Error != nil {
		log.Printf("An error occurred while deleting the task with ID %d: %v", taskID, result.Error)
		return result.Error
	}

	// Log the number of rows affected by the Delete operation
	log.Printf("Task with ID %d deleted successfully, Rows Affected: %d", taskID, result.RowsAffected)

	return nil
}

func (m *MySQLStore) UpdateTask(task *models.Task) error {

	result := m.db.Model(&models.Task{}).Where("id = ?", task.ID).Updates(task)

	if result.Error != nil {
		log.Printf("Error updating task with ID %d: %v", task.ID, result.Error)
		return result.Error
	}

	log.Printf("Task with ID %d updated successfully, Rows Affected: %d", task.ID, result.RowsAffected)
	return nil
}

func (m *MySQLStore) ViewTasks() ([]models.Task, error) {

	var tasks []models.Task

	// Retrieve all tasks from the database
	result := m.db.Find(&tasks)
	if result.Error != nil {
		log.Printf("Error retrieving tasks: %v", result.Error)
		return nil, result.Error
	}

	// Log the number of tasks retrieved
	log.Printf("Number of tasks retrieved: %d", len(tasks))

	return tasks, nil
}
