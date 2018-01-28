package main

type World struct {
	width  int
	height int
	board  string

	plants  []*Plant
	zombies []*Zombie
}

func (world *World) init() {
	world.width = 50
	world.height = 25

	world.initBoard()
}

func (world *World) initBoard() {
	world.board = ""

	for i := 0; i < world.height; i++ {
		for j := 0; j < world.width; j++ {
			if i == 0 || i == (world.height-1) || j == 0 || j == (world.width-1) {
				world.board += "#"
			} else {
				world.board += " "
			}
		}
	}
}
