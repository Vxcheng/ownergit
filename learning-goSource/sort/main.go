package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	log.Printf("start sort\n")
	SortHero()
	fmt.Println("-----------------")
	SortHeroList()
	fmt.Println("-----------------")
	IterString()
}

type HeroKind int

const (
	Up HeroKind = iota
	Middle
	Gank
	Down
	Assassin
)

type Hero struct {
	Name       string
	Profession HeroKind
}
type HeroList []*Hero

func (h *HeroList) SortBy() {
	list := []*Hero(*h)
	sort.Slice(list, func(i, j int) bool {
		if list[i].Profession != list[j].Profession {
			return list[i].Profession < list[j].Profession
		}
		return list[i].Name < list[j].Name
	})
}

func SortHeroList() {
	heros := []*Hero{
		&Hero{"a", Up},
		&Hero{"b", Middle},
		&Hero{"c", Gank},
		&Hero{"d", Down},
		&Hero{"e", Assassin},
		&Hero{"g", Middle},
	}
	h := make(HeroList, 0)
	h = append(h, heros...)
	(&h).SortBy()
	for _, hero := range h {
		fmt.Printf("kind: %v, name: %s\n", hero.Profession, hero.Name)
	}
}

func SortHero() {
	heros := []*Hero{
		&Hero{"a", Up},
		&Hero{"b", Middle},
		&Hero{"c", Gank},
		&Hero{"d", Down},
		&Hero{"e", Assassin},
		&Hero{"g", Middle},
	}

	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Profession != heros[j].Profession {
			return heros[i].Profession < heros[j].Profession
		}
		return heros[i].Name < heros[j].Name

	})

	for _, h := range heros {
		fmt.Printf("kind: %v, name: %s\n", h.Profession, h.Name)
	}
}

func IterString() {
	str := "abAB"
	fmt.Printf("str: %s\n", str)
	for s, value := range str {
		fmt.Printf("s: %v, value: %v, ascii: %v\n", s, value, rune(value))
	}

	r := 'a'
	fmt.Printf("r: %v, t: %T\n", r, r)
	fmt.Printf("r: %v\n", string(r))
	fmt.Print(rune(97))

	fmt.Println("-------------")

	heros := []string{"b", "c", "a"}
	sort.Strings(heros)
	for _, hero := range heros {
		fmt.Printf("hero: %v\n", hero)
	}
}
