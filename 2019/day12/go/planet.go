package main

import (
	"fmt"
)

type Planet struct {
	pos     *vector
	vel     *vector
	origpos *vector
}

type vector struct {
	x int
	y int
	z int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p *Planet) Reset() {
	p.pos.x, p.pos.y, p.pos.z = p.origpos.x, p.origpos.y, p.origpos.z
	p.vel.x, p.vel.y, p.vel.z = 0, 0, 0
}

func (p *Planet) GetEnergy() int {
	return (Abs(p.pos.x) + Abs(p.pos.y) + Abs(p.pos.z)) * (Abs(p.vel.x) + Abs(p.vel.y) + Abs(p.vel.z))
}

func (p *Planet) Print() {
	fmt.Printf(
		"pos=<x=%d, y=%d, z=%d, vel=<x=%d, y=%d, z=%d>\n",
		p.pos.x,
		p.pos.y,
		p.pos.z,
		p.vel.x,
		p.vel.y,
		p.vel.z,
	)
}

func (lhs *Planet) checkX(rhs *Planet) {
	if lhs.pos.x > rhs.pos.x {
		lhs.vel.x--
		rhs.vel.x++
	} else if lhs.pos.x < rhs.pos.x {
		lhs.vel.x++
		rhs.vel.x--
	}
}

func (lhs *Planet) checkY(rhs *Planet) {
	if lhs.pos.y > rhs.pos.y {
		lhs.vel.y--
		rhs.vel.y++
	} else if lhs.pos.y < rhs.pos.y {
		lhs.vel.y++
		rhs.vel.y--
	}
}

func (lhs *Planet) checkZ(rhs *Planet) {
	if lhs.pos.z > rhs.pos.z {
		lhs.vel.z--
		rhs.vel.z++
	} else if lhs.pos.z < rhs.pos.z {
		lhs.vel.z++
		rhs.vel.z--
	}
}

func (p *Planet) CheckSum() string {
	return fmt.Sprintf("%d%d%d%d%d%d",
		p.pos.x,
		p.pos.y,
		p.pos.z,
		p.vel.x,
		p.vel.y,
		p.vel.z)
}

func (p *Planet) addVelocity() {
	p.pos.x += p.vel.x
	p.pos.y += p.vel.y
	p.pos.z += p.vel.z
}

func (lhs *Planet) Update(otherPlanets []*Planet) {
	for _, rhs := range otherPlanets {
		lhs.checkX(rhs)
		lhs.checkY(rhs)
		lhs.checkZ(rhs)
	}
	lhs.addVelocity()
}

func (lhs *Planet) UpdateX(otherPlanets []*Planet) {
	for _, rhs := range otherPlanets {
		lhs.checkX(rhs)
	}
	lhs.addVelocity()
}

func (lhs *Planet) UpdateY(otherPlanets []*Planet) {
	for _, rhs := range otherPlanets {
		lhs.checkY(rhs)
	}
	lhs.addVelocity()
}

func (lhs *Planet) UpdateZ(otherPlanets []*Planet) {
	for _, rhs := range otherPlanets {
		lhs.checkZ(rhs)
	}
	lhs.addVelocity()
}

// NewPlanet: Create a new planet
func NewPlanet(x int, y int, z int) *Planet {
	pos := &vector{x, y, z}
	origpos := &vector{x, y, z}
	vel := &vector{0, 0, 0}
	return &Planet{pos, vel, origpos}
}
