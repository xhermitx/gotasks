package models

import (
	"encoding/json"
	"time"
)

type Task struct {
	ID       int       `gorm:"primary_key;AUTO_INCREMENT"` // Specify auto-increment
	TaskName string    `gorm:"not null"`
	Status   string    `gorm:"not null; default:'Pending'"`
	Date     time.Time `gorm:"not null"`
}

// PARSE THE FORMAT "YYYY-MM-DD" AS time.TIME
func (t *Task) UnmarshalJSON(b []byte) error {

	var temp interface{}

	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	body := temp.(map[string]interface{})

	date := body["Date"].(string)

	t.ID = int(body["ID"].(float64))
	t.TaskName = body["TaskName"].(string)
	t.Status = body["Status"].(string)
	t.Date, err = time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}

	return nil
}
