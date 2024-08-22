package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test"};
var wg sync.WaitGroup; //pointer
var mut sync.Mutex;

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
	}

	wg.Wait();
	fmt.Println(signals);
}

func getStatusCode(endpoint string) {

	defer wg.Done();	

	res, err := http.Get(endpoint);

	if(err != nil) {
		fmt.Println("OOPS...");
	} else {
		mut.Lock();
		signals = append(signals, endpoint);
		mut.Unlock();
		fmt.Printf("%d staus code for %s\n", res.StatusCode, endpoint);
	}


}