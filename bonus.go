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
	bonuses.affects["run"].Add("ActivatedRunSpeed")

	bonuses.affects["walk"] = NewStringSet()
	bonuses.affects["walk"].Add("ActiveWalkSpeed")

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

	bonuses.affects["targetDamage"] = NewStringSet()
	bonuses.affects["targetDamage"].Add("LRMDamage")

	bonuses.affects["targetHeat"] = NewStringSet()
	bonuses.affects["targetHeat"].Add("Inferno")
	bonuses.affects["targetHeat"].Add("ThermoBurn")
	bonuses.affects["targetHeat"].Add("InfernoIV")
	bonuses.affects["targetHeat"].Add("ImpHeatSink")
	bonuses.affects["targetHeat"].Add("HeatDamage")

	bonuses.affects["targetStab"] = NewStringSet()
	bonuses.affects["targetStab"].Add("StabDamage")

	bonuses.affects["selfHeat"] = NewStringSet()
	bonuses.affects["selfHeat"].Add("HeatGenerated")
	bonuses.affects["selfHeat"].Add("ActiveHeatPerTurn")
	bonuses.affects["selfHeat"].Add("AMSHeat")

	bonuses.affects["weaponHeat"] = NewStringSet()
	bonuses.affects["weaponHeat"].Add("TEHeatgen")

	return bonuses
}

func (b Bonuses) ApplyBonus(bonus Bonus, stat int) int {
	var newStat = float64(stat)

	if bonus.Type == Percentage && bonus.Direction == Add {
		newStat += newStat * float64(bonus.Value) / 100
	} else if bonus.Type == Percentage && bonus.Direction == Subtract {
		newStat -= newStat * float64(bonus.Value) / 100
	} else if bonus.Type == Number && bonus.Direction == Add {
		newStat += float64(bonus.Value)
	} else if bonus.Type == Number && bonus.Direction == Subtract {
		newStat -= float64(bonus.Value)
	} else if bonus.Type == Boolean {
		fmt.Printf("I don't know what to do with Boolean yet: %d\n", bonus.Value)
		return stat
	} else if bonus.Direction == Base {
		newStat += float64(bonus.Value)
		return stat
	}

	return int(math.Round(newStat))
}

func (b Bonuses) getBonus(bonus string) (bonusType string, newBonus Bonus) {
	s := strings.Split(bonus, ": ")
	newBonus.Name = s[0]

	for bt := range b.affects {
		if b.affects[bt].Contains(newBonus.Name) {
			bonusType = bt
		}
	}

	if len(s) == 1 {
		newBonus.Type = Boolean
		return bonusType, newBonus
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
		} else if newBonus.Name == "Inferno" {
			newBonus.Direction = Add
		} else {
			newBonus.Direction = Base
		}
	}

	r := strings.NewReplacer("-", "", "+", "", "%", "")
	newBonus.Value, _ = strconv.Atoi(r.Replace(s[1]))

	return bonusType, newBonus
}
