package figuras

import "math"

// Circulo
type Circle struct {
	Radio float64
}

// Calcular el area de un circulo
func (c *Circle) area() float64 {
	return math.Pi * (c.Radio * c.Radio)
}

// Calcular el perímetro de un circulo
func (c *Circle) perimeter() float64{
	return (2 * math.Pi) * c.Radio
}
