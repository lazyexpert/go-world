package main

func gameLoop(renderer *Renderer) {
	go renderer.render()

	for {
		// some calculations here
	}
}

func main() {
	world := &World{}
	world.init()
	println(world.board[5])
	renderer := &Renderer{500, world}
	gameLoop(renderer)
}
