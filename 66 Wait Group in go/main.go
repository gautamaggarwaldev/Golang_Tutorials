package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)
 
var wg sync.WaitGroup; //pointer

func main() {
	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://amazon.in",
		"https://fb.com",
	}

	for _,i := range websiteList {
		go getStatusCode(i);
		wg.Add(1);
		// time.Sleep(1*time.Second);
	}

	wg.Wait();
}

func getStatusCode(endpoint string) {

	defer wg.Done();	

	res, err := http.Get(endpoint);

	if(err != nil) {
		fmt.Println("OOPS...");
	} else {
		fmt.Printf("%d staus code for %s\n", res.StatusCode, endpoint);
	}


}