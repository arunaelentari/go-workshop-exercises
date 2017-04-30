package main

import "fmt"

// STRUCT GOES HERE
type player struct {
	nickname string
}

func (p player) Warcry() {
	return "I am the fiercest warrior, hahahaha!"
}

func main() {
	p := player{
		nickname: "Anonymous",
	}

	fmt.Println("Fearlessly,", p.nickname, "dived into the abyss, screaming:", p.warcry())
}
