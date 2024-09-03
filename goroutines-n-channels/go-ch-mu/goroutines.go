package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	start := time.Now()         // start time
	id := getUserByName("john") // calling function
	println(id)                 // printing results

	wg := sync.WaitGroup{}       // creating a wait group
	ch := make(chan *Message, 2) // creating a channel with buffer size 2

	wg.Add(2) // adding 2 goroutines to wait group

	// blocking and non-concurrent implementation
	// chats := getUserChats() // calling function
	// friends := getUserFriends() // calling a function
	// log.Println(chats)
	// log.Println(friends)

	// concurrent implementation
	go getUserChats(id, ch, &wg)
	go getUserFriends(id, ch, &wg)

	wg.Wait() // waiting for all goroutines to finish
	close(ch) // closing the channel

	log.Println(<-ch)              // receiving Message from channel
	log.Println(<-ch)              // receiving Message from channel
	log.Println(time.Since(start)) // counting time since start of main functin

}

// writing a simple non-concurrent function for fetching friends for user
func getUserFriends(_ string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 3) // simulate some time consuming operation

	ch <- &Message{ // giving Message struct initialised to the channel
		friends: []string{"John",
			"Mary",
			"Peter",
			"Paul",
			"George",
			"Ringo",
		},
	}

	// this is not recommended, if you want to use channel for other purposes later
	// close(ch) // closing the channel
	// instead we can use wait group to wait for the goroutine to finish
	wg.Done() // done with the second goroutine
}

// writing a simple non-concurrent function for fetching chats for user
func getUserChats(_ string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2) // simulate some time consuming operation

	ch <- &Message{ // giving Message struct initialised to the channel
		chats: []string{
			"chat-1",
			"chat-2",
			"chat-3",
		},
	}

	wg.Done() // done with the first goroutine
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1) // simulate some time consuming operation

	return fmt.Sprintf("%s-2", name) // return the id
}
