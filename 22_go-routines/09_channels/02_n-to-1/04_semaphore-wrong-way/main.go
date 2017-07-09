package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// we block here until done <- true
	<-done
	<-done
	close(c)

	// to unblock above
	// we need to take values off chan c here
	// but we never get here, because we're blocked above

	for n := range c {
		fmt.Println(n)
	}

	// for the rpgram to put onto channel c it has to 'receive' from c channel.
	// it has to be at the same time, in sync, like relay race, pass the baton
	// but because <- done is not inside a go routine it won't run paralel,
	// because it won't run paralel it won't get nothing from done channel
	// and it will block the rest of the program flow in that place.
	// When the first and second go routines want to put something into the c channel
	// they can't because we never got to the range c line, so they're out of sync
	// they can't put nothing into the channel if there is nothing to receive it.
	// We never got to range c line that would receive the value because program blocked at <-done
	// that's why <-done has be inside another go routine to do it's thing and be out of the way
	// and not block the control flow so we can reach range c.

}
