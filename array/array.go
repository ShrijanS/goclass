package array

import (
 "fmt"
 "reflect"
)

func ArrayMain() {
	simpleArray()
	twoDimensionalArray()
}

func twoDimensionalArray() {
	td := [2][3] int{{1, 2, 5}, {3, 4}}
	fmt.Println(td)

	var twoDArr [2][3]int
	for i := 0; i < 2; i++{
		for j := 0; j < 3; j++{
			twoDArr[i][j] = i + j
		}
	}

	fmt.Println("twoDArr value = ",twoDArr)
}

func simpleArray(){
	var arr [5]int
	fmt.Printf("value of array %v", arr)
	fmt.Printf("type of arr %t", arr)
	fmt.Println(reflect.TypeOf(arr).Kind())

	arr[4] = 4
	fmt.Println("setting the value", arr[4])
	fmt.Println("len of arr = ", len(arr))

	arr2 := [5] int{4, 5, 2, 6, 8}
	fmt.Println("va;ue of arr2", arr2)
}