package main

import (
	"fmt"
)

type boosterT struct {
	ID                                                           int
	Engineering, Experimental                                    string
	ShieldStrengthBonus, ExpResBonus, KinResBonus, ThermResBonus float32
}

/*
 *
 */
func loadBoosters(shortList bool) []boosterT {

	fmt.Println("Load shield booster variants")

	var shortBoosterVariants = []boosterT{
		boosterT{
			1, "Kinetic Resistance", "Thermo Block", 0.188, -0.05, 0.27, -0.029,
		},
		boosterT{
			2, "Kinetic Resistance", "Force Block", 0.188, -0.05, 0.2846, -0.05,
		},
		boosterT{
			3, "Kinetic Resistance", "Super Capacitors", 0.26, -0.071, 0.2554, -0.071,
		},
		boosterT{
			4, "Thermal Resistance", "Thermo Block", 0.188, -0.05, -0.05, 0.2846,
		},
		boosterT{
			5, "Thermal Resistance", "Force Block", 0.188, -0.05, -0.029, 0.27,
		},
		boosterT{
			6, "Thermal Resistance", "Super Capacitors", 0.26, -0.071, -0.071, 0.2554,
		},
		boosterT{
			7, "Heavy Duty", "Thermo Block", 0.639, 0, 0, 0.02,
		},
		boosterT{
			8, "Heavy Duty", "Force Block", 0.639, 0, 0.02, 0,
		},
		boosterT{
			9, "Heavy Duty", "Super Capacitors", 0.739, -0.02, -0.02, -0.02,
		},
		boosterT{
			10, "Resistance Augmented", "Thermo Block", 0.188, 0.17, 0.17, 0.1866,
		},
		boosterT{
			11, "Resistance Augmented", "Force Block", 0.188, 0.17, 0.1866, 0.17,
		},
		boosterT{
			12, "Resistance Augmented", "Super Capacitors", 0.26, 0.1534, 0.1534, 0.1534,
		},
	}

	var fullBoosterVariants = []boosterT{
		boosterT{
			1, "Kinetic Resistance", "Thermo Block", 0.188, -0.05, 0.27, -0.029,
		},
		boosterT{
			2, "Kinetic Resistance", "Force Block", 0.188, -0.05, 0.2846, -0.05,
		},
		boosterT{
			3, "Kinetic Resistance", "Blast Block", 0.188, -0.029, 0.27, -0.05,
		},
		boosterT{
			4, "Kinetic Resistance", "Super Capacitors", 0.26, -0.071, 0.2554, -0.071,
		},
		boosterT{
			5, "Thermal Resistance", "Thermo Block", 0.188, -0.05, -0.05, 0.2846,
		},
		boosterT{
			6, "Thermal Resistance", "Force Block", 0.188, -0.05, -0.029, 0.27,
		},
		boosterT{
			7, "Thermal Resistance", "Blast Block", 0.188, -0.029, -0.05, 0.27,
		},
		boosterT{
			8, "Thermal Resistance", "Super Capacitors", 0.26, -0.071, -0.071, 0.2554,
		},
		boosterT{
			9, "Blast Resistance", "Thermo Block", 0.188, 0.27, -0.05, -0.029,
		},
		boosterT{
			10, "Blast Resistance", "Force Block", 0.188, 0.27, -0.029, -0.05,
		},
		boosterT{
			11, "Blast Resistance", "Blast Block", 0.188, 0.2846, -0.05, -0.05,
		},
		boosterT{
			12, "Blast Resistance", "Super Capacitors", 0.26, -0.2554, 0.071, -0.071,
		},
		boosterT{
			13, "Heavy Duty", "Thermo Block", 0.639, 0, 0, 0.02,
		},
		boosterT{
			14, "Heavy Duty", "Force Block", 0.639, 0, 0.02, 0,
		},
		boosterT{
			15, "Heavy Duty", "Blast Block", 0.639, 0.02, 0, 0,
		},
		boosterT{
			16, "Heavy Duty", "Super Capacitors", 0.739, -0.02, -0.02, -0.02,
		},
		boosterT{
			17, "Resistance Augmented", "Thermo Block", 0.188, 0.17, 0.17, 0.1866,
		},
		boosterT{
			18, "Resistance Augmented", "Force Block", 0.188, 0.17, 0.1866, 0.17,
		},
		boosterT{
			19, "Resistance Augmented", "Blast Block", 0.188, 0.1866, 0.17, 0.17,
		},
		boosterT{
			20, "Resistance Augmented", "Super Capacitors", 0.26, 0.1534, 0.1534, 0.1534,
		},
	}

	if shortList {
		return shortBoosterVariants
	} else {
		return fullBoosterVariants
	}
}

func oneUp(currentLoadOut []int, ShieldBoosterCount int, ShieldBoosterVariants int, CurrentShieldBooster int) []int {

	var NextLoadout = currentLoadOut

	if NextLoadout[CurrentShieldBooster] < (ShieldBoosterVariants - 1) {
		NextLoadout[CurrentShieldBooster]++
	} else {
		if CurrentShieldBooster != (ShieldBoosterCount - 1) {
			NextLoadout = oneUp(NextLoadout, ShieldBoosterCount, ShieldBoosterVariants, CurrentShieldBooster+1)
			NextLoadout[CurrentShieldBooster] = NextLoadout[CurrentShieldBooster+1]
		}
	}

	return NextLoadout
}

func getBoosterLoadoutList(config configT, boosters []boosterT) []int {

	var Loadout []int = make([]int, config.ShieldBoosterCount)
	var LoadoutList []int = make([]int, config.ShieldBoosterCount)

	var ShieldBoosterVariants = len(boosters)

	for run := true; run; run = !(Loadout[config.ShieldBoosterCount-1] == (ShieldBoosterVariants - 1)) {
		Loadout = oneUp(Loadout, config.ShieldBoosterCount, ShieldBoosterVariants, 0)
		m := len(LoadoutList)
		n := m + len(Loadout)
		if n > cap(LoadoutList) {
			newSlice := make([]int, (n+1)*2)
			copy(newSlice, LoadoutList)
			LoadoutList = newSlice
		}
		LoadoutList = LoadoutList[0:n]
		copy(LoadoutList[m:n], Loadout)
	}

	return LoadoutList
}
