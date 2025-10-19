package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type TaskStatus string

const (
	StatusToDo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created-at"`
	UpdatedAt   time.Time  `json:"updated-at"`
}

type TaskFile struct {
	Tasks         []Task `json:"tasks"`
	TaskIDCounter int    `json:"task-id-counter"`
}

func loadTasks() TaskFile {
	var taskFile TaskFile

	data := LoadData()
	if len(data) == 0 {
		return TaskFile{}
	}

	err := json.Unmarshal(data, &taskFile)
	if err != nil {
		log.Fatalf("failed to unmarshal json data: %s", err)
	}
	return taskFile
}

func saveTasks(taskFile TaskFile) {
	data, err := json.Marshal(taskFile)
	if err != nil {
		log.Fatalf("failed to marshal json data: %s", err)
	}
	SaveData(data)
}

func GetTaskIndex(taskFile TaskFile, id int) int {
	for index, task := range taskFile.Tasks {
		if task.Id == id {
			return index
		}
	}
	return -1
}

func AddTask(description string) {
	taskFile := loadTasks()

	taskFile.TaskIDCounter++

	task := Task{
		Id:          taskFile.TaskIDCounter,
		Description: description,
		Status:      StatusToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	log.Print(taskFile)
	taskFile.Tasks = append(taskFile.Tasks, task)
	log.Print(taskFile)
	saveTasks(taskFile)
}

func UpdateTask(id int, description string) {
	taskFile := loadTasks()
	taskIndex := GetTaskIndex(taskFile, id)
	if taskIndex < 0 {
		log.Fatal("invalid task index")
	}
	taskFile.Tasks[taskIndex].Description = description
	taskFile.Tasks[taskIndex].UpdatedAt = time.Now()
	saveTasks(taskFile)
}

func ListTasks() {
	tasks := loadTasks().Tasks
	for _, value := range tasks {
		displayTask(value)
	}
}

func ListTasksFilter(filter string) {
	tasks := loadTasks().Tasks
	for _, value := range tasks {
		if value.Status == TaskStatus(filter) {
			displayTask(value)
		}
	}
}

func MarkDone(id int) {

	taskFile := loadTasks()
	taskIndex := GetTaskIndex(taskFile, id)
	if taskIndex < 0 {
		log.Fatal("invalid task index")
	}

	taskFile.Tasks[taskIndex].Status = StatusDone
	taskFile.Tasks[taskIndex].UpdatedAt = time.Now()

	saveTasks(taskFile)
}

func MarkInProgress(id int) {

	taskFile := loadTasks()
	taskIndex := GetTaskIndex(taskFile, id)
	if taskIndex < 0 {
		log.Fatal("invalid task index")
	}

	taskFile.Tasks[taskIndex].Status = StatusInProgress
	taskFile.Tasks[taskIndex].UpdatedAt = time.Now()

	saveTasks(taskFile)
}

func displayTask(task Task) {
	fmt.Printf("Task - %d -\n %s\n", task.Id, task.Description)
	fmt.Printf(" Status: %s\n", task.Status)
	fmt.Println(" Time created: ", task.CreatedAt)
	fmt.Println(" Time updated: ", task.UpdatedAt)
	fmt.Println()
}
