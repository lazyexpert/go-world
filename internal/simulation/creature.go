package main

// Creature - common interface for zombies and plants
type Creature interface {
	init() *Creature
	toString() string
	action()
}
