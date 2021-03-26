package main

import "fmt"

type Figure struct {
	x float64
	y float64
}

func moveUp(f *Figure) float64 {
	f.y += 1
	return f.x*4 + f.y*4
}

func (f *Figure) isCrashed() bool {
	f.x -= 1
	f.y -= 1
	return true
}

type Person struct {
	name string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.name)
}

type Npc struct {
	Person
	skill int
}

type Player struct {
	person Person
	age    int
}

type Guard struct {
	Person
	location string
}

func structures() {
	fmt.Println("-------------- Structures -------------------")
	var f0 Figure
	f1 := new(Figure)
	f2 := Figure{4, 5}

	fmt.Println(f2.x, f2.y)

	f1.x = 34
	f1.y = 21
	fmt.Println(f1.x, f1.y)

	f0.x = 43
	f0.y = 33
	fmt.Println(f0.x, f0.y)

	fmt.Println("Dest: ", moveUp(&f2))

	fmt.Println("Is crashed? ", f2.isCrashed())

	//
	npc1 := new(Npc)
	npc1.Person.talk()
	npc1.talk()

	player1 := new(Player)
	player1.person.name = "Alex"
	player1.person.talk()
	// player1.talk() - error

	guard1 := new(Guard)
	guard1.Person.name = "Zelon"
	guard1.talk()
}
