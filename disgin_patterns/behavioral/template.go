package behavioral

import "fmt"

type Game interface {
	Play()
}

type template interface {
	initialize()
	startPlay()
	endPlay()
}

type player struct {
	name string
	template
}

func newPlayer(name string, template template) player {
	return player{
		name:     name,
		template: template,
	}
}

func (p *player) Play() {
	p.template.initialize()
	p.template.startPlay()
	p.template.endPlay()
}

func (p *player) initialize() {
	fmt.Printf("%s default initialize\n", p.name)
}

type basketball struct {
	player
	category string
}

func NewBasketballGame(name string) Game {
	b := &basketball{
		category: "basketball",
	}
	b.player = newPlayer(name, b)
	return b
}

func (b *basketball) initialize() {
	fmt.Printf("%s is initialize when doing %s\n", b.name, b.category)
}

func (b *basketball) startPlay() {
	fmt.Printf("%s start Play %s\n", b.name, b.category)

}

func (b *basketball) endPlay() {
	fmt.Printf("%s end Play %s\n", b.name, b.category)
}

type football struct {
	player
	category string
}

func NewFootballGame(name string) Game {
	b := &football{
		category: "football",
	}
	b.player = newPlayer(name, b)
	return b
}

func (b *football) startPlay() {
	fmt.Printf("%s start Play %s\n", b.name, b.category)

}

func (b *football) endPlay() {
	fmt.Printf("%s end Play %s\n", b.name, b.category)
}
