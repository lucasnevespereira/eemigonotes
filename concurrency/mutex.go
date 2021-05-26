package main

import (
"fmt"
"sync"
)

func main() {
	var mx sync.Mutex	
	done := make(chan bool)
    m := make(map[string]string)
	m["name"] = "world"
	go func() {
		mx.Lock()
		m["name"] = "data race"
		mx.Unlock()		
		done <- true
	}()
	mx.Lock()
	fmt.Println("Hello,", m["name"])
	mx.Unlock()
	<-done
}
