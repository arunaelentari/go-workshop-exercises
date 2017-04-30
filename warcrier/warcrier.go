package main

import "fmt"

// STRUCTS AND INTERFACE GOES HERE

type player struct {
	nickname string
	phrase   string
}

type enemy struct {
	name   string
	phrase string
}

type warcrier interface {
	warcry() string
}

func (p player) warcry() string {
	return p.nickname + " says: " + p.phrase
}

func (e enemy) warcry() string {
	return e.name + " says: " + e.phrase
}

func battle(c warcrier) {
	fmt.Println(c.warcry())
}

func main() {
	thor := player{
		nickname: "Thor",
		phrase:   "I will defeat you!",
	}
	gorr := enemy{
		name:   "Gorr",
		phrase: "I will drink from your skull!",
	}

	sif := player{
		nickname: "Sif",
		phrase:   "I will strangle you with my golden hair!",
	}

	battle(thor)
	battle(gorr)
	battle(sif)
}
