package main

import (
	"fmt"
	"net/http"
)

var tasks = []string{"Buy groceries", "Complete Go tutorial", "Call Alex"}

func main() {
	http.HandleFunc("/tasks", showTasks)
	http.HandleFunc("/add-task", addTask)
	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	if len(tasks) == 0 {
		fmt.Fprintln(writer, "There are no tasks.")
		return
	}

	fmt.Fprintln(writer, "My tasks:")
	fmt.Fprintln(writer)

	for index, task := range tasks {
    	fmt.Fprintf(writer, "%d. %s\n", index+1, task)
	}
}

func addTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		var task = request.FormValue("task")
		tasks = append(tasks, task)
		fmt.Fprintln(writer, "Task added.")
	} else {
		fmt.Fprintln(writer, "Use POST method to add a task.")
	}
}
