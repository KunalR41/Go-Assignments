package main

import (
	"fmt"
	"time"
)

// State Design pattern
type State interface {
	Transition(traficLight *TraficLight)
}

type RedState struct{}

func (rs *RedState) Transition(traficlight *TraficLight) {
	fmt.Println("Trafic light is now RED")
	time.Sleep(3 * time.Second)
	traficlight.currentState = &GreenState{}
}

type YellowState struct{}

func (ys *YellowState) Transition(traficlight *TraficLight) {
	fmt.Println("Trafic light is now YELLOW")
	time.Sleep(1 * time.Second)
	traficlight.currentState = &RedState{}
}

type GreenState struct{}

func (gs *GreenState) Transition(traficlight *TraficLight) {
	fmt.Println("Trafic light is now GREEN")
	time.Sleep(2 * time.Second)
	traficlight.currentState = &YellowState{}
}

type TraficLight struct {
	currentState State
}

func NewTraficLight() *TraficLight {
	return &TraficLight{
		currentState: &RedState{},
	}
}

func (tl *TraficLight) ChangeState() {
	tl.currentState.Transition(tl)
}

func main() {
	traficlight := NewTraficLight()

	for i := 0; i < 5; i++ {
		traficlight.ChangeState()
	}
}
