package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func gameLoop() {
	for {
		fmt.Printf("Game width: %d, game height: %d", worldWidth, worldHeight)
		time.Sleep(500 * time.Millisecond)
		clear()
	}
}

func main() {
	fmt.Println("Starting app")
	gameLoop()
}
