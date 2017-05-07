package main

import (
	"errors"
	"fmt"
)

type player struct {
	stamina int
}

// jump returns an error if the player can't jump or nil if they could.
func (p *player) jump() error {
	if p.stamina <= 0 {
		return errors.New("no stamina left")
	}
	p.stamina = p.stamina - 1
	fmt.Printf("In jump, p has address %p and stamina %v\n", p, p.stamina)
	return nil
}

func main() {
	p := player{
		stamina: 5,
	}

	for i := 0; i < 10; i++ {
		if err := p.jump(); err != nil {
			fmt.Printf("I could't jump: %v\n", err)
			break
		}
		fmt.Printf("In main, p has address %p\n", &p)
		fmt.Printf("I jumped and now I have %v stamina left\n", p.stamina)
	}
}
