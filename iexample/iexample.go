package iexample

import (
	"errors"
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perim() (float64, error)
}

type rect struct{
	width, heigth float64
}

type circle struct{
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.heigth
}

func (r *rect) perim() (float64, error){
	if r.width <= 0 {
		return 0, errors.New("width cannot be 0")
	}
	return 2*r.width + 2*r.heigth, nil
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64{
	return 2 * math.Pi * c.radius 

}

func measure(g geometry) (float64, float64, error){
	perm, err := g.perim()
	if err != nil {
		return 0, 0, err
	}
	return g.area(), perm, nil

}

func MainStart(){	

	r := rect{width: 10, heigth: 5}
	circle := circle{radius: 5}
	
	// fmt.Println("Area", rectangle.area())
	// fmt.Println("Perimeter", rectangle.perim())
	// fmt.Println("Area", circle.area())
	// fmt.Println("Perimeter", circle.perim())
	a, b, err := measure(r)
	if err != nil {
		fmt.Println("%w",err)
	}
	fmt.Println(a)
	fmt.Println(b)
	// measure(c)
}

func GoRoutingExample() {
	Routine("direct")

	go Routing("goroutines")
	go Routine("goroutines--11")

	go func(msg string){
		fmt.Println(msg)
	}("going...")
	time.Sleep(time.Second)
	fmt.Println("done")
	}

	func Routine(from string)
		for i := 0; i < 5; i++ {
			fmt.Println(from, ":",i)
			
		}

}