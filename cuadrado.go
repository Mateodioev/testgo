package figuras

// Cuadrado
type Square struct {
	Lado float64
}

// Calcular el area de un cuadrado
func (s *Square) area() float64{
	return s.Lado * s.Lado
}
// Calcular el perímetro de un cuadrado
func (s *Square) perimeter() float64{
	return 4 * s.Lado
}
