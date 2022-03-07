package main

import (
	"fmt"

	fig "github.com/Mateodioev/testgo"
)

func main() {
  
  cir := fig.Circle{Radio: 5}
	fig.Show(&cir)
  
	fmt.Println("")
  
	squa := fig.Square{Lado: 5}
	fig.Show(&squa)
  
}
