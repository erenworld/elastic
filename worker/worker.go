package worker

import (
	"orkestra/task"

	"github.com/google/uuid"
	"github.com/golang-collections/collections/queue"
) 

// Run tasks as Docker containers
// Accept tasks to run from a manager
// Provide relevant statistics to the manager for the purpose of scheduling tasks
// Keep track of its tasks and their state

// TODOS: implement my own queue in Golang
// TODOS: implement my own DB
type Worker struct {
	Name 		string
	Queue   	queue.Queue					// a map of UUIDs to tasks
	Db			map[uuid.UUID]*task.Task   	// Tasks are handled in FIFO order  
	TaskCount 	int
}