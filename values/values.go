package values

import "fmt"


func Variable() {
	var a, b, c uint8 = 5, 6, 7

	x := 100

	// a += 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x)
}

func Values() {
	fmt.Println("First" + "Last")

	fmt.Println("10 - 1 = ", 10-1)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
