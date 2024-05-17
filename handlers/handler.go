package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/xhermitx/gotasks/models"
	"github.com/xhermitx/gotasks/store"
)

type TaskHandler struct {
	store store.Store
}

func NewTaskHandler(s store.Store) *TaskHandler {
	return &TaskHandler{
		store: s,
	}
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var body models.Task
	reqBody, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(reqBody, &body)

	fmt.Fprintf(w, "%+v", body)
}

func (t *TaskHandler) ViewTasks(w http.ResponseWriter, r *http.Request) {
	var result []models.Task
	result, err := t.store.ViewTasks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%+v", result)
}
