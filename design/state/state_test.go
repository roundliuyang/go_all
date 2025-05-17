package state

import (
	"fmt"
	"testing"
)

// 音乐播放器
// 想像你正在开发一个简单的音乐播放器，它有三种状态：播放、暫停和停止。每当用户按下按钮，播放器的状态就会改变。

// State 接口定义了处理状态的方法。
type State interface {
	Handle(context *MusicPlayerContext)
}

// PlayState 表示播放状态。
type PlayState struct{}

func (p *PlayState) Handle(context *MusicPlayerContext) {
	fmt.Println("正在播放音乐")
	context.SetState(&PauseState{})
}

// PauseState 表示暂停状态。
type PauseState struct{}

func (p *PauseState) Handle(context *MusicPlayerContext) {
	fmt.Println("音乐已暂停")
	context.SetState(&StopState{})
}

// StopState 表示停止状态。
type StopState struct{}

func (s *StopState) Handle(context *MusicPlayerContext) {
	fmt.Println("音乐已停止")
	context.SetState(&PlayState{})
}

// MusicPlayerContext 包含当前的状态。
type MusicPlayerContext struct {
	state State
}

func (m *MusicPlayerContext) SetState(state State) {
	m.state = state
}

func (m *MusicPlayerContext) PressPlayButton() {
	m.state.Handle(m)
}

func TestMusic(t *testing.T) {
	player := &MusicPlayerContext{state: &StopState{}}
	player.PressPlayButton() // 音乐已停止
	player.PressPlayButton() // 正在播放音乐
	player.PressPlayButton() // 音乐已暂停
}

// 如果你正在开发一个游戏，角色可能会有：正常、受傷和死亡等状态。在受傷状态下，角色的移动速度可能会减慢；在死亡状态下，角色可能不能移动。通过使用状态模式，你可以轻松管理角色在不同状态下的行为。
// CharacterState 是一個介面，定義了不同角色狀態的 Move 方法。
type CharacterState interface {
	Move()
}

// HealthyState 表示角色健康狀態。
type HealthyState struct{}

func (h *HealthyState) Move() {
	fmt.Println("角色快速移動！") // 角色快速移動！
}

// InjuredState 表示角色受傷狀態。
type InjuredState struct{}

func (i *InjuredState) Move() {
	fmt.Println("角色移動得有點慢...") // 角色移動得有點慢...
}

// DeadState 表示角色死亡狀態。
type DeadState struct{}

func (d *DeadState) Move() {
	fmt.Println("角色不能移動，他已經...RIP。") // 角色不能移動，他已經...RIP。
}

// GameCharacter 表示帶有特定狀態的遊戲角色。
type GameCharacter struct {
	state CharacterState
}

func (g *GameCharacter) SetState(state CharacterState) {
	g.state = state
}

func (g *GameCharacter) Move() {
	g.state.Move()
}

func TestGame(t *testing.T) {
	hero := &GameCharacter{state: &HealthyState{}}
	hero.Move() // 角色快速移動！

	hero.SetState(&InjuredState{})
	hero.Move() // 角色移動得有點慢...

	hero.SetState(&DeadState{})
	hero.Move() // 角色不能移動，他已經...RIP。
}
