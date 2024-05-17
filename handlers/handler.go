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
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		// ERROR HANDLING THE REQUEST BODY
		http.Error(w, "Error reading request body", http.StatusBadRequest) // 400
		return
	}
	defer r.Body.Close()

	if err = json.Unmarshal(reqBody, &body); err != nil {
		// ERROR PROCESSING THE REQUEST BODY
		http.Error(w, "Error processing the request", http.StatusBadRequest) // 400
		return
	}

	if err = t.store.CreateTask(&body); err != nil {
		// INTERNAL SERVER ERROR
		http.Error(w, "Error creating task", http.StatusInternalServerError) // 500
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated) // 201 CREATED

	fmt.Fprintln(w, "Task Created Successfully")
}

func (t *TaskHandler) ViewTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []models.Task

	tasks, err := t.store.ViewTasks()
	if err != nil {
		http.Error(w, "Error Fetching Data", http.StatusInternalServerError) // 500
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Serialize the result to JSON and write to response
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		// If an error occurs during encoding, log it and send a generic error message
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		log.Println(err)
	}
}

func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	reqBody, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	err := json.Unmarshal(reqBody, &task)
	if err != nil {
		http.Error(w, "Error processing the request", http.StatusBadRequest) // 400
		return
	}

	if err = t.store.UpdateTask(&task); err != nil {
		http.Error(w, "Error Updating the task", http.StatusInternalServerError) // 500
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "Task %s updated successfully!", task.TaskName)
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	if err = t.store.DeleteTask(id); err != nil {
		http.Error(w, "Error Deleting the task", http.StatusInternalServerError) // 500
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Task with ID %d deleted successfully!", id)
}
