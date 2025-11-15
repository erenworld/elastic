package manager

import (
	"orkestra/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
 	"github.com/google/uuid"
)

type Manager struct {
	Pending			queue.Queue						// Tasks in FIFO order
	TaskDb  		map[string][]*task.Task			// In-memory db to store task
	EventDb 		map[string][]*task.TaskEvent    // In-memory db to store taskEvent
	Workers			[]string						// Keep track of the workers in cluster
	WorkerTaskMap 	map[string][]uuid.UUID			// Task by UUID
	TaskWorkerMap 	map[uuid.UUID]string			// Task by name
}

func (m *Manager) SelectWorker() {
	fmt.Println("Select appropriate worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("Update tasks")
}

func (m *Manager) SendWorker() {
	fmt.Println("Send work to worker")
}
