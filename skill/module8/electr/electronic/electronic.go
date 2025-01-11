package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}
type StationPhone interface {
	ButtonsCount() int
}
type SmartPhone interface {
	OS() string
}
type applePhone struct {
	os    string
	model string
}
type androidPhone struct {
	brand string
	os    string
	model string
}
type radioPhone struct {
	brand   string
	model   string
	buttons int
}

func NewApplePhone(model string) Phone {
	return &applePhone{
		model: model,
		os:    "IOS",
	}
}
func (p *applePhone) Brand() string {
	return "apple"
}
func (p *applePhone) Model() string {
	return p.model
}
func (p applePhone) Type() string {
	return "smartphone"
}
func (p *applePhone) OS() string {
	return p.os
}

func NewAndroidPhone(brand, model string) Phone {
	return &androidPhone{
		brand: brand,
		model: model,
		os:    "Android",
	}
}
func (p *androidPhone) Brand() string {
	return p.brand
}
func (p *androidPhone) Model() string {
	return p.model
}
func (p *androidPhone) Type() string {
	return "smartphone"
}
func (p *androidPhone) OS() string {
	return p.os
}
func NewRadioPhone(brand, model string, buttons int) Phone {
	return &radioPhone{
		brand:   brand,
		model:   model,
		buttons: buttons,
	}
}
func (p *radioPhone) Brand() string {
	return p.brand
}
func (p *radioPhone) Model() string {
	return p.model
}
func (p *radioPhone) Type() string {
	return "station"
}
func (p *radioPhone) ButtonsCount() int {
	return p.buttons
}
