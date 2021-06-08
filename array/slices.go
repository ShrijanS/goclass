package array

import(
	"fmt"
)
func SlicesMain(){
	intro()
}

func intro(){
	// var arr []int {1, 3, 5}
	// fmt.Printf("value of array %v\n", arr)
	// fmt.Printf("type of arr %t\n", arr)
	// fmt.Printf("lenght of arr %v\n", len(arr))
	// fmt.Println(reflect.TypeOf(arr).Kind())

	s := make([]string, 6)
	fmt.Println("slices:", s)

	s[0] = "g"
	s[1] = "o"
	s[2] = "l"
	s[3] = "a"
	s[4] = "n"
	s[5] = "g"
	fmt.Println("Slices:", s)
	fmt.Printf("lens = %v, cap = %v", len(s), cap(s))

	s = s[:]
	fmt.Println(s)
	fmt.Printf("lens = %v, cap = %v", len(s), cap(s))



	s = s[2:4]
	fmt.Println(s)
	fmt.Printf("lens = %v, cap = %v", len(s), cap(s))

	s = s[:4]
	fmt.Println(s)
	fmt.Printf("lens = %v, cap = %v", len(s), cap(s))

	s = s[2:]
	fmt.Println(s)
	fmt.Printf("lens = %v, cap = %v", len(s), cap(s))
	
	s = s[:]
	fmt.Println(s)
	fmt.Printf("lens = %v, cap = %v", len(s), cap(s))

}
