//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	Name              string
	Health, MaxHealth uint
	Energy            uint
	MaxEnergy         uint
}

func (p *Player) addHealth(amount uint) {
	p.Health += amount
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}

	fmt.Println("Health added [", amount, "]")
	fmt.Println("    New Health:", p.Health)
}

func (p *Player) applyDamage(amount uint) {
	if amount > p.Health {
		p.Health = 0
	} else {
		p.Health -= amount
	}

	fmt.Println("Health expended [", amount, "]")
	fmt.Println("    New Health:", p.Health)
}

func (p *Player) addEnergy(amount uint) {
	p.Energy += amount
	if p.Energy > p.MaxEnergy {
		p.Energy = p.MaxEnergy
	}

	fmt.Println("Energy added [", amount, "]")
	fmt.Println("    New Energy:", p.Energy)
}

func (p *Player) consumeEnergy(amount uint) {
	if amount > p.Energy {
		p.Energy = 0
	} else {
		p.Energy -= amount
	}

	fmt.Println("Energy expended [", amount, "]")
	fmt.Println("    New Energy:", p.Energy)
}

func main() {
	player := Player{"Daniel", 100, 100, 75, 100}

	player.applyDamage(10)
	player.addEnergy(5)
	player.consumeEnergy(5)
	player.consumeEnergy(5)
	player.consumeEnergy(5)
	player.consumeEnergy(5)
	player.consumeEnergy(5)
	player.consumeEnergy(5)
	player.consumeEnergy(5)
	player.consumeEnergy(5)
	player.consumeEnergy(5)
	player.consumeEnergy(5)
	player.consumeEnergy(5)
	player.addEnergy(40)
	player.addHealth(15)
	player.consumeEnergy(999)
}
