package figuras

import "fmt"

type Figure interface {
	area() float64
	perimeter() float64
}

func Show(f Figure) {
	fmt.Println("Medidas", f)
	fmt.Println("Area:", f.area())
	fmt.Println("Perimeter:", f.perimeter())
}