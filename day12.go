package main

import (
	"2023-advent-of-code/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunDay12() {
	file, err := os.Open("./inputs/day12inputsmall.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalArrangements := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		halves := strings.Split(lineText, " ")
		damagedLengths := helpers.ParseInts(halves[1])
		damagedLengthStack := helpers.Stack{}
		for _, v := range damagedLengths {
			damagedLengthStack.Push(v)
		}
		springGroupsTemp := strings.Split(halves[0], ".")
		springGroups := []string{}
		for _, v := range springGroupsTemp {
			if len(v) == 0 {
				continue
			}
			springGroups = append(springGroups, v)
		}

		groupArrangements := 0
		for i := len(springGroups) - 1; i >= 0; i-- {
			group := springGroups[i]
			damaged, _ := damagedLengthStack.Top()
			damagedLengthStack.Pop()
			thisGroupArrangements := calcArrangements(
				group, []int{damaged}, damaged, 0, &damagedLengthStack)
			if thisGroupArrangements > 0 {
				if groupArrangements == 0 {
					groupArrangements = thisGroupArrangements
				} else {
					groupArrangements *= thisGroupArrangements
				}
			}
		}
		totalArrangements += groupArrangements
	}
	fmt.Println(totalArrangements)
}

func iterGroup(group string, damageList []int, operational int, current string) int {
	if group == "" && len(damageList) == 0 && operational == 0 {
		return 1
	}

	combos := 0
	combos += applyOperational(group, damageList, operational)
	combos += applySet(group, damageList, operational, current)
	return combos
}

func applyOperational(group string, damageList []int, operational int) int {
	if operational == 0 {
		return 0
	}

	if len(group) > 0 && group[0] == '#' {
		return 0
	}

	if len(group) == 0 {
		return 0
	}

	return iterGroup(group[1:], damageList, operational-1, group)
}

func applySet(group string, damageList []int, operational int, current string) int {
	if len(damageList) == 0 {
		return 0
	}

	if damageList[len(damageList)-1] > len(group) {
		return 0
	}
	nextGroup := group[damageList[len(damageList)-1]:]
	nextDamageList := damageList[:len(damageList)-1]
	nextOperational := operational

	appliedForcedOperational := applyOperational(nextGroup, nextDamageList, nextOperational)

	if appliedForcedOperational == 0 && len(nextDamageList) > 0 {
		return 0
	}

	if appliedForcedOperational == 0 && len(nextDamageList) == 0 {
		return 1
	}

	return appliedForcedOperational
}

func calcArrangements(group string, damaged []int, damagedTotal int, operational int, damageStack *helpers.Stack) int {
	if len(group) == damagedTotal+operational {
		return iterGroup(group, damaged, operational, group)
	}

	if len(group) < damagedTotal+operational {
		return 0
	}

	lastDamageAdded := 0
	for damagedTotal+operational < len(group) {
		d, _ := damageStack.Top()
		damageStack.Pop()
		lastDamageAdded = d
		damagedTotal += d
		if d > 0 {
			damaged = append(damaged, d)
		}
		operational++
	}

	if damagedTotal+operational > len(group) {
		damagedTotal -= lastDamageAdded
		if len(damaged) > 0 {
			damaged = damaged[1:]
		}
		damageStack.Push(lastDamageAdded)
	}

	return calcArrangements(group, damaged, damagedTotal, operational, damageStack)
}
