package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//Channel uses keyword chan, channel syncs data and comm data through channels
	// //Declaration
	// Channel := make(chan int)
	// //Sending data to channel
	// Channel <- 15
	// //Receiving data to channel
	// value := <- Channel
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewherintheinternet.com",
		"https://graph.microsoft.com",
	}
	//Creating channel
	ch := make(chan string)

	for _, api := range apis {
		//Concurrency
		//Go run time with keyword Go to execute different func in different goroutines

		go checkAPI(api, ch)
	}
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
	//Concurrency executes routines faster, so we cant see the results from fnc checkAPI bcs its happens so fast then, we apply sleep
	time.Sleep(1 * time.Second)
	tlapse := time.Since(start)
	fmt.Printf("Succeeded in %v seconds", tlapse.Seconds())
}
func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		//Sending data to channel successfully
		ch <- fmt.Sprintf("Error getting %s\n", api)
		return
	}
	//Sending data to channel failed
	ch <- fmt.Sprintf("Successfully getting API %s\n", api)
}
