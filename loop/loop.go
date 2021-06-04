package loop

import "fmt"

func LoopStatement() {
	sum := 0
	i := 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}
