package main

import "fmt"

func output_stuff(message string, chann chan bool) {
	fmt.Println(message)
	chann <- true
}

func main() {
	// normally, if you start a function call with "go", it is a go routine
	/* But they do not return anything. So, we need something to catch what they return.
	They are called channels. They act as communication mediums to send out whatever happens in routines.*/
	done := make(chan bool)
	go output_stuff("Hi bro I am going to be shoved into a channel", done)
	<-done

	for ind := 1; ind <= 20; ind++ {
		fmt.Println(ind * ind)
	}

}
