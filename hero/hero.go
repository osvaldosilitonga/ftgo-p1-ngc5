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
	}

	return hero.BaseAttack + hero.Weapon.Attack
}

func (hero Hero) isAttackBy(damage int) int {
	total := damage - hero.Defence

	if total > 0 {
		hero.HealthPoint -= total
	}

	return hero.HealthPoint
}

func Battle(hero1, hero2 Hero) {

	damagePenyerang := hero1.CountDamage()
	healthHeroDiserang := hero2.isAttackBy(damagePenyerang)

	fmt.Println("Health point hero yang diserang : ", healthHeroDiserang)

}
