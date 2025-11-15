package task

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
) 

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

// A Task that a user wants to run on our cluster (pending, scheduled, running, completed, failed).
type Task struct {
	ID			  uuid.UUID
	ContainerId	  string
	Name 		  string
	State 		  State
	Image 		  string
	Memory  	  int
	Disk 		  int
	ExposedPorts  nat.PortSet
	PortBinding   map[string]string
	RestartPolicy string
	StartTime	  time.Time
	FinishTime	  time.Time
}

// Event struct to tell the system to stop a task.
type TaskEvent struct {
	ID			uuid.UUID
	State		State		// From state one to state two
	Timestamp 	time.Time	// To record the time the event was requested
	Task 		Task
}

type Config struct {
	Name		  string
	AttachStdin	  bool
	AttachStdout  bool
	AttachStderr  bool
	ExposedPorts  nat.PortSet
	Cmd			  []string
	Image		  string
	Cpu 		  float64
	Memory		  int64
	Disk		  int64
	Env			  []string
	RestartPolicy string
}

// Encapsulation to run our task as docker container.
type Docker struct {
	Client *client.Client
	Config Config
}

type DockerResult struct {
	ContainerId	string
	Error 		error
	Action 		string
	Result 		string
}

func (d *Docker) Run() DockerResult {
	ctx := context.Background()
	reader, err := d.Client.ImagePull(ctx, d.Config.Image, image.PullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s: %v\n", d.Config.Image, err)
		return DockerResult{Error: err}
	}
	// Copy the stream of logs from Docker to the terminal
	io.Copy(os.Stdout, reader) 
	return DockerResult{}
}