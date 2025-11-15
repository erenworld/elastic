package node

// A node is an object that represents any machine in our cluster. 
// For example, the manager is one type of node in Cube. 

type Node struct {
	Name			string
	Ip				string
	Cores   		int
	Memory  		int
	MemoryAllocated int
	Disk 			int 
	DiskAllocated 	int
	Role			string
	TaskCount 		int
}
