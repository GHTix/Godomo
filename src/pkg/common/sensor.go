package common

type Sensor struct {
	Name  string
	Value float64
	Unit  string
}

func NewSensor(name string, value float64, unit string) *Sensor {
	return &Sensor{
		Name:  name,
		Value: value,
		Unit:  unit,
	}
}
