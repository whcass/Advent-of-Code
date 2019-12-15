package main

import "fmt"

// <x=8, y=0, z=8>
// <x=0, y=-5, z=-10>
// <x=16, y=10, z=-5>
// <x=19, y=-10, z=-7>

// <x=-1, y=0, z=2>
// <x=2, y=-10, z=-7>
// <x=4, y=-8, z=8>
// <x=3, y=5, z=-1>
func main() {
	// A := NewPlanet(8, 0, 8)
	// B := NewPlanet(0, -5, -10)
	// C := NewPlanet(16, 10, -5)
	// D := NewPlanet(19, -10, -7)
	Part1()
	Part2()

}

func Part2() {
	// A := NewPlanet(-1, 0, 2)
	// B := NewPlanet(2, -10, -7)
	// C := NewPlanet(4, -8, 8)
	// D := NewPlanet(3, 5, -1)
	A := NewPlanet(8, 0, 8)
	B := NewPlanet(0, -5, -10)
	C := NewPlanet(16, 10, -5)
	D := NewPlanet(19, -10, -7)
	otherPlanets := []*Planet{A, B, C, D}
	steps := 0
	// xmatch, ymatch, zmatch := false, false, false
	states := []string{}
	for {
		A.UpdateZ(otherPlanets[1:])
		B.UpdateZ(otherPlanets[2:])
		C.UpdateZ(otherPlanets[3:])
		D.UpdateZ(otherPlanets[4:])
		checksum := A.CheckSum() + B.CheckSum() + C.CheckSum() + D.CheckSum()
		match := false
		for _, v := range states {
			if v == checksum {
				match = true
				break
			}
		}
		if match {
			fmt.Printf("[*] Steps - %d\n", steps)
			break
		} else {
			states = append(states, checksum)
		}
		steps++
	}
}

func Part1() {
	A := NewPlanet(8, 0, 8)
	B := NewPlanet(0, -5, -10)
	C := NewPlanet(16, 10, -5)
	D := NewPlanet(19, -10, -7)

	otherPlanets := []*Planet{A, B, C, D}
	for i := 1; i <= 1000; i++ {
		A.Update(otherPlanets[1:])
		B.Update(otherPlanets[2:])
		C.Update(otherPlanets[3:])
		D.Update(otherPlanets[4:])
	}

	total := 0
	for _, p := range otherPlanets {
		total += p.GetEnergy()
	}

	fmt.Printf("[*] Energy - %d\n", total)
}
