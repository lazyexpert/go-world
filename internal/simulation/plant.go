package main

type Plant struct {
	symbol      string
	level       int
	hp          float32
	damage      float32
	levelUpKoef float32
	currentPos  int
}

func (p *Plant) init() *Plant {
	p.symbol = "p"
	p.level = 0
	p.hp = 100
	p.damage = 1
	p.levelUpKoef = 1.1

	return p
}

func (p *Plant) toString() string {
	return p.symbol
}

func (p *Plant) action(creatures []Creature) {
	// think of possible actions in here
	// will be performe on each tick
}
