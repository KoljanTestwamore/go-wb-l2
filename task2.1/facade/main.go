package main

type Memory struct {}
func (m *Memory) LoadUp() {}


type HardDrive struct {}
func (h *HardDrive) Read() {}

type ComputerFacade struct {
	m Memory
	hd HardDrive
}

func NewComputerFacade() ComputerFacade {
	return ComputerFacade{
		m: Memory{},
		hd: HardDrive{},
	}
}

func (c *ComputerFacade) Start() {
	c.m.LoadUp()
	c.hd.Read()
}

func main() {
	f := NewComputerFacade()
	f.Start()
}