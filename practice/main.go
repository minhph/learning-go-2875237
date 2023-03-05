package main

import (
	"fmt"
)

func main() {
	myAircraft := Aircraft{"A6M", "H6-Renzo"}
	myAircraft.Fire()
	myAircraft.FireDouble()
	myAircraft.FireDouble()
}

type Aircraft struct {
	Name    string
	Missile string
}

func (a Aircraft) Fire() {
	fmt.Println(a.Missile)
}

func (a Aircraft) FireDouble() {
	a.Missile = fmt.Sprintf("%v %v", a.Missile, a.Missile)
	fmt.Println(a.Missile)
}
