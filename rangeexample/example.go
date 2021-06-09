package rangeexample

import "fmt"

func RangeExample(){
	// exampleOne()
	// exampleMap()
	exampleTwo()

}

func exampleTwo(){
	for i, value := range "golang"{
		fmt.Println(i, value)
	}

}
func exampleMap(){
	//map key => value one => 1 
	// n := map[string]int{"one": 1, "two":2}
	// fmt.Println("map =", n)

	
	m := make(map[string]int)

	m["one"]= 1
	m["two"]= 2
	m["three"]= 3	
	m["four"]= 4
	delete(m, "three")
	fmt.Println("map =", m)

	for key, value := range m {
		fmt.Printf("%v = %v\n", key, value)
	
	}
}

	func exampleOne() {
	nums := []int {2,3,4,5}
	sum := 0
	// sum = nums[0]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("sum = ", sum)

	sum = 0
	for _, num := range nums {
		sum += num

	}
	fmt.Println("sum = ", sum)
}