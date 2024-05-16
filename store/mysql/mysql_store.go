package crud

import (
	"fmt"
	"log"

	models "github.com/xhermitx/gotasks/models"
	"gorm.io/gorm"
)


func AddTask(db *gorm.DB, task models.Task) error {

	result := db.Create(&task)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return result.Error // Return the error to the caller
	}

	log.Print("Task ID: ", task.Tid)
	log.Print("Rows Affected: ", result.RowsAffected)

	return nil
}

func DeleteTask(db *gorm.DB, taskID int) error {

	result := db.Delete(&models.Task{}, taskID)
	if result.Error != nil {
		log.Println("An Error occured while deleting the Task")
		return result.Error
	}

	log.Println("Number of Rows Affected : ", result.RowsAffected)

	return nil
}

func UpdateTask(db *gorm.DB, task models.Task) error {

	result := db.Save(&task)
	if result.Error != nil {
		log.Println("Error updating the task")
		return result.Error
	}

	log.Println("Rows Affected : ", result.RowsAffected)
	return nil
}

func ViewTasks(db *gorm.DB) (models.Task, error) {

	var tasks []models.Task

	result := db.Find(&tasks)
	if result.Error != nil {
		log.Print("Error retrieving Data")
		return models.Task{},result.Error
	}

	log.Println("Rows retrieved: ")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Name: %s, Date: %s, Status: %s\n", task.Tid, task.TaskName, task.Date, task.Status)
	}

	return models.Task{},nil
}