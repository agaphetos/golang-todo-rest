package main

// GetTasks returns all Tasks
func GetTasks() []Task {
	db := InitEngine()
	defer db.Close()
	tasks := []Task{}
	db.Find(&tasks)
	return tasks
}

// GetTask returns a Task record based from given ID
func GetTask(id int) (*Task, error) {
	db := InitEngine()
	defer db.Close()
	task := Task{ID: id}
	if _, err := db.Get(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

// CreateTask creates a Task record to the database
func CreateTask(task *Task) (int64, error) {
	db := InitEngine()
	defer db.Close()
	affected, err := db.Insert(task)
	if err != nil {
		return 0, err
	}
	return affected, nil
}

// UpdateTask updates a Task record
func UpdateTask(id int, task *Task) (int64, error) {
	db := InitEngine()
	defer db.Close()
	affected, err := db.Id(id).Update(task)
	if err != nil {
		return 0, err
	}
	return affected, nil
}

// DeleteTask deletes a Task record from a given ID
func DeleteTask(id int) (int64, error) {
	db := InitEngine()
	defer db.Close()
	affected, err := db.Id(id).Delete(&Task{})
	if err != nil {
		return 0, err
	}
	return affected, nil
}
