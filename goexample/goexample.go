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
	// // msg := make(chan string, 3)

	// // go message(msg)

	// // fmt.Println(<-msg)
	// // fmt.Println(<-msg)
	// // fmt.Println(<-msg)

	// done := make(chan bool, 1)
	// go worker(done)
	// <-done

	// ................

// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func () {
// 		time.Sleep(4 * time.Second)
// 		c1 <- "passing two"
// 	}()

// 	go func () {
// 		time.Sleep(1 * time.Second)
// 		c1 <-"passing one"
// 	}()

// 	for i := 0 ; i <2 ; i++ {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println(msg1)
// 		case msg2 := <-c2:
// 			fmt.Println(msg2)
// 		case <- time.After(4 * time.Second):
// 			fmt.Println("timeout")

// 		}
// 	}

	// closingChannel()
	// rangeOnChannel()
	
	exampleTimer()
 }

func exampleTimer(){
	timer1 :=time.NewTimer(2 * time.Second)

	<- timer1.C
	fmt.Println("timer 1 fired")

	timer2 :=time.NewTimer(time.Second)

	go func() {
		<- timer2.C
		fmt.Println("timer 2 fired")

	}()
	stop2 := timer2.Stop()
	if stop2{
		fmt.Println("timer 2 stopped")
	
	}
} 

func rangeOnChannel(){
	queue := make( chan	string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close (queue)

	for value := range queue {
		fmt.Println(value)

	}

}

func closingChannel(){
	jobs := make(chan int,5)
	done := make(chan bool)

	go func () {
		for {
			rec, close := <-jobs //here close is boolean
			if close{
				fmt.Println("recevied jobs=", rec)
			}else{
				fmt.Println("all jobs finished")
				done <- true
				return

				}

			}
		}()
		
		for i:= 1 ; i <= 3; i++{
			jobs <- i
			fmt.Println("send jobs =", i)

		}
		close(jobs)
		fmt.Println("closed")
		<- done 
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

