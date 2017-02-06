package gg

import (
	"fmt"
)

// Pony
type Pony interface {
	Walk()
	Run()
}

// EarthPony
type EarthPony struct {
	Pony
}

// SpacePony
type SpacePony interface {
	Pony
	Fly()
}

// Pegasus
type Pegasus struct {
	SpacePony
	name string
}

// Walk EarthPony
func (ep *EarthPony) Walk() {
	fmt.Println("歩くよ")
}

// Run EarthPony
func (ep *EarthPony) Run() {
	fmt.Println("ばたばた")
}

// Walk Pegasus
func (p *Pegasus) Walk() {
	fmt.Println("あるーくよ")
}

// Run Pegasus
func (p *Pegasus) Run() {
	fmt.Println("はしーるよ")
}

// Fly Pegasus
func (p *Pegasus) Fly() {
	fmt.Println("飛ぶよ")
}

// Sprint
func Sprint(p Pony) {
	p.Walk()
	p.Run()
}

func Sprinte(p SpacePony) {
	p.Walk()
	p.Run()
	p.Fly()
}

func main() {
	appleJack := &EarthPony{}
	Sprint(appleJack)

	rainbowDash := &Pegasus{}
	Sprinte(rainbowDash)
	rainbowDash.name = "name"
	fmt.Println(rainbowDash.name)
}
