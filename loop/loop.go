package loop

import "fmt"

// func LoopStatement() {
// 	sum := 0
// 	i := 5
// 	for i < 10{
// 	for i < 10 {
// 		sum += i // sum = sum + 1 
// 		i++ //i = i + 1 
// 		if i <= 5{
// 			fmt.Println("break")
// 		}
// 	}

// }

func GetPrimeNumbers(last int) {
	// var i, j int
	
	// for i = 2; i < last; i++{
	// 	for j = 2; j <= (i/j); j++ {
	// 		if i%j == 0 {
	// 			break
	// 		}
	// 	}
	// 	if j > (i / j) {
	// 		fmt.Printf("%d is prime\n",i)
	// 	}

	// }
	plus, minus  := PlusMinus(1, 2) //calling function 
	fmt.Println("a+b =", plus)
	fmt.Println("a-b=", minus)
}

func PlusMinus( num1, num2 int) (int, int) {	
	plus := num1 + num2
	minus := num1 - num2 	
	return plus, minus 
}

func CheckSwap()	{
	var a, b int
	a = 5
	b = 10
	fmt.Printf("before swapping, a = %d\n", a)
	fmt.Printf("before swapping, b = %d\n", b)
	swapbyref(&a, &b)
	fmt.Printf("after swapping, a = %d\n", a)
	fmt.Printf("after swapping, b = %d\n", b)

}

// //call by value cant change the return for a abd b above
// func swap(x int, y int) int {
// 	var temp int
// 	temp = x //save x to temp variable
// 	x = y	// put y into x
// 	y = temp //put temp into y 
// 	fmt.Printf("inside swap function after swapping, x = %d\n", x)
// 	fmt.Printf("inside sawp function after swapping, y = %d\n", y)
// 	return temp 
// } 


//call by reference 
func swapbyref (x *int, y *int) int {
	var temp int
	temp = *x //save x to temp variable
	*x = *y	// put y into x
	*y = temp //put temp into y 
	return temp 
} 









// }
// func CheckSwap() {
// 	var a, b int
// 	a = 5
// 	b = 10
// 	fmt.Printf("Before swapping, a = %d\n", a)
// 	fmt.Printf("Before swapping, a = %d\n", b)	
// 	swap(a,b)
// 	fmt.Printf("After swapping, a = %d\n", a)
// 	fmt.Printf("After swapping, a = %d\n", b)
// }
//  func swap( x int, y int) {
// 	 var temp int
// 	 temp = *x 
// 	 *x = *y
// 	 *y = temp
// 	 return temp 
//  }
	
