package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("I am root!")
}

func handleGetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := GetTasks()
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(tasks)
}

func handleGetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if id, err := strconv.Atoi(params["id"]); err == nil {
		task, err := GetTask(id)
		switch {
		case err != nil:
			fmt.Println(err.Error())
		case task != nil:
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func handlePostTask(w http.ResponseWriter, r *http.Request) {
	var result string
	decoder := json.NewDecoder(r.Body)
	var task Task
	err := decoder.Decode(&task)
	if err != nil {
		log.Println(fmt.Errorf("Error reading received data: %v", err))
		enc := json.NewEncoder(w)
		enc.Encode("")
		return
	}
	id := GenerateID()
	task.ID = id
	task.Status = 1
	task.CreatedDate = time.Now()
	status, err := CreateTask(&task)
	switch {
	case err != nil:
		log.Println(fmt.Errorf("Error inserting data: %v", err))
		result = fmt.Sprintf("Error inserting data: %v", err)
	case status > 0:
		result = "Insert Data Success!"
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(result)
}

func handlePutTask(w http.ResponseWriter, r *http.Request) {
	var result string
	decoder := json.NewDecoder(r.Body)
	var task Task
	err := decoder.Decode(&task)
	if err != nil {
		log.Println(fmt.Errorf("Error reading received data: %v", err))
		enc := json.NewEncoder(w)
		enc.Encode("")
		return
	}
	newTask := new(Task)
	newTask.Description = task.Description
	status, err := UpdateTask(task.ID, newTask)
	switch {
	case err != nil:
		log.Println(fmt.Errorf("Error updating data: %v", err))
		result = fmt.Sprintf("Error updating data: %v", err)
	case status > 0:
		result = "Update Data Success!"
	default:
		result = "ID Not Data Found!"
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(result)
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	var result string
	params := mux.Vars(r)
	if id, err := strconv.Atoi(params["id"]); err == nil {
		status, err := DeleteTask(id)
		switch {
		case err != nil:
			log.Println(fmt.Errorf("Error deleting data: %v", err))
			result = fmt.Sprintf("Error deleting data: %v", err)
		case status > 0:
			result = "Delete Data Success!"
		default:
			result = "ID Not Data Found!"
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(result)
	}
}
