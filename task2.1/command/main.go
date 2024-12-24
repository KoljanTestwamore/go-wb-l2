package main

import "fmt"

type ICommand interface {
	Execute()
}

type Reciever struct {}
func (r *Reciever) Do(input string) {
	fmt.Println(input)
}

// invoker
type Switch struct {
	Commands map[string]ICommand
}

func (s *Switch) Register(id string, command ICommand) {
	s.Commands[id] = command
}

func (s *Switch) Execute(id string) {
	for _, command := range s.Commands {
		command.Execute()
	}
}

// concrete command 1
type TurnOnCommand struct {
	r *Reciever
}
func (t *TurnOnCommand) Execute() {
	t.r.Do("Turn on")
}

// concrete command 2
type TurnOffCommand struct {
	r *Reciever
}
func (t *TurnOffCommand) Execute() {
	t.r.Do("Turn off")
}

func main() {
	r := Reciever{}

	turnOnCommand := TurnOnCommand{
		r: &r,
	}
	turnOffCommand := TurnOffCommand{
		r: &r,
	}

	sw := Switch{}

	sw.Register("on", &turnOnCommand)
	sw.Register("off", &turnOffCommand)

	sw.Execute("on")
	sw.Execute("off")
}