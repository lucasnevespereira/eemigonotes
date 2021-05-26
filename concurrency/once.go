package main

import (
  "fmt"
  "sync"
)

var doOnce sync.Once

func main() {
  DoSomething()
}

func DoSomething() {
  doOnce.Do(func() {
	fmt.Println("Run once, loading...")
  })
  
  fmt.Println("Run every time")
}


