package main

import (
	"fmt"
	"strconv"
	"time"
)

// Task model
type Task struct {
	ID          int       `xorm:"pk autoincr", json:"id"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedDate time.Time `json:"createdDate"`
}

// GenerateID assigns a new ID to Task
func GenerateID() int {
	db := InitEngine()
	defer db.Close()

	sql := "SELECT MAX(id) + 1 AS ID FROM public.task"
	results, err := db.Query(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	id, _ := strconv.Atoi(string(results[0]["id"]))
	return id
}
