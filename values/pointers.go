package values

import (
	"fmt"
)

func PointersValue() {
	var p *int
	i := 42
	p = &i
	fmt.Println(*p)

}
