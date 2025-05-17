package main

import "fmt"

// Command 是命令的接口
type Command interface {
	Execute()
}

// 开灯 Command
type LightOnCommand struct {
	Light *Light
}

func (l *LightOnCommand) Execute() {
	l.Light.TurnOn()
}

// 关灯 Command
type LightOffCommand struct {
	Light *Light
}

func (l *LightOffCommand) Execute() {
	l.Light.TurnOff()
}

// 灯
type Light struct{}

// 开灯
func (l *Light) TurnOn() {
	fmt.Println("燈亮了！")
}

// 关灯
func (l *Light) TurnOff() {
	fmt.Println("燈熄了！")
}

// Remote 遙控器
type Remote struct {
	Command Command
}

func (r *Remote) PressButton() {
	r.Command.Execute()
}

func main() {
	light := &Light{}
	lightOn := &LightOnCommand{Light: light}
	lightOff := &LightOffCommand{Light: light}

	remote := &Remote{}

	remote.Command = lightOn
	remote.PressButton()

	remote.Command = lightOff
	remote.PressButton()
}
