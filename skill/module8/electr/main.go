package main

import (
	"fmt"
	"skill/module8/electr/electronic"
)

func printCharacteristic(p electronic.Phone) {
	fmt.Printf("Brand: %s\n", p.Brand())
	fmt.Printf("Model: %s\n", p.Model())
	fmt.Printf("Type: %s\n", p.Type())
	if smartphone, ok := p.(electronic.SmartPhone); ok {
		fmt.Printf("OS: %s\n", smartphone.OS())
	}
	if stationPhone, ok := p.(electronic.StationPhone); ok {
		fmt.Printf("Buttons Number: %d\n", stationPhone.ButtonsCount())
	}

}
func main() {
	applePhone := electronic.NewApplePhone("iPhone 99")
	androidPhone := electronic.NewAndroidPhone("Samsung", "Note11+")
	radioPhone := electronic.NewRadioPhone("Panasonic", "KX34234", 18)
	printCharacteristic(applePhone)
	printCharacteristic(androidPhone)
	printCharacteristic(radioPhone)
}
