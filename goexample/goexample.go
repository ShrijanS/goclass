package goexample

import (
"fmt"
"time"
)
type User struct{
	name string
	email string
}

type Admin struct{
	User
	email string
	keyrings string
}

func Examples() {
	admin := Admin{
		User: User{name: "example", email: "example@gmail.com" },
		email: "admin@gmail.com",
		keyrings: "my keys",

	}
	fmt.Println(admin.name)
	fmt.Println(admin.email)
	fmt.Println(admin.User.email)
}





//Channels
//connection between concurrent goroutines
func Example2() {
	// msg := make(chan string, 3)

	// go message(msg)

	// fmt.Println(<-msg)
	// fmt.Println(<-msg)
	// fmt.Println(<-msg)

	done := make(chan bool, 1)
	go worker(done)
	<-done

}

func worker(done chan bool) {
	fmt.Println("working....")
	time.Sleep(time.Second * 2)
	fmt.Println("done")
	done <- true
}
func message(channel chan<- string){
	channel <- "first data"
	channel <- "second data"
	channel <- " thrid data"
}

