package rangeexample

import "fmt"


type rect struct{
	width, heigth int
}

func (r *rect) area() int{
	return r.width * r.heigth
}

func (r *rect) perim() int{
	return 2*r.width + 2*r.heigth


}

func MethodExample(){	

	rectangle := rect{width: 5, heigth: 2}

	fmt.Println("Area", rectangle.area())
	fmt.Println("Perimeter", rectangle.perim())
}