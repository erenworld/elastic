package node

// A Node is basically a computer that can run tasks.
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