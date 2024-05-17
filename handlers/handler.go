package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	err := json.Unmarshal(reqBody, &body)
	if err != nil {
		log.Println("Error processing the request")
	}

	if err = t.store.CreateTask(&body); err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "Task Created Successfully")
}

func (t *TaskHandler) ViewTasks(w http.ResponseWriter, r *http.Request) {
	var result []models.Task
	result, err := t.store.ViewTasks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%+v", result)
}

func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &task)
	if err != nil {
		log.Println("Error processing the request")
	}

	if err = t.store.UpdateTask(&task); err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "Task %s successfully updated!", task.TaskName)
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("Please provide a valid id")
	}

	if err = t.store.DeleteTask(id); err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "Task with id %d deleted successfully!", id)
}
