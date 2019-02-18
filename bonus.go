package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type BonusType int

const (
	Percentage BonusType = 0
	Number     BonusType = 1
	Boolean    BonusType = 2
)

type BonusDirection int

const (
	Add      BonusDirection = 0
	Subtract BonusDirection = 1
	Base     BonusDirection = 2
)

type Bonus struct {
	Name      string
	Value     int
	Type      BonusType
	Direction BonusDirection
}

type Bonuses struct {
	Bonus   []Bonus
	affects map[string]*StringSet
}

func NewBonuses() Bonuses {
	bonuses := Bonuses{}
	bonuses.affects = make(map[string]*StringSet)

	bonuses.affects["run"] = NewStringSet()
	bonuses.affects["run"].Add("RunSpeed")

	bonuses.affects["walk"] = NewStringSet()
	bonuses.affects["walk"].Add("WalkSpeed")

	bonuses.affects["jump"] = NewStringSet()
	bonuses.affects["jump"].Add("JumpDistance")

	bonuses.affects["dfa"] = NewStringSet()
	bonuses.affects["dfa"].Add("DFA")
	bonuses.affects["dfa"].Add("DFA2")

	bonuses.affects["dfaSelf"] = NewStringSet()
	bonuses.affects["dfaSelf"].Add("DFASelfDamage")

	bonuses.affects["melee"] = NewStringSet()
	bonuses.affects["melee"].Add("Melee")
	bonuses.affects["melee"].Add("Melee2")

	bonuses.affects["meleeStab"] = NewStringSet()
	bonuses.affects["meleeStab"].Add("MeleeStab")

	bonuses.affects["targetHeat"] = NewStringSet()
	bonuses.affects["targetHeat"].Add("Inferno")
	bonuses.affects["targetHeat"].Add("ThermoBurn")
	bonuses.affects["targetHeat"].Add("InfernoIV")
	bonuses.affects["targetHeat"].Add("ImpHeatSink")
	bonuses.affects["targetHeat"].Add("HeatDamage")

	bonuses.affects["selfHeat"] = NewStringSet()
	bonuses.affects["selfHeat"].Add("HeatGenerated")
	bonuses.affects["selfHeat"].Add("HeatPerTurn")

	bonuses.affects["weaponHeat"] = NewStringSet()
	bonuses.affects["weaponHeat"].Add("TEHeatgen")

	return bonuses
}

func (b Bonuses) ApplyBonus(bonusGroup string, stat int) int {
	var newStat = float64(stat)
	bonuses := b.GetBonus(bonusGroup)

	for b := range bonuses {
		fmt.Printf("%v\n", bonuses[b])

		if bonuses[b].Type == Percentage && bonuses[b].Direction == Add {
			newStat += newStat * float64(bonuses[b].Value) / 100
		} else if bonuses[b].Type == Percentage && bonuses[b].Direction == Subtract {
			newStat -= newStat * float64(bonuses[b].Value) / 100
		} else if bonuses[b].Type == Number && bonuses[b].Direction == Add {
			newStat += float64(bonuses[b].Value)
		} else if bonuses[b].Type == Number && bonuses[b].Direction == Subtract {
			newStat -= float64(bonuses[b].Value)
		} else if bonuses[b].Type == Boolean {
			fmt.Printf("I don't know what to do with Boolean yet: %d\n", bonuses[b].Value)
			return stat
		} else if bonuses[b].Direction == Base {
			fmt.Printf("I don't know what to do with Base yet: %d\n", bonuses[b].Value)
			return stat
		}
	}

	return int(math.Round(newStat))
}

func (b Bonuses) GetBonus(bonusGroup string) []Bonus {
	var bList []Bonus
	for bonus := range b.Bonus {
		if b.affects[bonusGroup].Contains(b.Bonus[bonus].Name) {
			bList = append(bList, b.Bonus[bonus])
		}
	}

	return bList
}

func (b *Bonuses) AddBonus(bonus string) {
	newBonus := Bonus{}
	s := strings.Split(bonus, ": ")
	newBonus.Name = s[0]

	if len(s) == 1 {
		newBonus.Type = Boolean
		b.Bonus = append(b.Bonus, newBonus)
		return
	}

	if strings.Contains(s[1], "%") {
		newBonus.Type = Percentage
	} else {
		newBonus.Type = Number
	}

	if strings.Contains(s[1], "-") {
		newBonus.Direction = Subtract
	} else if strings.Contains(s[1], "+") {
		newBonus.Direction = Add
	} else {
		if newBonus.Type == Percentage {
			newBonus.Direction = Add
		} else {
			newBonus.Direction = Base
		}
	}

	r := strings.NewReplacer("-", "", "+", "", "%", "")
	newBonus.Value, _ = strconv.Atoi(r.Replace(s[1]))
	b.Bonus = append(b.Bonus, newBonus)
}
