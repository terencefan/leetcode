package main

import "sort"

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func effective(spell, potion int, threshold int64) bool {
	return int64(spell)*int64(potion) >= threshold
}

type SpellIndex struct {
	potency int
	index   int
}

func successfulPairs2Pointers(spells []int, potions []int, success int64) []int {
	if len(spells) == 0 || len(potions) == 0 {
		return []int{}
	}
	r := make([]int, len(spells))

	var spellMaxToMin = make([]SpellIndex, len(spells))
	for i, spell := range spells {
		spellMaxToMin[i] = SpellIndex{spell, i}
	}
	sort.Slice(spellMaxToMin, func(i, j int) bool {
		return spellMaxToMin[i].potency > spellMaxToMin[j].potency
	})

	sort.Slice(potions, func(i, j int) bool {
		return potions[i] < potions[j]
	})

	currentPotionIndex := 0
	for _, spell := range spellMaxToMin {
		for currentPotionIndex < len(potions) && int64(potions[currentPotionIndex])*int64(spell.potency) < success {
			currentPotionIndex++
		}
		r[spell.index] = len(potions) - currentPotionIndex
	}
	return r
}

func binarySearch(spell int, success int64, potions []int) int {
	l, r := 0, len(potions)
	for l < r {
		mid := l + (r-l)/2
		if int64(spell)*int64(potions[mid]) >= success {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Slice(potions, func(i, j int) bool {
		return potions[i] > potions[j]
	})

	var r = make([]int, len(spells))
	for i, spell := range spells {
		r[i] = binarySearch(spell, success, potions)
	}
	return r
}
