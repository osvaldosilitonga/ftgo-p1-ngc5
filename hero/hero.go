package hero

import (
	"fmt"
	"math/rand"
)

type Hero struct {
	Name           int
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealthPoint    int
	Weapon         Weapon
}

type Weapon struct {
	Attack int
}

func (hero Hero) CountDamage() int {
	randNumber := rand.Intn(100)
	if randNumber%2 == 0 {
		return hero.BaseAttack + hero.Weapon.Attack + hero.CriticalDamage

	} else {
		return hero.BaseAttack + hero.Weapon.Attack
	}
}

func (hero Hero) isAttackBy(damage int) int {
	total := damage - hero.Defence

	if total > 0 {
		hero.HealthPoint -= total
	}

	return hero.HealthPoint
}

func Battle(hero1, hero2 map[string]int) {

	penyerang := Hero{
		Name:           hero1["name"],
		BaseAttack:     hero1["baseAttack"],
		Defence:        hero1["defence"],
		CriticalDamage: hero1["criticalDamage"],
		HealthPoint:    hero1["healthPoint"],
		Weapon:         Weapon{Attack: hero1["weapon"]},
	}

	diserang := Hero{
		Name:           hero2["name"],
		BaseAttack:     hero2["baseAttack"],
		Defence:        hero2["defence"],
		CriticalDamage: hero2["criticalDamage"],
		HealthPoint:    hero2["healthPoint"],
		Weapon:         Weapon{Attack: hero2["weapon"]},
	}

	damagePenyerang := penyerang.CountDamage()
	healthHeroDiserang := diserang.isAttackBy(damagePenyerang)

	fmt.Println("Health point hero yang diserang : ", healthHeroDiserang)

}
