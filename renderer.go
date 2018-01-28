package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Renderer struct {
	delay time.Duration
	world *World
}

func (renderer *Renderer) clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (r *Renderer) render() {
	for {
		r.clear()
		fmt.Print("\n")

		for i := 0; i < r.world.height; i++ {
			for j := 0; j < r.world.width; j++ {
				index := i*r.world.width + j
				fmt.Print(string(r.world.board[index]))
			}
			fmt.Print("\n")
		}

		time.Sleep(r.delay * time.Millisecond)
	}
}
