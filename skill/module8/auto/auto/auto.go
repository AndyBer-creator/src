package main

import "fmt"

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value
	if t != u.T {
		switch t {
		case Inch:
			if u.T == CM {
				value = u.Value / 2.54 // конвертация см в дюймы
			}
		case CM:
			if u.T == Inch {
				value = u.Value * 2.54 // конвертация дюймов в см
			}
		}
	}
	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type DimensionsInches struct {
	length Unit
	width  Unit
	height Unit
}

func (d DimensionsInches) Length() Unit {
	return d.length
}

func (d DimensionsInches) Width() Unit {
	return d.width
}

func (d DimensionsInches) Height() Unit {
	return d.height
}

type DimensionsCM struct {
	length Unit
	width  Unit
	height Unit
}

func (d DimensionsCM) Length() Unit {
	return d.length
}

func (d DimensionsCM) Width() Unit {
	return d.width
}

func (d DimensionsCM) Height() Unit {
	return d.height
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type BMW struct {
	model       string
	dimensions  DimensionsCM
	maxSpeed    int
	enginePower int
}

func (a BMW) Brand() string {
	return "BMW"
}

func (a BMW) Model() string {
	return a.model
}

func (a BMW) Dimensions() Dimensions {
	return a.dimensions
}

func (a BMW) MaxSpeed() int {
	return a.maxSpeed
}

func (a BMW) EnginePower() int {
	return a.enginePower
}

type Mercedes struct {
	model       string
	dimensions  DimensionsCM
	maxSpeed    int
	enginePower int
}

func (a Mercedes) Brand() string {
	return "Mercedes"
}

func (a Mercedes) Model() string {
	return a.model
}

func (a Mercedes) Dimensions() Dimensions {
	return a.dimensions
}

func (a Mercedes) MaxSpeed() int {
	return a.maxSpeed
}

func (a Mercedes) EnginePower() int {
	return a.enginePower
}

type Dodge struct {
	model       string
	dimensions  DimensionsInches
	maxSpeed    int
	enginePower int
}

func (a Dodge) Brand() string {
	return "Dodge"
}

func (a Dodge) Model() string {
	return a.model
}

func (a Dodge) Dimensions() Dimensions {
	return a.dimensions
}

func (a Dodge) MaxSpeed() int {
	return a.maxSpeed
}

func (a Dodge) EnginePower() int {
	return a.enginePower
}
func main() {
	bmw := BMW{
		model: "X5",
		dimensions: DimensionsCM{
			length: Unit{Value: 4922, T: CM},
			width:  Unit{Value: 2004, T: CM},
			height: Unit{Value: 1696, T: CM},
		},
		maxSpeed:    250,
		enginePower: 400,
	}

	mercedes := Mercedes{
		model: "C-Class",
		dimensions: DimensionsCM{
			length: Unit{Value: 4860, T: CM},
			width:  Unit{Value: 1810, T: CM},
			height: Unit{Value: 1440, T: CM},
		},
		maxSpeed:    250,
		enginePower: 350,
	}

	dodge := Dodge{
		model: "Charger",
		dimensions: DimensionsInches{
			length: Unit{Value: 197, T: Inch},
			width:  Unit{Value: 75, T: Inch},
			height: Unit{Value: 58, T: Inch},
		},
		maxSpeed:    200,
		enginePower: 300,
	}

	printAutoDetails(bmw)
	printAutoDetails(mercedes)
	printAutoDetails(dodge)
}
func printAutoDetails(a Auto) {
	fmt.Printf("Brand: %s, Model: %s\n", a.Brand(), a.Model())
	dimensions := a.Dimensions()
	fmt.Printf("Dimensions - Length: %.2f %s, Width: %.2f %s, Height: %.2f %s\n",
		dimensions.Length().Get(CM), dimensions.Length().T,
		dimensions.Width().Get(CM), dimensions.Width().T,
		dimensions.Height().Get(CM), dimensions.Height().T)
	fmt.Printf("Max Speed: %d km/h, Engine Power: %d HP\n\n", a.MaxSpeed(), a.EnginePower())
}
