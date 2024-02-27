package main

import (
	"fmt"
	"math/rand"
	"time"
)

// type Person struct {
// 	Name string
// 	Age  int
// 	Job  string
// }

// func main() {
// 	people := []Person{
// 		{"Bambang", 20, "Gambler"},
// 		{"Joko", 30, "Engineer"},
// 		{"Siti", 25, "Doctor"},
// 	}

// 	fmt.Println("Initial Information:")
// 	for _, person := range people {
// 		person.GetInfo()
// 		fmt.Println()

// 		for i := 0; i < 5; i++ {
// 			person.AddYear()
// 		}

// 		fmt.Println("After 5 years:")
// 		person.GetInfo()
// 		fmt.Println()
// 	}
// }
// func (p *Person) GetInfo() {
// 	fmt.Printf("Name: %s\nAge: %d\nJob: %s\n", p.Name, p.Age, p.Job)
// }
// func (p *Person) AddYear() {
// 	p.Age++
// 	if p.Age >= 50 {
// 		p.Job = "Retired"
// 	}
// }

// -----------------

type Weapon struct {
	Attack int
}

type Hero struct {
	Name           string
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealthPoint    int
	Weapon         Weapon
}

func (h *Hero) CountDamage() int {
	rand.Seed(time.Now().UnixNano())
	totalDamage := h.BaseAttack + h.Weapon.Attack
	if rand.Intn(2) == 0 {
		totalDamage += h.CriticalDamage
	}
	return totalDamage
}

func (h *Hero) isAttackedBy(attacker Hero) {
	totalDamageReceived := attacker.CountDamage() - h.Defence
	if totalDamageReceived < 0 {
		totalDamageReceived = 0
	}
	h.HealthPoint -= totalDamageReceived
	fmt.Printf("%s receives %d damage. HealthPoint now: %d\n", h.Name, totalDamageReceived, h.HealthPoint)
}

func Battle(attacker, defender Hero) {
	defender.isAttackedBy(attacker)
}

func main() {
	weapon := Weapon{Attack: 20}
	hero := Hero{
		Name:           "Hero1",
		BaseAttack:     50,
		Defence:        30,
		CriticalDamage: 25,
		HealthPoint:    100,
		Weapon:         weapon,
	}
	totalDamage := hero.CountDamage()

	fmt.Printf("Total damage dealt by %s: %d\n", hero.Name, totalDamage)

	weapon2 := Weapon{Attack: 15}
	hero2 := Hero{Name: "Hero2", BaseAttack: 40, Defence: 35, CriticalDamage: 20, HealthPoint: 100, Weapon: weapon2}
	Battle(hero, hero2)
}
