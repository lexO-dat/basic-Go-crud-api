//to start the import i have to run the comand:
//go mod init [name of the project]
//then i can import the packages that i need from github with go get [package url]

package main

//the fmt package is used to print to the console
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// this is the struct of the task
type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

// this is the array of all the tasks
type allTasks []task

// the tasks var contains the array of all the tasks
var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	//the header of the response will be set to the json format
	w.Header().Set("Content-Type", "application/json")

	//this get will return the tasks in the json format
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	//this post will create a new task
	var newTask task

	//the ioutil package is used to read the body of the request
	reqBody, err := ioutil.ReadAll(r.Body)

	//if the body is not valid it will return an error
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	//the json package is used to unmarshal the body of the request
	json.Unmarshal(reqBody, &newTask)

	//the new task will have the id of the length of the tasks array + 1 (to always have a unique and consecutive id)
	newTask.ID = len(tasks) + 1

	//the new task will be added to the tasks array
	tasks = append(tasks, newTask)

	//the header of the response will be set to the json format
	w.Header().Set("Content-Type", "application/json")

	//the status of the response will be set to 201 (created)
	w.WriteHeader(http.StatusCreated)

	//the new task will be returned in the json format
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	//the header of the response will be set to the json format
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	//the id of the task will be converted to an int
	task_id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//the task will be searched in the tasks array
	for _, task := range tasks {
		if task.ID == task_id {
			//if the task is found it will be returned in the json format
			json.NewEncoder(w).Encode(task)
			return
		} else {
			fmt.Fprintf(w, "Task not found")
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	task_id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//this code search a task in the tasks array and remove it by manteining the ones before and after it
	for i, task := range tasks {
		if task.ID == task_id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "The task with ID %v has been removed successfully", task_id)
		}
	}
}

// in this case the updateTask will delete the task and create a new one with the same id
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	task_id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	var updatedTask task

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range tasks {
		if task.ID == task_id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updatedTask.ID = task_id
			tasks = append(tasks, updatedTask)
			fmt.Fprintf(w, "The task with ID %v has been updated successfully", task_id)
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func main() {
	//all the routes are made by the mux package
	//this is the router
	router := mux.NewRouter().StrictSlash(true)

	//this is the index route
	router.HandleFunc("/", indexRoute)

	//this is the route to get the tasks
	router.HandleFunc("/tasks", getTasks).Methods("GET")

	//this is the route to create a new task
	router.HandleFunc("/tasks", createTask).Methods("POST")

	//this is the route to get a single task
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")

	//this is the route to delete a single task
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	//this is the route to update a single task
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	//here using the http module i started a server on the port 300 and the router
	//the log.Fatal is used to log the error if the server can't start
	log.Fatal(http.ListenAndServe(":3000", router))

}
